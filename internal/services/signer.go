package services

import (
	"api/internal/channels"
	"api/internal/config"
	"api/internal/types"
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func Signer() {
	transactionRequestChannel := channels.GetTransactionRequestChannel()
	signedTransactionChannel := channels.GetSignedTransactionChannel()

	for {
		select {
		case transactionRequest := <-transactionRequestChannel:
			go signTransaction(transactionRequest, signedTransactionChannel)
		}
	}
}

func signTransaction(transactionForSigning types.TransactionForSigning, signedTransactionChannel chan types.SignedTransaction) {
	defer close(transactionForSigning.ResponseChannel)
	defer close(transactionForSigning.ErrorChannel)

	log.Printf("Signing transaction: %v", transactionForSigning.Request.Address)

	privateKey, err := crypto.HexToECDSA(config.ApplicationConfig.PrivateKey)

	if err != nil {
		log.Printf("Error loading private key: %v", err)
		transactionForSigning.ErrorChannel <- err
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	log.Printf("pub key: %v", fromAddress)

	signature := "0x123"
	transactionHash := "0x123"

	transactionForSigning.ResponseChannel <- types.ExecuteResponse{TransactionHash: transactionHash}
	signedTransactionChannel <- types.SignedTransaction{SignedTransaction: signature, TransactionHash: transactionHash}

}
