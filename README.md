# google-hunter
Fast &amp; Lightweight Go-based Security Agent for automated reconnaissance and vulnerability detection in Bug Bounty programs.

# 🎯 Google Hunter Agent

Un agent de sécurité automatisé écrit en **Go**, conçu pour la reconnaissance rapide et la détection de vulnérabilités (Bug Bounty).

## 🚀 Fonctionnalités
- **Performance Go** : Utilise les goroutines pour un scan multi-thread ultra-rapide.
- **Détection Automatisée** : Analyse les en-têtes HTTP et les configurations de sécurité.
- **GitHub Actions** : Planification de scans quotidiens via CI/CD.

## 🛠️ Installation

```bash
# Cloner le dépôt
git clone [https://github.com/TON_PSEUDO/google-hunter.git](https://github.com/TON_PSEUDO/google-hunter.git)
cd google-hunter

# Compiler l'agent
go build -o hunter main.go
