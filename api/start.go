package api

import (
	"github.com/airchains-studio/settlement_layer_calls_api/api/handler"
	"log"
	"net/http"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"context"
	"sync"
	"github.com/syndtr/goleveldb/leveldb"
)

func StartAPI(wg *sync.WaitGroup, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, db *leveldb.DB ) {
	
	defer wg.Done() 

	// Register the handlers
	http.HandleFunc("/addexelayer", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleAddExecutionLayerPostAPI(w, r, client,ctx, account,addr,db)
	})
	
	// Start the server
	port := ":8080"
	log.Printf("Server started on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
