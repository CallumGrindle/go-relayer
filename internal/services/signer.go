package services

import (
	"api/internal/channels"
	"api/internal/types"
	"log"
)

func Signer() {
	transactionRequestChannel := channels.GetTransactionRequestChannel()

	for {
		select {
		case transactionRequest := <-transactionRequestChannel:
			go func(req types.TransactionForSigning) {
				log.Print("Signing transaction")

				transactionRequest.ResponseChannel <- types.ExecuteResponse{TransactionHash: "0x123"}
				log.Printf("Signed Transaction: %v", transactionRequest.Request)
			}(transactionRequest)
		}
	}
}
