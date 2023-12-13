package chain 

import (
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"fmt"
)

func CreateAccount(accountName string, accountPath string) {
	registry, err := cosmosaccount.New(cosmosaccount.WithHome(accountPath))
	if err != nil {
		fmt.Println(err)
		return
	}

	account, mnemonic, err := registry.Create(accountName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("new account created: ",account, mnemonic)
}