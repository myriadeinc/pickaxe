package monero_api

import (
	"github.com/ybbus/jsonrpc"
)

type Request struct {
	Wallet_Address 	string `json:"wallet_address"`
	Reserve_Size		int		 `json:"reserve_size"`
}

type JobTemplateResponse struct {
	Ok 					 	 bool
	Difficulty     int64  `json:"difficulty"`
	Height         int64  `json:"height"`
	Blob           string `json:"blocktemplate_blob"`
	ReservedOffset int    `json:"reserved_offset"`
	PrevHash       string `json:"prev_hash"`
}

var rpcClient := jsonrpc.NewClient("http://0.0.0.0:8080/rpc")

func GetJobTemplate(reserveSize int, address string) (*JobTemplateResponse) {
	result, err := rpcClient.Call("getblocktemplate", &Request{address, reserveSize})
	JobTemplateResponse* response;
	if err != nil || response.Error != nil {
		response.Ok = false
		return response
		// error handling goes here e.g. network / http error
	}
	response = &JobTemplateResponse(result)
	fmt.Println(response)
	return response
}