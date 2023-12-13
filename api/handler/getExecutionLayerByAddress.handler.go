package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/airchains-studio/settlement_layer_calls_api/chain"
	"github.com/airchains-studio/settlement_layer_calls_api/model"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	cosmosclient "github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/syndtr/goleveldb/leveldb"
)

// HandlePostAPI handles the POST request
func HandleGetExecutionLayerByAddress(c *gin.Context, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, dbIPaddress *leveldb.DB, sAPI string) {

	// Parse the request body into a struct
	var requestBody model.RequestBodyGetExecutionLayerByAddress
	if err := c.BindJSON(&requestBody); 
	err != nil {
		respondWithError(c, "Invalid JSON format")
		return
	}

	// check if address is empty
	Address := requestBody.Address // "air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t"
	if len(Address) == 0 {
		respondWithError(c, "VerificationKey cannot be empty")
		return
	}

	success, chainDetails := chain.GetExecutionLayerByAddress(Address,sAPI)
	if success {
		respondWithSuccess(c, chainDetails , "get execution layer successfully")
		return
	}else{
		respondWithError(c, "chain id not found in this address")
		return
	}

}