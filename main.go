package main
import (
	"fmt"
	"github.com/airchains-studio/settlement_layer_calls_api/chain"
	"github.com/airchains-studio/settlement_layer_calls_api/api"
	"github.com/airchains-studio/settlement_layer_calls_api/admin"
	"github.com/airchains-studio/settlement_layer_calls_api/connect"
	"sync"

)

func main(){

	// connect to blockchain (settlement layer) && create admin wallet if do not exists
	client, account, addr , ctx, sAPI := connect.SettlementLayer()

	// connect to levelDB
	dbIPaddress := connect.LevelDB()

	// check admin balance
	isBalance, amount, err := chain.CheckBalance(ctx, addr, client)
	if err != nil {
		fmt.Println("Error in checking balance",err)
		return
	}

	// call faucet (if balance is 0)
	if !isBalance {
		fmt.Println("hey admin ! please restart api after getting tokens in this wallet:",addr)
		return
	} else{
		fmt.Println("Admin wallet have",amount, "tokens \nstarting api...")

		var wg sync.WaitGroup 
		wg.Add(2)
		go api.StartAPI(&wg, client,ctx, account,addr,dbIPaddress,sAPI)
		go admin.AdminBalanceCheckerTimer(&wg, ctx, client,account, addr,dbIPaddress)
	
		// Wait for both functions to finish
		wg.Wait()
	
		// wg Crashed. send report to admin
		fmt.Println("wg Crashed: faucet api stopped")
	}
}