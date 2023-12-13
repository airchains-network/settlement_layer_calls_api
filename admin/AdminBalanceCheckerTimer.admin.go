package admin 

import (
	"sync"
	"time"
	"context"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/airchains-studio/settlement_layer_calls_api/chain"
	"github.com/syndtr/goleveldb/leveldb"
	"strconv"
	"fmt"
)

func AdminBalanceCheckerTimer(wg *sync.WaitGroup, ctx context.Context,client cosmosclient.Client,account cosmosaccount.Account, addr string, db *leveldb.DB){
	defer wg.Done()
	_ = db 

	// hourly send token report to admin. eg. how many tokens are left in faucet wallet
	minuteCount := 0 // 60 minutes = 1 hour

	// check admin balance every minute
	for{
		// wait 60 seconds
		time.Sleep(5 * 60 * time.Second)

		// check client connection
		_, err := client.Status(ctx)
		if err != nil {
			InformAdmin("error","Blockchain is offline")
		}

		// check if admin have enough balance to send tokens
		isBalance, amount, _ := chain.CheckBalance(ctx, addr, client)


		// no balance in admin wallet
		if isBalance == false {
			InformAdmin("error","No balance in execution layer faucet/admin wallet")
		}

		// admin have less than 10 tokens (not enough balance)
		amountStr := strconv.FormatInt(amount, 10)
		if amount < 100 {
			InformAdmin("error","Admin dont have enough tokens: " + amountStr)
		}

		// send report: inform admin about his/her balance in faucet wallet.
		minuteCount++
		if minuteCount > 59 {
			minuteCount = 0
			// inform / send report to admin about currecnt balance
			InformAdmin("information","Admin have Currently" + amountStr + " tokens") 
		}

	}
}

func InformAdmin(message string, msgtype string){
	switch(msgtype){
		case "error":
			fmt.Println("Error: ",message)
		case "information":
			fmt.Println("Information: ",message)
		default: // default not possible in out case, but still just for safety
			fmt.Println("Error: ",message)
	}

	if msgtype == "error" {
		// after Error, wait for 1 hour. don't waste API in sending same error every minute
		time.Sleep(1 * time.Hour)
	}
}