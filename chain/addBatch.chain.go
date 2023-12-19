package chain

import (
	"context"
	"fmt"
	"github.com/airchains-network/airsettle/x/airsettle/types"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
)

func AddBatch(batchNumber uint64, chainId string, witness string, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, sAPI string) (status bool, data string, error string) {

	msg := &types.MsgAddBatch{
		Creator:     addr,
		BatchNumber: batchNumber,
		ChainId:     chainId,
		Witness:      witness,
	}

	txResp, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		error_msg := formatErrorMessage(err)
		return false, "error in transaction", error_msg
	}

	data = fmt.Sprintf(`{"txDetails":"%s"}`, txResp)
	return true, data, "data sended successfully"
}
