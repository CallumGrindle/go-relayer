package services

import (
	"api/internal/channels"
	"api/internal/types"
	"log"
	"time"
)

func Listener() {
	dispatchedTransactionChannel := channels.GetDispatchedTransactionChannel()

	for {
		select {
		case dispatchedTransaction := <-dispatchedTransactionChannel:
			go waitForTransactionConfirmation(dispatchedTransaction)
		}
	}
}

func waitForTransactionConfirmation(dispatchedTransaction types.DispatchedTransaction) {
	time.Sleep(1 * time.Second)
	log.Printf("Transaction validated: %v", dispatchedTransaction.TransactionHash)
}
