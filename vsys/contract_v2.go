package vsys

type ContractV2 struct {
	NewIssuer string
	NewRegulator string
	Address string
	AddToList bool
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

func (c *ContractV2) DecodeSupersede(data []byte) {
	de := DataEncoder{}
	list := de.Decode(data)
	c.NewIssuer = list[0].Value.(string)
	c.NewRegulator = list[1].Value.(string)
}

func (c *ContractV2) DecodeUpdateList(data []byte) {
	de := DataEncoder{}
	list := de.Decode(data)
	c.Address = list[0].Value.(string)
	c.AddToList = list[1].Value.(bool)
}
