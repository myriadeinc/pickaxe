package MoneroApi

import (
	"fmt"
	"github.com/ybbus/jsonrpc"
	"github.com/myriadeinc/pickaxe/src/util/logger"
)

type Request struct {
	Wallet_Address 	string `json:"wallet_address"`
	Reserve_Size		uint16		 `json:"reserve_size"`
}

var debug bool = false

type JobTemplateResponse struct {
	Ok 					 	 bool		`json:"ok"`
	Difficulty     uint64  `json:"difficulty"`
	Height         uint64  `json:"height"`
	Blob           string `json:"blocktemplate_blob"`
	ReservedOffset int    `json:"reserved_offset"`
	PrevHash       string `json:"prev_hash"`
}

var rpcClient jsonrpc.RPCClient

func Init(url string, debugMode bool) () {
	debug = debugMode
	rpcClient = jsonrpc.NewClient(url)
}

func GetJobTemplate(reserveSize uint16, address string) (*JobTemplateResponse) {
	if (debug) {
		return fakeNewJobTemplate()
	}
	result, err := rpcClient.Call("getblocktemplate", &Request{address, reserveSize})
	var response *JobTemplateResponse
	if err != nil || result.Error != nil {
		response.Ok = false
		fmt.Println(`Error`)
		return response
		// error handling goes here e.g. network / http error
	}
	result.GetObject(&response)

	LoggerUtil.Logger.Info(response)
	return response
}

func GetFakeBlockHeight() uint64 {
	var number uint64 = 90
	return number
}


func fakeNewJobTemplate() (*JobTemplateResponse) {
	// fmt.Println("Querying monero API")
	return &JobTemplateResponse{
		Ok: true,		 	
		Difficulty: 10000,
		Height: 100,         
		Blob: "heyeheyeyeyhe",        
		ReservedOffset: 120,
		PrevHash: "c82411b6b6ac",       
	}
}