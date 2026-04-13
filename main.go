package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"google-hunter/core"
	"net/http"
	"sync"
)

type DiscordMessage struct {
	Content string `json:"content"`
}

const webhookURL = "https://discord.com/api/webhooks/1493290406048436386/6whKo-aV73x_wRR4ELaoUIJwgLjHeh4LLUklc1XqOoHn5srj-ExzUwhlr0tPdKNWNqxl"

func sendDiscordAlert(message string) {
	msg := DiscordMessage{Content: "🚀 **Agent Alert** : " + message}
	payload, _ := json.Marshal(msg)
	http.Post(webhookURL, "application/json", bytes.NewBuffer(payload))
}

func main() {
	targets := []string{
		"https://neurolistech-prog.github.io/Portfolio-2",
		"https://www.google.com",
	}

	var wg sync.WaitGroup
	fmt.Println("--- 🛡️ Agent Google-Hunter en ligne ---")

	for _, url := range targets {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			found, details := core.ScanTarget(u)
			if found {
				sendDiscordAlert(fmt.Sprintf("Faille sur %s : %s", u, details))
				fmt.Printf("[!!!] Alerte envoyée pour %s\n", u)
			} else {
				fmt.Printf("[+] %s : RAS\n", u)
			}
		}(url)
	}
	wg.Wait()
	fmt.Println("--- ✅ Scan terminé ---")
}
