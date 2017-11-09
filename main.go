package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	urls := map[string]string{
		"valutac":   "URL_NOTIFICATION_ONE",
		"valutac2": "URL_NOTIFICATION_TWO",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method: %s. RemoteAddr: %s\n", r.Method, r.RemoteAddr)
		// early return
		if r.Method == http.MethodGet {
			fmt.Fprintf(w, "Midtrans Bridge\n")
			return
		}
		var resp Response
		if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
			fmt.Fprintf(w, "Error: %s!\n", err.Error())
			return
		}
		prefix := strings.SplitN(resp.OrderID, "-", 2)[0]
		switch strings.ToLower(prefix) {
		case "valutac":
			fmt.Fprintf(w, "URL: %s!\n", urls["valutac"])
			send(urls["valutac"], resp)
		case "valutac2":
			fmt.Fprintf(w, "URL: %s!\n", urls["valutac2"])
			send(urls["valutac2"], resp)
		default:
			fmt.Fprintf(w, "Unknown Transaction\n")
		}
	})
	log.Println("Running HTTP Server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func send(url string, resp Response) error {
	jsonStr, _ := json.Marshal(resp)
	_, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Callback for %s, sent to %s. Got Err: %s\n", resp.OrderID, url, err.Error())
		return err
	}
	fmt.Fprintf(os.Stdout, "Callback for %s, sent to %s\n", resp.OrderID, url)
	return nil
}

// Response received callback
type Response struct {
	StatusCode        string `json:"status_code"`
	StatusMessage     string `json:"status_message"`
	PermataVaNumber   string `json:"permata_va_number"`
	SignKey           string `json:"signature_key"`
	Bank              string `json:"bank"`
	ReURL             string `json:"redirect_url"`
	ECI               string `json:"eci"`
	FraudStatus       string `json:"fraud_status"`
	PaymentType       string `json:"payment_type"`
	OrderID           string `json:"order_id"`
	TransactionId     string `json:"transaction_id"`
	TransactionTime   string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
	GrossAmount       string `json:"gross_amount"`
	PaymentCode       string `json:"payment_code"`
}