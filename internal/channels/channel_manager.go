package channels

import (
	"api/internal/types"
	"sync"
)

var transactionRequestChannel chan types.TransactionForSigning
var signedTransactionChannel chan types.SignedTransaction
var dispatchedTransactionChannel chan types.DispatchedTransaction

var onceTransactionRequestChannel sync.Once
var onceSignedTransactionChannel sync.Once
var onceDispatchedTransactionChannel sync.Once

func GetTransactionRequestChannel() chan types.TransactionForSigning {
	onceTransactionRequestChannel.Do(func() {
		transactionRequestChannel = make(chan types.TransactionForSigning)
	})

	return transactionRequestChannel
}

func GetSignedTransactionChannel() chan types.SignedTransaction {
	onceSignedTransactionChannel.Do(func() {
		signedTransactionChannel = make(chan types.SignedTransaction)
	})

	return signedTransactionChannel
}

func GetDispatchedTransactionChannel() chan types.DispatchedTransaction {
	onceDispatchedTransactionChannel.Do(func() {
		dispatchedTransactionChannel = make(chan types.DispatchedTransaction)
	})

	return dispatchedTransactionChannel
}
