package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	log.Println("Starting the PingPong application...")
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/pong", pongHandler)
	http.HandleFunc("/professional-ping-pong", professionalPingPongHandler)
	http.HandleFunc("/amateur-ping-pong", amateurPingPongHandler)
	http.HandleFunc("/chance-ping-pong", chancePingPongHandler)
	log.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("PingHandler called")
	response := map[string]string{"response": "pong"}
	jsonResponse(w, response)
}

func pongHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("PongHandler called")
	response := map[string]string{"response": "ping"}
	jsonResponse(w, response)
}

func professionalPingPongHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ProfessionalPingPongHandler called")
	if rand.Intn(100) < 90 {
		response := map[string]string{"response": "pong"}
		jsonResponse(w, response)
	} else {
		response := map[string]string{"response": "ping"}
		jsonResponse(w, response)
	}
}

func amateurPingPongHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("AmateurPingPongHandler called")
	if rand.Intn(100) < 70 {
		response := map[string]string{"response": "pong"}
		jsonResponse(w, response)
	} else {
		response := map[string]string{"response": "ping"}
		jsonResponse(w, response)
	}
}

func chancePingPongHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ChancePingPongHandler called")
	if rand.Intn(100) < 50 {
		response := map[string]string{"response": "pong"}
		jsonResponse(w, response)
	} else {
		response := map[string]string{"response": "ping"}
		jsonResponse(w, response)
	}
}

func jsonResponse(w http.ResponseWriter, data map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
