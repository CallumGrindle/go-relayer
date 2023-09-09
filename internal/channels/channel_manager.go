package channels

import "api/internal/types"

var transactionRequestChannel chan types.TransactionForSigning

func GetTransactionRequestChannel() chan types.TransactionForSigning {
	if transactionRequestChannel == nil {
		transactionRequestChannel = make(chan types.TransactionForSigning)
	}

	return transactionRequestChannel
}
