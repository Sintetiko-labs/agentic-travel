package session

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDoctorIncapsulaChallengeOn200(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`<html><script src="/_Incapsula_Resource?SWJIYLWA=x"></script></html>`))
	}))
	defer srv.Close()

	res := Doctor(DoctorOptions{
		Slug:      "iberiaexpress",
		EnvPrefix: "IBERIAEXPRESS",
		BaseURL:   "https://www.iberiaexpress.com",
		Cookie:    "visid_incap=1; incap_ses_1=2",
		ProbeURL:  srv.URL,
	})
	if res.Status != DoctorBlocked {
		t.Fatalf("status=%s msg=%s", res.Status, res.Message)
	}
}
