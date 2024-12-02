package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var rates map[string]float64

// Fetch rates from Flask service
func fetchRates() {
	resp, err := http.Get("http://express_service:3000/api/rates")
    if err != nil {
        fmt.Println("Error fetching rates:", err)
        return
    }
    defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error fetching rates: received status code %d\n", resp.StatusCode)
		return
	}

    // Use an interface to handle both float64 and string
    var rawRates map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&rawRates); err != nil {
        fmt.Println("Error decoding response:", err)
        return
    }

    // Convert raw rates to float64
    rates = make(map[string]float64)
    for k, v := range rawRates {
        switch value := v.(type) {
        case float64:
            rates[k] = value
        case string:
            if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
                rates[k] = floatValue
            } else {
                fmt.Printf("Error converting %s: %v\n", k, err)
            }
        default:
            fmt.Printf("Unexpected type for %s: %T\n", k, v)
        }
    }
}

func convertCurrency(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	amountStr := r.URL.Query().Get("amount")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	var amount float64
	fmt.Sscanf(amountStr, "%f", &amount)

	rateFrom := rates[from]
	rateTo := rates[to]
	convertedAmount := (amount * rateTo / rateFrom)

	response := map[string]float64{
		"convertedAmount": convertedAmount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	fetchRates() // Initial fetch of rates

	http.HandleFunc("/api/convert", convertCurrency)

	fmt.Println("Go service running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
