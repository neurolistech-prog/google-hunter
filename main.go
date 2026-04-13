package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"google-hunter/core"
	"net/http"
	"sync"
	"time"
)

// Structure pour le message Discord
type DiscordMessage struct {
	Content string `json:"content"`
}

const webhookURL = "https://discord.com/api/webhooks/1493290406048436386/6whKo-aV73x_wRR4ELaoUIJwgLjHeh4LLUklc1XqOoHn5srj-ExzUwhlr0tPdKNWNqxl"

// sendDiscordAlert envoie une notification push sur ton serveur Discord
func sendDiscordAlert(message string) {
	msg := DiscordMessage{Content: "🚀 **Agent Alert** : " + message}
	payload, _ := json.Marshal(msg)

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Printf("Erreur lors de l'envoi Discord : %v\n", err)
		return
	}
	defer resp.Body.Close()
}

func main() {
	// Cible de test : ton portfolio
	targets := []string{
		"https://neurolistech-prog.github.io/Portfolio-2",
		"https://www.google.com",
	}

	var wg sync.WaitGroup
	fmt.Println("--- 🛡️ Lancement de l'Agent Google-Hunter ---")
	fmt.Printf("--- 🎯 Cibles : %v ---\n", len(targets))

	for _, url := range targets {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			
			// On appelle le scanner que nous avons mis à jour dans core/scanner.go
			found, details := core.ScanTarget(u)
			
			if found {
				fullMsg := fmt.Sprintf("Faille détectée sur %s\nDétails : %s", u, details)
				fmt.Println("[!!!]", fullMsg)
				sendDiscordAlert(fullMsg)
			} else {
				fmt.Printf("[+] %s analysé (aucune faille critique).\n", u)
			}
		}(url)
	}

	wg.Wait()
	fmt.Println("--- ✅ Scan terminé ---")
}
