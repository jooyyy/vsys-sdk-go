package vsys

const (
	Funcidx2Supersede    = 0
	Funcidx2Issue	     = 1
	Funcidx2Destroy      = 2
	Funcidx2UpdateList   = 3
	Funcidx2Send     	 = 4
	Funcidx2Transfer     = 5
	Funcidx2Deposit      = 6
	Funcidx2Withdraw   	 = 7
	Funcidx2TotalSupply  = 8
	Funcidx2MaxSupply    = 9
	Funcidx2BalanceOf    = 10
	Funcidx2GetIssuer    = 11
)

type ContractV2 struct {
	Address string
	AddToList 	bool
}

// BuildSupersedeData transfer issuing right of Contract
func (c *ContractV2) BuildSupersedeData(newIssuer string, newRegulator string) []byte {
	data := DataEncoder{}
	data.EncodeArgAmount(2)
	data.Encode(newIssuer, DeTypeAddress)
	data.Encode(newRegulator, DeTypeAddress)
	return data.Result()
}

func (c *ContractV2) BuildUpdateListData(address string, add bool) []byte {
	data := DataEncoder{}
	data.EncodeArgAmount(2)
	data.Encode(address, DeTypeAddress)
	data.Encode(add, DeTypeBoolean)
	return data.Result()
}