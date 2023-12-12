package chain 

import (
	"fmt"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/cosmos/cosmos-sdk/types/query"
	"context"
)


func CheckBalance(ctx context.Context, accountAddress string, client cosmosclient.Client) (bool,int64,error) {
	pageRequest := &query.PageRequest{} // Add this line to create a new PageRequest

	balances, err := client.BankBalances(ctx, accountAddress, pageRequest) // Add pageRequest as the third argument
	if err != nil {
		fmt.Printf("Error querying bank balances: %v\n", err)
		return false,0,err
	}

	for _, balance := range balances {
		if balance.Denom == "token" {
			return true,balance.Amount.Int64(), nil
		}
	}
	
	// no stake found
	return false,0, nil
}