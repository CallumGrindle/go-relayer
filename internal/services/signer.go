package services

import (
	"api/internal/channels"
	"api/internal/types"
	"log"
)

func Signer() {
	transactionRequestChannel := channels.GetTransactionRequestChannel()
	signedTransactionChannel := channels.GetSignedTransactionChannel()

	keyStore, err := newKeyStore()

	if err != nil {
		log.Fatal("Failed to initialize key store")
	}

	for {
		select {
		case transactionRequest := <-transactionRequestChannel:
			go signTransaction(transactionRequest, signedTransactionChannel, keyStore)
		}
	}
}

func signTransaction(transactionForSigning types.TransactionForSigning, signedTransactionChannel chan types.SignedTransaction, ks *KeyStore) {
	defer close(transactionForSigning.ResponseChannel)
	defer close(transactionForSigning.ErrorChannel)

	log.Printf("Signing transaction: %v", transactionForSigning.ExecutePayload)

	address, privateKey, nonce, err := ks.getRandomSigningKey()

	if err != nil {
		transactionForSigning.ErrorChannel <- err
		log.Print(err)
		return
	}

	log.Printf("pub key: %v", address)
	log.Printf("priv key: %v", privateKey)
	log.Printf("nonce: %v", nonce)

	signature := "0x123"
	transactionHash := "0x123"

	transactionForSigning.ResponseChannel <- types.ExecuteResponse{TransactionHash: transactionHash}
	signedTransactionChannel <- types.SignedTransaction{SignedTransaction: signature, TransactionHash: transactionHash}
}
