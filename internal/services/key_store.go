package services

import (
	"api/internal/config"
	"context"
	"errors"
	"log"
	"math/big"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type SigningKey struct {
	nonce      *big.Int
	privateKey string
}

type KeyStore struct {
	mu          sync.Mutex
	signingKeys map[string]*SigningKey
}

func newKeyStore() (*KeyStore, error) {
	keyStore := &KeyStore{
		signingKeys: make(map[string]*SigningKey),
	}

	keyStore.initialize()

	return keyStore, nil
}

func (ks *KeyStore) initialize() error {
	ks.mu.Lock()
	defer ks.mu.Unlock()

	for _, privateKeyHex := range config.ApplicationConfig.PrivateKeys {
		privateKeyHex = strings.TrimSpace(privateKeyHex)
		privateKeyBytes, err := crypto.HexToECDSA(privateKeyHex)

		if err != nil {
			log.Printf("Failed to decode private key hex: %v", err)
			continue
		}

		publicAddress := crypto.PubkeyToAddress(privateKeyBytes.PublicKey).Hex()

		nonce, error := fetchNonceFromBlockchain(publicAddress)

		if error != nil {
			log.Printf("Error loading nonce for key: %v", publicAddress)
		}

		ks.signingKeys[publicAddress] = &SigningKey{
			nonce:      nonce,
			privateKey: privateKeyHex,
		}
	}

	return nil
}

func fetchNonceFromBlockchain(address string) (*big.Int, error) {
	client := getBlockchainClient()

	if !common.IsHexAddress(address) {
		return nil, errors.New("invalid Ethereum address: " + address)
	}

	commonAddress := common.HexToAddress(address)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	nonce, err := client.PendingNonceAt(ctx, commonAddress)

	if err != nil {
		return nil, err
	}

	return big.NewInt(int64(nonce)), nil
}

func (ks *KeyStore) getRandomSigningKey() (address string, privateKey string, nonce *big.Int, err error) {
	keys := make([]string, 0, len(ks.signingKeys))
	for key := range ks.signingKeys {
		keys = append(keys, key)
	}

	if len(keys) == 0 {
		return "", "", nil, errors.New("Error: no available signing keys")
	}

	address = keys[rand.Intn(len(keys))]
	signingKey, exists := ks.signingKeys[address]

	if !exists {
		return "", "", nil, errors.New("unexpected error: key not found in store: " + address)
	}

	nonce, err = ks.getAndIncrementNonce(address)

	if err != nil {
		return "", "", nil, err
	}

	privateKey = signingKey.privateKey

	return address, signingKey.privateKey, signingKey.nonce, nil
}

func (ks *KeyStore) getAndIncrementNonce(address string) (*big.Int, error) {
	ks.mu.Lock()
	defer ks.mu.Unlock()

	if ks.signingKeys[address].nonce == nil {
		var err error
		ks.signingKeys[address].nonce, err = fetchNonceFromBlockchain(address)

		if err != nil {
			return nil, err
		}
	}

	nonce := new(big.Int).Set(ks.signingKeys[address].nonce)

	ks.signingKeys[address].nonce.Add(ks.signingKeys[address].nonce, big.NewInt(1))

	return nonce, nil
}
