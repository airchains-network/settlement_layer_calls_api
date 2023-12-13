package chain

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func GetExecutionLayerById(id string, sAPI string) (success bool, chainId string){

	apiURL := sAPI+"/airchains-network/airsettle/airsettle/show_execution_layer_by_id/"+id

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