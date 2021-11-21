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

func (c *ContractV2) BuildUpdateListData(address string, add bool) []byte {
	data := DataEncoder{}
	data.EncodeArgAmount(2)
	data.Encode(address, DeTypeAddress)
	data.Encode(add, DeTypeBoolean)
	return data.Result()
}