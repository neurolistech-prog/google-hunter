package core

import (
	"fmt"
	"net/http"
	"time"
)

// Result structure pour stocker les découvertes
type Result struct {
	URL   string
	Found bool
	Msg   string
}

// ScanTarget vérifie les vulnérabilités de base
func ScanTarget(url string) Result {
	client := &http.Client{
		Timeout: 5 * time.Second,
		// Empêcher de suivre les redirections pour analyser le serveur précis
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		return Result{URL: url, Found: false, Msg: err.Error()}
	}
	defer resp.Body.Close()

	// Logique de détection : Absence de headers de sécurité
	if resp.Header.Get("Content-Security-Policy") == "" {
		return Result{URL: url, Found: true, Msg: "Manque la politique CSP (potentiel XSS)"}
	}

	return Result{URL: url, Found: false, Msg: "RAS"}
}
