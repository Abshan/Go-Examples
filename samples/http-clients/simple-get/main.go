package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get("https://example.com")
	if err != nil {
		fmt.Println("request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body error:", err)
		return
	}

	fmt.Printf("status: %s\n", resp.Status)
	previewLen := 120
	if len(body) < previewLen {
		previewLen = len(body)
	}
	fmt.Printf("body preview (%d bytes): %s\n", previewLen, string(body[:previewLen]))
}
