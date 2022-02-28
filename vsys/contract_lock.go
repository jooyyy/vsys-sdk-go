package vsys

import "time"

type LockContract struct {
	TokenId  	string
}

func (c *LockContract) BuildRegisterData() []byte {
	data := DataEncoder{}
	data.EncodeArgAmount(1)
	data.Encode(c.TokenId, DeTypeTokenId)
	return data.result
}

func (c *LockContract) BuildLockData(timestamp time.Time) []byte {
	data := DataEncoder{}
	data.EncodeArgAmount(1)
	data.Encode(timestamp.UnixNano(), DeTypeTimestamp)
	return data.result
}
