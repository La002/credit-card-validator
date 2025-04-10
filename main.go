package main

import (
	"CreditCard/checker"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting")
	r := mux.NewRouter()
	r.HandleFunc("/{number}", handleGet).Methods("GET")

	// Add CORS middleware
	corsHandler := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))

	srv := &http.Server{
		Handler: corsHandler(r),
		Addr:    ":8082",
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

	//	fmt.Println(checker.CheckNumberValid("3469354508162306"))
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	num := mux.Vars(r)["number"]
	isValid := checker.CheckNumberValid(num)
	if !isValid {
		_, _ = w.Write([]byte("The number is not valid"))
		return
	}

	network := checker.GetNetwork(num)
	if network == "" {
		_, _ = w.Write([]byte("Number has no valid network"))
		return
	}

	respBody := resBody{
		CreditNetwork:    network,
		CreditCardNumber: num,
	}
	respData, err := json.Marshal(respBody)
	if err != nil {
		fmt.Println("Error marshalling response body:", err)
		return
	}

	_, err = w.Write(respData)
	if err != nil {
		fmt.Println("Error writing response:", err)
		return
	}

}

type resBody struct {
	CreditCardNumber string `json:"credit_card_number"`
	CreditNetwork    string `json:"credit_network"`
}
