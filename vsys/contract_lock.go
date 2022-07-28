package vsys

import "time"

const (
	FuncidxLock = 0
)

const (
	LockStateVariableIdxMaker   = 0x00
	LockStateVariableIdxTokenId = 0x01

	LockStateMapIdxBalance  = 0x00
	LockStateMapIdxLockTime = 0x01
)

type LockContract struct {
	TokenId   string
	Timestamp int64
}

type LockResp struct {
	ContractId string `json:"contractId"`
	Key        string `json:"key"`
	Height     int64  `json:"height"`
	DbName     string `json:"dbName"`
	DateType   string `json:"dateType"`
}

type LockMakerResp struct {
	LockResp
	Value string `json:"value"`
}

type LockTokenIdResp struct {
	LockResp
	Value string `json:"value"`
}

type LockBalanceResp struct {
	LockResp
	Value int64 `json:"value"`
}

type LockTimeResp struct {
	LockResp
	Value int64 `json:"value"`
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

func (c *LockContract) DecodeRegister(data []byte) {
	de := DataEncoder{}
	list := de.Decode(data)
	c.TokenId = list[0].Value.(string)
}

func (c *LockContract) DecodeLock(data []byte) {
	de := DataEncoder{}
	list := de.Decode(data)
	c.Timestamp = list[0].Value.(int64)
}

func (c *LockContract) GenQueryKeyMaker() string {
	var b []byte
	b = append(b, LockStateVariableIdxMaker)
	return Base58Encode(b)
}

func (c *LockContract) GenQueryKeyTokenId() string {
	var b []byte
	b = append(b, LockStateVariableIdxTokenId)
	return Base58Encode(b)
}

func (c *LockContract) GenQueryKeyBalance(userAddress string) string {
	var b []byte
	b = append(b, LockStateMapIdxBalance)
	b = append(b, DeTypeAddress)
	b = append(b, Base58Decode(userAddress)...)
	return Base58Encode(b)
}

func (c *LockContract) GenQueryKeyLockTime(userAddress string) string {
	var b []byte
	b = append(b, LockStateMapIdxLockTime)
	b = append(b, DeTypeAddress)
	b = append(b, Base58Decode(userAddress)...)
	return Base58Encode(b)
}
