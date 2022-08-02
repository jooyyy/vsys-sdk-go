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
	SenderAddress string
	TokenIndex    int64
	Recipient     string
	ContractId    string
	Description   string
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

func BuildDepositData(sender string, contract string, index int32) []byte {
	data := DataEncoder{}
	data.EncodeArgAmount(3)
	data.Encode(sender, DeTypeAddress)
	data.Encode(contract, DeTypeContractAccount)
	data.Encode(index, DeTypeInt32)
	return data.Result()
}

func BuildWithdrawData(contract string, recipient string, index int32) []byte {
	data := DataEncoder{}
	data.EncodeArgAmount(3)
	data.Encode(contract, DeTypeContractAccount)
	data.Encode(recipient, DeTypeAddress)
	data.Encode(index, DeTypeInt32)
	return data.Result()
}

func (c *NFTContract) DecodeIssueData(data []byte) {
	de := DataEncoder{}
	list := de.Decode(data)
	c.Description = list[0].Value.(string)
}

func (c *NFTContract) DecodeSendData(data []byte) {
	de := DataEncoder{}
	list := de.Decode(data)
	c.Recipient = list[0].Value.(string)
	c.TokenIndex = list[1].Value.(int64)
}

func (c *NFTContract) DecodeTransferData(data []byte) {
	de := DataEncoder{}
	list := de.Decode(data)
	c.SenderAddress = list[0].Value.(string)
	c.Recipient = list[1].Value.(string)
	c.TokenIndex = list[2].Value.(int64)
}

func (c *NFTContract) DecodeDepositData(data []byte) {
	de := DataEncoder{}
	list := de.Decode(data)
	c.SenderAddress = list[0].Value.(string)
	c.ContractId = list[1].Value.(string)
	c.TokenIndex = list[2].Value.(int64)
}

func (c *NFTContract) DecodeWithdrawData(data []byte) {
	de := DataEncoder{}
	list := de.Decode(data)
	c.ContractId = list[0].Value.(string)
	c.Recipient = list[1].Value.(string)
	c.TokenIndex = list[2].Value.(int64)
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
