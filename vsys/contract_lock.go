package vsys

import "time"

const (
	FuncidxLock    = 0
)

type LockContract struct {
	TokenId  	string
	Timestamp   int64
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

func (c *LockContract) DecodeLock(data []byte) {
	de := DataEncoder{}
	list := de.Decode(data)
	c.Timestamp = list[0].Value.(int64)
}
