package types

type ExecutePayload struct {
	Address     string `json:"address"`
	Transaction struct {
		Abi       string `json:"abi"`
		Nonce     uint   `json:"nonce"`
		Signature string `json:"signature"`
	}
}

type ExecuteResponse struct {
	TransactionHash string `json:"transactionHash"`
}

type TransactionForSigning struct {
	Request         ExecutePayload
	ResponseChannel chan ExecuteResponse
}
