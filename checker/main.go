package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var validFlags = map[string]bool{
	"CTF{d0m_xss_w1th_l0c4lst0r4g3}": true,
	"CTF{h4sh_b4s3d_xss_1s_r3al}":    true,
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	flag := r.URL.Query().Get("flag")
	if flag == "" {
		http.Error(w, "❌ flag param missing", http.StatusBadRequest)
		return
	}

	flag = strings.TrimSpace(flag)

	if validFlags[flag] {
		fmt.Fprintf(w, "✅ Correct! Flag accepted. \n")
		log.Printf("[+] Valid flag: %s | from %s", flag, r.RemoteAddr)
	} else {
		fmt.Fprintf(w, "❌ Wrong flag. \n")
		log.Printf("[-] Invalid flag: %s | from %s", flag, r.RemoteAddr)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "✅ Checker is alive. \n")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port

	http.HandleFunc("/submit", submitHandler)
	http.HandleFunc("/health", healthHandler)

	log.Println("Checker started", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
