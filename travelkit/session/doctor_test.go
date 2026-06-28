package session

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDoctorPOSTProbeBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.Header.Get("content-type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	res := Doctor(DoctorOptions{
		Slug:             "demo",
		EnvPrefix:        "DEMO",
		BaseURL:          "https://example.com",
		Cookie:           "_abck=1; bm_sz=2",
		ProbeURL:         srv.URL,
		ProbeMethod:      http.MethodPost,
		ProbeBody:        `{"text":"Madrid"}`,
		ProbeContentType: "application/json",
	})
	if res.Status != DoctorOK {
		t.Fatalf("status=%s msg=%s", res.Status, res.Message)
	}
	if res.ProbeHTTPStatus != http.StatusOK {
		t.Fatalf("probe status=%d", res.ProbeHTTPStatus)
	}
}

func TestDoctorSessionOptional(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	res := Doctor(DoctorOptions{
		Slug:            "demo",
		ProbeURL:        srv.URL,
		SessionOptional: true,
	})
	if res.Status != DoctorOK {
		t.Fatalf("status=%s msg=%s", res.Status, res.Message)
	}
}
