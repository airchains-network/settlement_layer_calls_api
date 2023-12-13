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
func HandleGetExecutionLayerById(c *gin.Context, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, dbIPaddress *leveldb.DB, sAPI string) {

	// Parse the request body into a struct
	var requestBody model.RequestBodyGetExecutionLayerById
	if err := c.BindJSON(&requestBody); 
	err != nil {
		respondWithError(c, "Invalid JSON format")
		return
	}

	// check if chain id  is empty
	chainId := requestBody.ChainId
	if len(chainId) == 0 {
		respondWithError(c, "VerificationKey cannot be empty")
		return
	}

	success, chainDetails := chain.GetExecutionLayerById(chainId,sAPI)

	if success {
		respondWithSuccess(c, chainDetails , "get execution layer successfully")
		return
	}else{
		respondWithError(c, "chain id not found in this address")
		return
	}

}