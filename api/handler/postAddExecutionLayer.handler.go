package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/airchains-studio/settlement_layer_calls_api/chain"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	cosmosclient "github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/airchains-studio/settlement_layer_calls_api/model"

)

// HandlePostAPI handles the POST request
func HandlePostAddExecutionLayer(c *gin.Context, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, dbIPaddress *leveldb.DB, sAPI string) {

	// Parse the request body into a struct
	var requestBody model.RequestBodyAddExecutionLayer
	if err := c.BindJSON(&requestBody); 
	err != nil {
		respondWithError(c, "Invalid JSON format")
		return
	}

	verificationKey := requestBody.VerificationKey
	chainInfo := requestBody.ChainInfo
	
	// Validate the verification_key (add your validation logic here)
	if len(verificationKey) == 0 {
		respondWithError(c, "VerificationKey cannot be empty")
		return
	}

	// Validate the chain_info (add your validation logic here)
	if len(chainInfo) == 0 {
		respondWithError(c, "ChainInfo cannot be empty")
		return
	}

	success, data, error_msg:= chain.AddExecutionLayer(verificationKey, chainInfo, client, ctx, account, addr, sAPI)
	if !success {
		respondWithError(c, error_msg)
		return
	}

	respondWithSuccess(c, data , "add execution layer successfully")
	return
}