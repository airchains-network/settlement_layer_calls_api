package chain

import (
	"context"
	"github.com/airchains-network/airsettle/x/airsettle/types"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"fmt"
)

// chain.VerifyBatch(batchNumber, chainId, merkleRootHash, prevMerkleRoot, zkProof, client, ctx, account, addr, sAPI)

func VerifyBatch(batchNumber uint64, chainId string, merkleRootHash string, prevMerkleRoot string, zkProof string, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, sAPI string ) (status bool,data string, error string) {

	fmt.Println("batchNumber: ", batchNumber)
	fmt.Println("chainId: ", chainId)
	fmt.Println("merkleRootHash: ", merkleRootHash)
	fmt.Println("prevMerkleRoot: ", prevMerkleRoot)
	fmt.Println("zkProof: ", zkProof)
	

	msg := &types.MsgVerifyBatch{
		Creator:     addr,
		BatchNumber: batchNumber,
		ChainId:     chainId,
		MerkleRootHash: merkleRootHash,
		PrevMerkleRoot: prevMerkleRoot,
		ZkProof:      zkProof,
	}

	txResp, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		error_msg := formatErrorMessage(err)
		return false,"error in transaction",error_msg
	}

	data = fmt.Sprintf(`{"txDetails":"%s"}`, txResp)
	return true, data, "verify batch successfully"
}