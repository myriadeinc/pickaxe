package MoneroApi

import (
	"github.com/ybbus/jsonrpc"
	"github.com/myriadeinc/pickaxe/src/util/logger"
)

type Request struct {
	Wallet_Address 		string `json:"wallet_address"`
	Reserve_Size		int		`json:"reserve_size"`
}

var debug bool = false

type JobTemplateResponse struct {
	Blob            *string  `json:"blocktemplate_blob"`
	ReservedOffset  *uint64  `json:"reserved_offset"`
	Difficulty      *uint64  `json:"difficulty"`
	Height          *uint64  `json:"height"`
	ExpectedReward  *uint64  `json:"expected_reward"`
	SeedHash		*string  `json:"seed_hash"`
	PrevHash        *string  `json:"prev_hash"`
}

var rpcClient jsonrpc.RPCClient

func Init(url string, debugMode bool) () {
	debug = debugMode
	rpcClient = jsonrpc.NewClient(url)
}

func GetJobTemplate(reserveSize int, address string) (*JobTemplateResponse) {
    var jobTemplateResponse *JobTemplateResponse
    err := rpcClient.CallFor(&jobTemplateResponse, "get_block_template", &Request{address, reserveSize})
    if err != nil || jobTemplateResponse == nil {
		LoggerUtil.Logger.Error(err.Error())
		panic("Bad response from Monero Node! Shutting down")
	}
	return jobTemplateResponse
}
