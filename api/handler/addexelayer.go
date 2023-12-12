package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/airchains-studio/settlement_layer_calls_api/chain"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	cosmosclient "github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/syndtr/goleveldb/leveldb"
)

// RequestBody is the structure for the incoming JSON request
type RequestBody struct {
	VerificationKey string `json:"verification_key"`
	ChainInfo       string `json:"chain_info"`
}

// ResponseBody is the structure for the JSON response
type ResponseBody struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Message string `json:"message"`
	TxHash  string `json:"txhash"`
}

// HandlePostAPI handles the POST request
func HandleAddExecutionLayerPostAPI(w http.ResponseWriter, r *http.Request, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, dbIPaddress *leveldb.DB) {

	// Parse the incoming JSON request
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		respondWithError(w, http.StatusOK, "Invalid JSON format")
		return
	}

	// Validate the verification_key (add your validation logic here)
	if len(requestBody.VerificationKey) == 0 {
		respondWithError(w, http.StatusOK, "VerificationKey cannot be empty")
		return
	}

	// Validate the chain_info (add your validation logic here)
	if len(requestBody.ChainInfo) == 0 {
		respondWithError(w, http.StatusOK, "ChainInfo cannot be empty")
		return
	}

	verificationKey := requestBody.VerificationKey
	chainInfo := requestBody.ChainInfo

	success, err, message, txhash := chain.AddExecutionLayer(verificationKey, chainInfo, client, ctx, account, addr)

	response := ResponseBody{
		Success: success,
		Error:   formatErrorMessage(err),
		Message: message,
		TxHash:  txhash,
	}

	respondWithJSON(w, http.StatusOK, response)
	return

}

func formatErrorMessage(err error) string {
    if err == nil {
        return "No error"
    }
    return fmt.Sprintf("Failed to add execution layer: %v", err.Error())
}

// respondWithError sends a JSON error response
func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	response := ResponseBody{
		Success: false,
		Error:   message,
		Message: message,
		TxHash:  "",
	}
	respondWithJSON(w, statusCode, response)
}

// respondWithJSON sends a JSON response
func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func storeData(db *leveldb.DB, key, value string) error {
	// Store the key-value pair in LevelDB
	err := db.Put([]byte(key), []byte(value), nil)
	if err != nil {
		return err
	}
	return nil
}

func retrieveData(db *leveldb.DB, key string) (string, error) {
	// Retrieve the value for the given key from LevelDB
	data, err := db.Get([]byte(key), nil)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
