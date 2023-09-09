package services

import (
	"api/internal/channels"
	"api/internal/types"
	"log"
	"time"
)

func Sender() {
	signedTransactionChannel := channels.GetSignedTransactionChannel()
	dispatchedTransactionChannel := channels.GetDispatchedTransactionChannel()

	for {
		select {
		case signedTransaction := <-signedTransactionChannel:
			go sendTransaction(signedTransaction, dispatchedTransactionChannel)
		}
	}
}

func sendTransaction(signedTransaction types.SignedTransaction, dispatchedTransactionChannel chan types.DispatchedTransaction) {
	time.Sleep(1 * time.Second)
	log.Printf("Got signed transaction: %v", signedTransaction.SignedTransaction)

	dispatchedTransactionChannel <- types.DispatchedTransaction{TransactionHash: signedTransaction.TransactionHash}
}
