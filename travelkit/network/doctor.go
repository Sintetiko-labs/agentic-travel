package network

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	egressURL = "https://ifconfig.me/ip"
	ipAPIURL  = "http://ip-api.com/json/?fields=status,message,query,isp,org,as,hosting"
	doctorUA  = "travelkit-network-doctor/1.0"
)

type DoctorStatus string

const (
	DoctorOK           DoctorStatus = "ok"
	DoctorProxySet     DoctorStatus = "proxy_env_set"
	DoctorEgressFailed DoctorStatus = "egress_failed"
	DoctorDatacenter   DoctorStatus = "datacenter_egress"
)

type DoctorResult struct {
	Status   DoctorStatus `json:"status"`
	ProxyEnv []string     `json:"proxy_env,omitempty"`
	EgressIP string       `json:"egress_ip,omitempty"`
	ISP      string       `json:"isp,omitempty"`
	Org      string       `json:"org,omitempty"`
	AS       string       `json:"as,omitempty"`
	Hosting  bool         `json:"hosting"`
	Message  string       `json:"message"`
	NextStep string       `json:"next_step,omitempty"`
}

func Doctor(ctx context.Context) DoctorResult {
	res := DoctorResult{}
	if proxy := SetProxyVars(); len(proxy) > 0 {
		res.ProxyEnv = proxy
		res.Status = DoctorProxySet
		res.Message = "proxy env vars are set — CLIs ignore them but Chrome session capture may still use a proxy"
		res.NextStep = "unset HTTP_PROXY HTTPS_PROXY ALL_PROXY, or run with --no-proxy"
	}
	if ctx == nil {
		ctx = context.Background()
	}
	ip, err := fetchEgressIP(ctx)
	if err != nil {
		if res.Status == "" {
			res.Status = DoctorEgressFailed
		}
		if res.Message != "" {
			res.Message += "; "
		}
		res.Message += fmt.Sprintf("egress probe failed: %v", err)
		if res.NextStep == "" {
			res.NextStep = "curl -s https://ifconfig.me from Terminal on your Mac home network"
		}
		return res
	}
	res.EgressIP = strings.TrimSpace(ip)
	meta, err := lookupIP(ctx, res.EgressIP)
	if err != nil {
		if res.Status == "" {
			res.Status = DoctorOK
		}
		if res.Message == "" {
			res.Message = fmt.Sprintf("egress IP %s (ISP lookup failed: %v) — verify manually at https://ifconfig.me", res.EgressIP, err)
		} else {
			res.Message += fmt.Sprintf("; egress IP %s (ISP lookup failed)", res.EgressIP)
		}
		if res.NextStep == "" {
			res.NextStep = "confirm IP is your Mac home/office Wi‑Fi, not VPN or cloud egress"
		}
		return res
	}
	res.ISP, res.Org, res.AS, res.Hosting = meta.ISP, meta.Org, meta.AS, meta.Hosting
	if meta.Hosting {
		res.Status = DoctorDatacenter
		res.Message = fmt.Sprintf("egress IP %s looks like datacenter/hosting (%s)", res.EgressIP, meta.Org)
		res.NextStep = "disconnect VPN/cloud egress; run CLIs from Mac home residential IP"
		return res
	}
	if res.Status == DoctorProxySet {
		res.Message += fmt.Sprintf("; egress IP %s (%s) appears residential", res.EgressIP, firstNonEmpty(meta.ISP, meta.Org))
		return res
	}
	res.Status = DoctorOK
	res.Message = fmt.Sprintf("residential egress OK — %s (%s)", res.EgressIP, firstNonEmpty(meta.ISP, meta.Org))
	return res
}

type ipMeta struct {
	ISP, Org, AS string
	Hosting      bool
}

func fetchEgressIP(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, egressURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("user-agent", doctorUA)
	resp, err := DirectClient(15 * time.Second).Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, 64))
	if err != nil {
		return "", err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("ifconfig.me HTTP %d", resp.StatusCode)
	}
	return string(body), nil
}

func lookupIP(ctx context.Context, ip string) (ipMeta, error) {
	url := strings.Replace(ipAPIURL, "/json/", "/json/"+ip+"?", 1)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return ipMeta{}, err
	}
	req.Header.Set("user-agent", doctorUA)
	resp, err := DirectClient(15 * time.Second).Do(req)
	if err != nil {
		return ipMeta{}, err
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(io.LimitReader(resp.Body, 4096))
	if err != nil {
		return ipMeta{}, err
	}
	var out struct {
		Status, Message string
		ISP, Org, AS    string
		Hosting         bool
	}
	if err := json.Unmarshal(raw, &out); err != nil {
		return ipMeta{}, err
	}
	if out.Status != "success" {
		return ipMeta{}, fmt.Errorf("ip-api: %s", firstNonEmpty(out.Message, out.Status))
	}
	return ipMeta{ISP: out.ISP, Org: out.Org, AS: out.AS, Hosting: out.Hosting}, nil
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

func PrintDoctor(res DoctorResult) {
	if len(res.ProxyEnv) > 0 {
		fmt.Fprintf(os.Stderr, "proxy env: %s\n", strings.Join(res.ProxyEnv, ", "))
	}
	if res.EgressIP != "" {
		fmt.Fprintf(os.Stderr, "egress ip: %s\n", res.EgressIP)
	}
	if res.ISP != "" || res.Org != "" {
		fmt.Fprintf(os.Stderr, "isp/org:   %s / %s\n", res.ISP, res.Org)
	}
	if res.AS != "" {
		fmt.Fprintf(os.Stderr, "as:        %s\n", res.AS)
	}
	if res.EgressIP != "" {
		fmt.Fprintf(os.Stderr, "hosting:   %v\n", res.Hosting)
	}
	fmt.Fprintf(os.Stderr, "status:    %s\n", res.Status)
	fmt.Fprintln(os.Stderr, res.Message)
	if res.NextStep != "" {
		fmt.Fprintln(os.Stderr, "next:", res.NextStep)
	}
}
