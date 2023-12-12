package chain

import (
	"context"
	"github.com/airchains-network/airsettle/x/airsettle/types"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
)

func AddExecutionLayer(verificationKey string,chainInfo string , client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string ) (success bool,err error, message string ,txhash string) {

	msg := &types.MsgAddExecutionLayer{
		Creator:         addr,
		VerificationKey: verificationKey,
		ChainInfo:       chainInfo,
	}

	// http://0.0.0.0:1317/cosmos/tx/v1beta1/txs/{transaction_hash}
	txResp, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		return false,err,"error in transaction",""
	}

	return true,nil,"Success",txResp.TxHash
}