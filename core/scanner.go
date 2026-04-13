package core

import (
	"net/http"
	"time"
)

// ScanTarget analyse la cible et retourne (VRAI, Détails) si une anomalie est trouvée
func ScanTarget(url string) (bool, string) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	
	// On teste la racine et un dossier sensible théorique
	paths := []string{"/", "/.git/config"}

	for _, path := range paths {
		fullURL := url + path
		resp, err := client.Get(fullURL)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		// Si le serveur répond 200 (OK) sur /.git/config, c'est une faille critique
		if resp.StatusCode == 200 && path == "/.git/config" {
			return true, "Dossier .git exposé ! (Fuite de code source)"
		}
	}

	return false, ""
}
