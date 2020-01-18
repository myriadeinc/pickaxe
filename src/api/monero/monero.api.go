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
	Blocktemplate_blob            *string  `json:"blocktemplate_blob"`
	Blockhashing_blob            *string  `json:"blockhashing_blob"`
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
		return getGenesisBlockTemplate()
	}
	return jobTemplateResponse
}

func getGenesisBlockTemplate() (*JobTemplateResponse) {
	var Blob 			string = "0000000000000000000000000000000000000000000000000000000000000000"
	var ReservedOffset  uint64 = 0
	var Difficulty      uint64 = 1
	var Height          uint64 = 0
	var ExpectedReward  uint64 = 0
	var SeedHash		string = "0000000000000000000000000000000000000000000000000000000000000000"
	var PrevHash        string = "0000000000000000000000000000000000000000000000000000000000000000"
	return &JobTemplateResponse{
		Blocktemplate_blob : &Blob,
		Blockhashing_blob  : &Blob,
		ReservedOffset  : &ReservedOffset,
		Difficulty      : &Difficulty,
		Height          : &Height,
		ExpectedReward  : &ExpectedReward,
		SeedHash		: &SeedHash,
		PrevHash        : &PrevHash,
	}

}