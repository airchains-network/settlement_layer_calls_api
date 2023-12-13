package api

import (
	"github.com/airchains-studio/settlement_layer_calls_api/api/handler"
	// "log"
	// "net/http"
	"fmt"
    "github.com/gin-gonic/gin"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"context"
	"sync"
	"github.com/syndtr/goleveldb/leveldb"
)

func StartAPI(wg *sync.WaitGroup, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, db *leveldb.DB,sAPI string  ) {
	defer wg.Done() 

	// Create a new router
    router := gin.Default()

	// define port for api
	port := ":8080"
	
	// * Register the Handlers / Routers

	// add execution layer
	router.POST("/addexelayer", func(c *gin.Context) {
		handler.HandlePostAddExecutionLayer(c,client,ctx, account,addr,db, sAPI)
	})
	
	// get execution layer by address
	router.GET("/getexelayer_by_address", func(c *gin.Context) {
		handler.HandleGetExecutionLayerByAddress(c,client,ctx, account,addr,db, sAPI)
	})
	
	// get execution layer by chain id
	router.GET("/getexelayer_by_id", func(c *gin.Context) {
		handler.HandleGetExecutionLayerById(c,client,ctx, account,addr,db, sAPI)
	})

	// get all execution layers
	router.GET("/get_all_exelayer", func(c *gin.Context) {
		handler.HandleGetAllExecutionLayers(c,client,ctx, account,addr,db, sAPI)
	})

	
	// Run the server on given port
	if err := router.Run(port); err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
