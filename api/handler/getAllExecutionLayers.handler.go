package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/airchains-studio/settlement_layer_calls_api/chain"
	// "github.com/airchains-studio/settlement_layer_calls_api/model"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	cosmosclient "github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/syndtr/goleveldb/leveldb"
)

// HandlePostAPI handles the POST request
func HandleGetAllExecutionLayers(c *gin.Context, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, dbIPaddress *leveldb.DB, sAPI string) {

	success, chainDetails := chain.GetAllExecutionLayers(sAPI)
	if success {
		respondWithSuccess(c, chainDetails , "get execution layer successfully")
		return
	}else{
		respondWithError(c, "error in getting execution layers list")
		return
	}

}