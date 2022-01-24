package vsys

const (
	NFTFuncIdxSupersede = 0
	NFTFuncIdxIssue     = 1
	NFTFuncIdxSend      = 2
	NFTFuncIdxTransfer  = 3
	NFTFuncIdxDeposit   = 4
	NFTFuncIdxWithdraw  = 5
)

type NFTIssueResponse struct {
	Type       int
	Id         string
	Fee        int64
	Timestamp  int64
	Attachment string
	FeeScale   int
	Recipient  string
}

type NFTContract struct {
	ContractId  string
	Description string
}

func BuildIssueData(tokenDescription string) []byte {
	data := DataEncoder{}
	data.EncodeArgAmount(1)
	data.Encode(tokenDescription, DeTypeShortText)
	return data.Result()
}

func BuildSendData(recipient string, nftIndex int32) []byte {
	data := DataEncoder{}
	data.EncodeArgAmount(2)
	data.Encode(recipient, DeTypeAddress)
	data.Encode(nftIndex, DeTypeInt32)
	return data.Result()
}

func (c *NFTContract) BuildRegisterData() []byte {
	data := DataEncoder{}
	data.EncodeArgAmount(0)
	return data.result
}

func (c *NFTContract) BuildIssue(nftDesc string) []byte {
	data := DataEncoder{}
	data.EncodeArgAmount(1)
	data.Encode(nftDesc, DeTypeShortText)
	return data.result
}

func (c *NFTContract) BuildSend(recipient string, index int32) []byte {
	data := DataEncoder{}
	data.EncodeArgAmount(2)
	data.Encode(recipient, DeTypeAddress)
	data.Encode(index, DeTypeInt32)
	return data.result
}