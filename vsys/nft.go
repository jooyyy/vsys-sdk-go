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
	Recipient string
	Index int
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

func (c *NFTContract) BuildIssue() []byte {
	data := DataEncoder{}
	data.EncodeArgAmount(1)
	data.Encode(c.Description, DeTypeShortText)
	return data.result
}

func (c *NFTContract) BuildSend() []byte {
	data := DataEncoder{}
	data.EncodeArgAmount(2)
	data.Encode(c.Recipient, DeTypeAddress)
	data.Encode(c.Index, DeTypeInt32)
	return data.result
}