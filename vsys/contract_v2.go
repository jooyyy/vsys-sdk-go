package vsys

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

func (c *ContractV2) BuildUpdateListData() []byte {
	data := DataEncoder{}
	data.EncodeArgAmount(2)
	data.Encode(c.Address, DeTypeAddress)
	data.Encode(c.AddToList, DeTypeBoolean)
	return data.Result()
}