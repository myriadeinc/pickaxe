package MoneroAPI

import (
	"fmt"
	"github.com/ybbus/jsonrpc"
)

type Request struct {
	Wallet_Address 	string `json:"wallet_address"`
	Reserve_Size		int		 `json:"reserve_size"`
}

type JobTemplateResponse struct {
	Ok 					 	 bool		`json:ok`
	Difficulty     int64  `json:"difficulty"`
	Height         int64  `json:"height"`
	Blob           string `json:"blocktemplate_blob"`
	ReservedOffset int    `json:"reserved_offset"`
	PrevHash       string `json:"prev_hash"`
}

func Init(url string) (jsonrpc.RPCClient) {
	return jsonrpc.NewClient(url)
}

func GetJobTemplate(rpcClient jsonrpc.RPCClient, reserveSize int, address string) (*JobTemplateResponse) {
	result, err := rpcClient.Call("getblocktemplate", &Request{address, reserveSize})
	var response *JobTemplateResponse
	if err != nil || result.Error != nil {
		response.Ok = false
		fmt.Println(`Error`)
		return response
		// error handling goes here e.g. network / http error
	}
	result.GetObject(&response)
	fmt.Println(response)
	return response
}