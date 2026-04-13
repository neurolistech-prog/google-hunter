package main

import (
	"fmt"
	"google-hunter/core"
	"sync"
)

func main() {
	// Liste des sous-domaines de Google (à titre d'exemple)
	targets := []string{
		"https://www.google.com",
		"https://mail.google.com",
		"https://drive.google.com",
	}

	var wg sync.WaitGroup
	results := make(chan core.Result, len(targets))

	fmt.Println("--- Lancement de l'Agent de Reconnaissance ---")

	for _, url := range targets {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			results <- core.ScanTarget(u)
		}(url)
	}

	// Fermer le canal quand tout est fini
	go func() {
		wg.Wait()
		close(results)
	}()

	// Affichage des rapports
	for res := range results {
		if res.Found {
			fmt.Printf("[!] VULNÉRABILITÉ : %s | %s\n", res.URL, res.Msg)
		} else {
			fmt.Printf("[+] OK : %s\n", res.URL)
		}
	}
}
