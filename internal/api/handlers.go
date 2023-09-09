package api

import (
	"api/internal/channels"
	"api/internal/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func executeRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var payload types.ExecutePayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, "Malformed request", http.StatusBadRequest)
		return
	}

	// TODO: Validate request params

	transactionRequestChannel := channels.GetTransactionRequestChannel()
	responseChannel := make(chan types.ExecuteResponse)

	transactionRequestChannel <- types.TransactionForSigning{Request: payload, ResponseChannel: responseChannel}
	transactionResponse := <-responseChannel

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactionResponse)
}
