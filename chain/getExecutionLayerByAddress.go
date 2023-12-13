package chain

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type ExecutionLayerTrueResponseBody struct {
	ExeLayer struct {
		Validator             []string `json:"validator"`
		VotingPower           []string `json:"votingPower"`
		LatestBatch           string   `json:"latestBatch"`
		LatestMerkleRootHash  string   `json:"latestMerkleRootHash"`
		VerificationKey       string   `json:"verificationKey"`
		ChainInfo             string   `json:"chainInfo"`
		ID                    string   `json:"id"`
		Creator               string   `json:"creator"`
	} `json:"exelayer"`
}

type ExecutionLayerErrorResponseBody struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

func GetExecutionLayerByAddress(address string, sAPI string) (success bool, chainId string){

	apiURL := sAPI+"/airchains-network/airsettle/airsettle/show_execution_layer_by_address/"+address

	// Make the GET request
	response, err := http.Get(apiURL)
	if err != nil {
		return false , ""
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, "Error in Requesting to Execution Layer Blockchain API"
	}

	fmt.Println(string(body))

	// Check the structure of the response body to determine the appropriate struct
	var executionLayerResponse ExecutionLayerTrueResponseBody
	if err := json.Unmarshal(body, &executionLayerResponse); err == nil {
		if len(executionLayerResponse.ExeLayer.ID) == 0 {
			return false, ""
		}else{
			return true , string(body)
		}
	}

	// code may not reach here... but just in case
	var executionLayerErrResponse ExecutionLayerErrorResponseBody
	if err := json.Unmarshal(body, &executionLayerErrResponse); err == nil {
		// Successfully unmarshaled into ExecutionLayerErrorResponseBody
		return false, ""
	}
	
	// if not both data type
	return false, ""
}