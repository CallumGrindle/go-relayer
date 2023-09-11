package api

import (
	"api/internal/channels"
	"api/internal/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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
	responseChannel := make(chan types.ExecuteResponse, 1)
	errorChannel := make(chan error, 1)

	transactionRequestChannel <- types.TransactionForSigning{ExecutePayload: payload, ResponseChannel: responseChannel, ErrorChannel: errorChannel}

	select {
	case transactionResponse := <-responseChannel:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(transactionResponse)
	case <-errorChannel:
		http.Error(w, "An error occurred", http.StatusInternalServerError)
	case <-time.After(5 * time.Second):
		http.Error(w, "request timed out", http.StatusRequestTimeout)

	}
}
