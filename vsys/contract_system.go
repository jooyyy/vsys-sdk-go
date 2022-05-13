package vsys

const (
	FuncIdxSysSend     = 0
	FuncIdxSysDeposit  = 1
	FuncIdxSysWithdraw = 2
	FuncIdxSysTransfer = 3
)

type SystemContract struct {
	Sender     string
	ContractId string
	Recipient  string
	Amount     int64
}

func (c *SystemContract) BuildSendData(recipient string, amount int64) []byte {
	data := DataEncoder{}
	data.EncodeArgAmount(2)
	data.Encode(recipient, DeTypeAddress)
	data.Encode(amount, DeTypeAmount)
	return data.Result()
}

func (c *SystemContract) BuildDepositData(sender string, contractId string, amount int64) []byte {
	data := DataEncoder{}
	data.EncodeArgAmount(3)
	data.Encode(sender, DeTypeAddress)
	data.Encode(contractId, DeTypeContractAccount)
	data.Encode(amount, DeTypeAmount)
	return data.Result()
}

func (c *SystemContract) BuildWithdrawData(contractId string, recipient string, amount int64) []byte {
	data := DataEncoder{}
	data.EncodeArgAmount(3)
	data.Encode(contractId, DeTypeContractAccount)
	data.Encode(recipient, DeTypeAddress)
	data.Encode(amount, DeTypeAmount)
	return data.Result()
}

func (c *SystemContract) BuildTransferData(sender string, recipient string, amount int64) []byte {
	data := DataEncoder{}
	data.EncodeArgAmount(3)
	data.Encode(sender, DeTypeAddress)
	data.Encode(recipient, DeTypeAddress)
	data.Encode(amount, DeTypeAmount)
	return data.Result()
}

func (c *SystemContract) DecodeSend(data []byte) {
	de := DataEncoder{}
	list := de.Decode(data)
	c.Recipient = list[0].Value.(string)
	c.Amount = list[1].Value.(int64)
}

func (c *SystemContract) DecodeDeposit(data []byte) {
	de := DataEncoder{}
	list := de.Decode(data)
	c.Sender = list[0].Value.(string)
	c.ContractId = list[1].Value.(string)
	c.Amount = list[2].Value.(int64)
}

func (c *SystemContract) DecodeWithdraw(data []byte) {
	de := DataEncoder{}
	list := de.Decode(data)
	c.ContractId = list[0].Value.(string)
	c.Recipient = list[1].Value.(string)
	c.Amount = list[2].Value.(int64)
}

func (c *SystemContract) DecodeTransfer(data []byte) {
	de := DataEncoder{}
	list := de.Decode(data)
	c.Sender = list[0].Value.(string)
	c.Recipient = list[1].Value.(string)
	c.Amount = list[1].Value.(int64)
}
