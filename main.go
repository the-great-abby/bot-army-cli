package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
)

var nc *nats.Conn

func init() {
	var err error
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = "nats://localhost:4222"
	}

	nc, err = nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	fmt.Printf("Connected to NATS at %s\n", natsURL)
}

func briefingToday() {
	fmt.Println("\n📋 Today's Briefing...")

	msg, err := nc.Request("bridge.brief.today", []byte("{}"), 5*time.Second)
	if err != nil {
		fmt.Printf("❌ Briefing request failed: %v\n", err)
		return
	}

	var response map[string]interface{}
	if err := json.Unmarshal(msg.Data, &response); err != nil {
		fmt.Printf("❌ Failed to parse response: %v\n", err)
		return
	}

	if data, ok := response["data"].(map[string]interface{}); ok {
		if briefing, ok := data["briefing"].(string); ok {
			fmt.Println(briefing)
			return
		}
	}

	fmt.Printf("Briefing response: %v\n", response)
}

func contextCurrent() {
	fmt.Println("\n🎯 Current Context...")

	msg, err := nc.Request("bridge.context.current", []byte("{}"), 5*time.Second)
	if err != nil {
		fmt.Printf("❌ Context request failed: %v\n", err)
		return
	}

	var response map[string]interface{}
	if err := json.Unmarshal(msg.Data, &response); err != nil {
		fmt.Printf("❌ Failed to parse response: %v\n", err)
		return
	}

	if data, ok := response["data"].(map[string]interface{}); ok {
		if context, ok := data["context"].(map[string]interface{}); ok {
			b, _ := json.MarshalIndent(context, "", "  ")
			fmt.Println(string(b))
			return
		}
	}

	fmt.Printf("Context response: %v\n", response)
}

func chat(query string) {
	fmt.Printf("\n🤖 Chat: %s\n", query)

	payload := map[string]string{"query": query}
	payloadBytes, _ := json.Marshal(payload)

	msg, err := nc.Request("bridge.chat", payloadBytes, 10*time.Second)
	if err != nil {
		fmt.Printf("❌ Chat request failed: %v\n", err)
		return
	}

	var response map[string]interface{}
	if err := json.Unmarshal(msg.Data, &response); err != nil {
		fmt.Printf("❌ Failed to parse response: %v\n", err)
		return
	}

	if data, ok := response["data"].(map[string]interface{}); ok {
		if answer, ok := data["answer"].(string); ok {
			fmt.Println(answer)
			return
		}
	}

	fmt.Printf("Chat response: %v\n", response)
}

func main() {
	defer nc.Close()

	fmt.Println("=== Bot Army CLI ===")
	fmt.Println("Commands: brief, context, chat <query>, help, quit")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		parts := strings.SplitN(input, " ", 2)
		cmd := parts[0]

		switch cmd {
		case "brief":
			briefingToday()
		case "context":
			contextCurrent()
		case "chat":
			if len(parts) > 1 {
				chat(parts[1])
			} else {
				fmt.Println("❌ Usage: chat <query>")
			}
		case "help":
			fmt.Println(`
Commands:
  brief             - Get today's briefing
  context           - Show current context
  chat <query>      - Chat with LLM
  quit              - Exit
  help              - Show this help
`)
		case "quit", "exit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Printf("❌ Unknown command: %s\n", cmd)
		}
	}
}
