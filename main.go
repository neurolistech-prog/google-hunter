package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"google-hunter/core"
	"net/http"
	"os"
	"sync"
)

type DiscordMessage struct {
	Content string `json:"content"`
}

const webhookURL = "https://discord.com/api/webhooks/1493290406048436386/6whKo-aV73x_wRR4ELaoUIJwgLjHeh4LLUklc1XqOoHn5srj-ExzUwhlr0tPdKNWNqxl"

func sendDiscordAlert(message string) {
	msg := DiscordMessage{Content: message}
	payload, _ := json.Marshal(msg)
	http.Post(webhookURL, "application/json", bytes.NewBuffer(payload))
}

func main() {
	// 1. Lire les cibles depuis targets.txt
	file, err := os.Open("targets.txt")
	var targets []string
	if err != nil {
		fmt.Println("[!] Fichier targets.txt introuvable, utilisation des cibles par défaut.")
		targets = []string{"https://www.google.com"}
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			targets = append(targets, scanner.Text())
		}
		file.Close()
	}

	var wg sync.WaitGroup
	var countFound int
	fmt.Println("--- 🛡️ Agent Google-Hunter en ligne ---")

	for _, url := range targets {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			found, details := core.ScanTarget(u)
			if found {
				countFound++
				sendDiscordAlert("🚨 **FAILLE DÉTECTÉE**\nCible : " + u + "\nType : " + details)
			}
		}(url)
	}

	wg.Wait()

	// 2. Message de garantie (Rapport quotidien)
	report := fmt.Sprintf("✅ **Rapport Quotidien**\nL'agent a scanné **%d** cibles.\nFailles trouvées : **%d**.\nStatut : Opérationnel 🟢", len(targets), countFound)
	sendDiscordAlert(report)
	
	fmt.Println("--- ✅ Scan terminé et rapport envoyé ---")
}
