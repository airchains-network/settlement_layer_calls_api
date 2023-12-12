package connect

import (
	"fmt"
	"context"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/airchains-studio/settlement_layer_calls_api/chain"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
)

func SettlementLayer() ( cosmosclient.Client, cosmosaccount.Account, string , context.Context ){

	// connect to blockchain (settlement layer)
	accountName := "admin"
	accountPath := "./accounts"
	ctx := context.Background()
	gasLimit := "300000"
	addressPrefix := "air" // "cosmos"
	sRPC := "http://localhost:26657"

	client, err := cosmosclient.New(ctx, cosmosclient.WithGas(gasLimit), cosmosclient.WithAddressPrefix(addressPrefix), cosmosclient.WithNodeAddress(sRPC), cosmosclient.WithKeyringDir(accountPath))
	if err != nil {
		panic(err)
	}

	// check if admin account exists
	isAccountExists,_ := chain.CheckIfAccountExists(accountName,client,addressPrefix,accountPath)

	// create admin account if don't exists
	if !isAccountExists {
		chain.CreateAccount(accountName, accountPath)

		// now recheck, its impossible its not exists. if error occure it will panic and stop.
		isAccountExists,_ = chain.CheckIfAccountExists(accountName,client,addressPrefix,accountPath)
	}

	account, err := client.Account(accountName)
	if err != nil {
		panic(err)
	}

	addr, err := account.Address(addressPrefix)
	if err != nil {
		panic(err)
	}
	fmt.Println("admin address:",addr)

	return client, account, addr , ctx
}