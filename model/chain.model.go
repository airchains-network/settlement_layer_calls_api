package model

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

type VerificationKeyResponseBody struct {
	Vkey string   `json:"vkey"`
}

// possible response body
// {"exelayer":[],"pagination":{"next_key":null,"total":"0"}}
// {"exelayer":{"validator":["air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t"],"votingPower":["100"],"latestBatch":"0","latestMerkleRootHash":"0","verificationKey":"/verificationKey/fddd253b-8097-431b-bb94-a158fd51fdd8/","chainInfo":"some information about the chain, e.g. its made for DeFi","id":"fddd253b-8097-431b-bb94-a158fd51fdd8","creator":"air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t"}}
// {"exelayer":[{"validator":["air1r34mk94d74523l6yqwp32dg5durs0vz6mxscds"],"votingPower":["100"],"latestBatch":"0","latestMerkleRootHash":"0","verificationKey":"/verificationKey/65304c71-c05a-4ea7-a1d2-b900a95b778b/","chainInfo":"verificationkey","id":"65304c71-c05a-4ea7-a1d2-b900a95b778b","creator":"air1r34mk94d74523l6yqwp32dg5durs0vz6mxscds"},{"validator":["air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t"],"votingPower":["100"],"latestBatch":"0","latestMerkleRootHash":"0","verificationKey":"/verificationKey/fddd253b-8097-431b-bb94-a158fd51fdd8/","chainInfo":"some information about the chain, e.g. its made for DeFi","id":"fddd253b-8097-431b-bb94-a158fd51fdd8","creator":"air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t"}],"pagination":{"next_key":null,"total":"2"}}
type AllExecutionLayersResponseBody struct {
	ExeLayer []struct {
		Validator            []string `json:"validator"`
		VotingPower          []string `json:"votingPower"`
		LatestBatch          string   `json:"latestBatch"`
		LatestMerkleRootHash string   `json:"latestMerkleRootHash"`
		VerificationKey      string   `json:"verificationKey"`
		ChainInfo            string   `json:"chainInfo"`
		ID                   string   `json:"id"`
		Creator              string   `json:"creator"`
	} `json:"exelayer"`
	Pagination struct {
		NextKey string `json:"next_key"`
		Total   string `json:"total"`
	} `json:"pagination"`
}


// {"batch":{"batchNumber":"1","chainId":"f0722463-03f1-485e-8d91-f592cad02d23","prevMerkleRootHash":"","merkleRootHash":"0xMerkleRoot","zkProof":"zkproof","inputs":["input1","input2"],"verified":"true","batchSubmitter":"air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t","batchVerifier":"air15nt3l400td56dhvy7tk4pehv2rqu2fw53fw59t"}}
type BatchResponseBody struct {
	Batch struct {
		BatchNumber    string   `json:"batchNumber"`
		ChainId        string   `json:"chainId"`
		PrevMerkleRootHash string   `json:"prevMerkleRootHash"`
		MerkleRootHash string   `json:"merkleRootHash"`
		ZkProof        string   `json:"zkProof"`
		Inputs         []string `json:"inputs"`
		Verified       string   `json:"verified"`
		BatchSubmitter string   `json:"batchSubmitter"`
		BatchVerifier  string   `json:"batchVerifier"`
	} `json:"batch"`
}

