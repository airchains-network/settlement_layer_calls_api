package chain

import (
	"context"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	cosmosclient "github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/airchains-studio/settlement_layer_calls_api/model"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/airchains-network/airsettle/x/airsettle/types"
)

func DeleteExecutionLayer(client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, addr string, sAPI string) (success bool, data string, error_msg string){
	
	// check if there is a chain id under this account
	apiURL := sAPI+"/airchains-network/airsettle/airsettle/show_execution_layer_by_address/"+addr

	// Make the GET request
	response, err := http.Get(apiURL)
	if err != nil {
		return false , "", "Error in Requesting to Execution Layer Blockchain API"
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, "", "Error in Requesting to Execution Layer Blockchain API"
	}

	// Check the structure of the response body to determine the appropriate struct
	var executionLayerResponse model.ExecutionLayerTrueResponseBody
	err = json.Unmarshal(body, &executionLayerResponse);
	
	if err != nil {
		return false , "", "Error in Requesting to Execution Layer Blockchain API"	
	}

	if (executionLayerResponse.ExeLayer.LatestBatch != "0") {
		return false, "", "Cannot delete a chain with batch number greater than 0"
	}


	// delete the execution layer associated with this address
	msg := &types.MsgDeleteExecutionLayer{
		Creator:         addr,
	}

	txResp, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		error_msg := formatErrorMessage(err)
		return false,"error in transaction",error_msg
	}

	data = fmt.Sprintf(`{"txDetails":"%s"}`, txResp)
	return true, data, "execution layer deleted successfully"

}