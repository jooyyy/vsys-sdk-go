package vsys

type LockContract struct {
	TokenId  	string
}

func (c *LockContract) BuildRegisterData() []byte {
	data := DataEncoder{}
	data.EncodeArgAmount(1)
	data.Encode(c.TokenId, DeTypeTokenId)
	return data.result
}