package core

import (
	"fmt"
	"net/http"
	"time"
)

// ScanTarget analyse la cible et retourne un message si une anomalie est trouvée
func ScanTarget(url string) (bool, string) {
	client := &http.Client{Timeout: 5 * time.Second}
	
	// Liste de fichiers à tester sur ton portfolio
	files := []string{"/", "/index.html", "/.git/config", "/README.md"}

	for _, file := range files {
		fullURL := url + file
		resp, err := client.Get(fullURL)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		// Si on trouve un fichier qui ne devrait pas être là (ex: .git) 
		// ou si on confirme que le site répond bien
		if resp.StatusCode == 200 {
			if file == "/.git/config" {
				return true, "ALERTE : Dossier Git exposé sur " + fullURL
			}
			fmt.Printf("[+] Check réussi sur : %s\n", fullURL)
		}
	}
	return false, ""
}
