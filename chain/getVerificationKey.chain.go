package chain

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/airchains-studio/settlement_layer_calls_api/model"
)

func GetVerificationKeyById(id string, sAPI string) (success bool, chainId string){

	apiURL := sAPI+"/airchains-network/airsettle/airsettle/verification_key/"+id

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
	var verificationKeyResponseBody model.VerificationKeyResponseBody
	if err := json.Unmarshal(body, &verificationKeyResponseBody); err == nil {
		if len(verificationKeyResponseBody.Vkey) == 0 {
			return false, ""
		}else{
			return true , string(body)
		}
	}

	// if not both data type
	return false, ""
}