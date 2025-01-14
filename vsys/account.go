package vsys

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/curve25519"
)

type Account struct {
	publicKey  []byte
	privateKey []byte
	network    NetType
	accSeed    string
}

// get account address string
func (acc *Account) Address() string {
	return publicKeyToAddress(acc.publicKey, acc.network)
}

func publicKeyToAddress(publicKey []byte, network NetType) string {
	uAddr := make([]byte, 0)
	uAddr = append(uAddr, addrVersion, byte(network))
	uAddr = append(uAddr, HashChain(publicKey)[:20]...)
	return Base58Encode(append(uAddr, HashChain(uAddr)[:4]...))
}

// get account privateKey string
func (acc *Account) PrivateKey() string {
	return Base58Encode(acc.privateKey)
}

// get account publicKey string
func (acc *Account) PublicKey() string {
	return Base58Encode(acc.publicKey)
}

func (acc *Account) AccountSeed() string {
	return acc.accSeed
}

type AccountInfo struct {
	Address        string
	Regular        int64
	MintingAverage int64
	Available      int64
	Effective      int64
	Height         int64
}

// Get account V Coin balance
func (acc *Account) GetInfo() (AccountInfo, error) {
	path := fmt.Sprintf(ApiGetAddressDetail, acc.Address())
	resp, err := GetVsysApi().Get(path)
	if err != nil {
		return AccountInfo{}, err
	}

	var info AccountInfo
	err = json.Unmarshal(resp, &info)
	if err != nil {
		return AccountInfo{}, err
	}

	return info, nil
}

type TokenBalance struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
	Address string `json:"address"`
	TokenId string `json:"tokenId"`
	Balance int64  `json:"balance"`
	Unity   int64  `json:"unity"`
}

func (acc *Account) GetTokenBalance(tokenId string) (TokenBalance, error) {
	path := fmt.Sprintf(ApiContractTokenBalance, acc.Address(), tokenId)
	resp, err := GetVsysApi().Get(path)
	if err != nil {
		return TokenBalance{}, nil
	}

	var tBalance TokenBalance
	err = json.Unmarshal(resp, &tBalance)
	if err != nil {
		return TokenBalance{}, nil
	}

	return tBalance, nil
}

type TxHistoryList struct {
	Error        int    `json:"error"`
	Message      string `json:"message"`
	TotalCount   int64  `json:"totalCount"`
	Size         int    `json:"size"`
	Transactions []struct {
		Type      int64
		Id        string
		Fee       int64
		FeeScale  int64 `json:"feeScale"`
		Timestamp int64 `json:"timestamp"`
		Proofs    []struct {
			ProofType string `json:"proofType"`
			PublicKey string `json:"publicKey"`
			Address   string `json:"address"`
			Signature string `json:"signature"`
		}
		Recipient  string `json:"recipient"`
		Amount     int64  `json:"amount"`
		Attachment string `json:"attachment"`
		Status     string `json:"status"`
		FeeCharged int64  `json:"feeCharged"`
		Height     int64  `json:"height"`
	} `json:"transactions"`
}

func (acc *Account) GetTransferHistory(limit int64, offset int64, txType int64) (TxHistoryList, error) {
	return GetTransactionList(acc.Address(), limit, offset, txType)
}

// SignData sign data bytes and
// the output is base58 encoded data
func (acc *Account) SignData(data []byte) string {
	return Base58Encode(Sign(acc.privateKey, data, genRandomBytes(64)))
}

// VerifySignature check if signature is correct
func (acc *Account) VerifySignature(data, signature []byte) bool {
	return Verify(acc.publicKey, data, signature) == 1
}

// InitAccount return account with network initiated
func InitAccount(network NetType) *Account {
	return &Account{network: network}
}

// BuildFromPrivateKey build account using privateKey
func (acc *Account) BuildFromPrivateKey(privateKey string) {
	var bPrivateKey [32]byte
	var originPublicKey = new([32]byte)
	copy(bPrivateKey[:], Base58Decode(privateKey)[:])
	curve25519.ScalarBaseMult(originPublicKey, &bPrivateKey)
	acc.publicKey = originPublicKey[:]
	acc.privateKey = bPrivateKey[:]
}

// BuildFromPrivateKey build account using seed and nonce
func (acc *Account) BuildFromSeed(seed string, nonce int) {
	seedHash := BuildSeedHash(seed, nonce)
	keyPair := GenerateKeyPair(seedHash)
	acc.publicKey = keyPair.publicKey
	acc.privateKey = keyPair.privateKey
	acc.accSeed = seed
}

// BuildPayment build payment transaction
// recipient should be address
// amount is in minimum unit
// attachment can be empty
func (acc *Account) BuildPayment(recipient string, amount int64, attachment string) *Transaction {
	transaction := NewPaymentTransaction(recipient, amount, attachment)
	transaction.SenderPublicKey = acc.PublicKey()
	transaction.Signature = acc.SignData(transaction.BuildTxData())
	return transaction
}

// BuildLeasing build leasing transaction
// recipient should be address
// amount is in minimum unit
func (acc *Account) BuildLeasing(recipient string, amount int64) *Transaction {
	transaction := NewLeaseTransaction(recipient, amount)
	transaction.SenderPublicKey = acc.PublicKey()
	transaction.Signature = acc.SignData(transaction.BuildTxData())
	return transaction
}

// BuildCancelLeasing build Cancel transaction
func (acc *Account) BuildCancelLeasing(txId string) *Transaction {
	transaction := NewCancelLeaseTransaction(txId)
	transaction.SenderPublicKey = acc.PublicKey()
	transaction.Signature = acc.SignData(transaction.BuildTxData())
	return transaction
}

// BuildRegisterContract build RegisterContract transaction
func (acc *Account) BuildRegisterContract(contract string, max int64, unity int64, tokenDescription string, contractDescription string) *Transaction {
	c := &Contract{
		Max:              max * unity,
		Unity:            unity,
		TokenDescription: tokenDescription,
	}
	data := c.BuildRegisterData()
	transaction := NewRegisterTransaction(contract, Base58Encode(data), contractDescription)
	transaction.SenderPublicKey = acc.PublicKey()
	transaction.Signature = acc.SignData(transaction.BuildTxData())
	return transaction
}

// BuildRegisterContractV2 build RegisterContract transaction
func (acc *Account) BuildRegisterContractV2(contract string, max int64, unity int64, tokenDescription string, contractDescription string) *Transaction {
	c := &Contract{
		Max:              max * unity,
		Unity:            unity,
		TokenDescription: tokenDescription,
	}
	data := c.BuildRegisterData()
	transaction := NewRegisterTransaction(contract, Base58Encode(data), contractDescription)
	transaction.SenderPublicKey = acc.PublicKey()
	transaction.Signature = acc.SignData(transaction.BuildTxData())
	return transaction
}


// BuildExecuteContract build ExecuteContract transaction
func (acc *Account) BuildExecuteContract(contractId string, funcIdx int16, funcData []byte, attachment string) *Transaction {
	transaction := NewExecuteTransaction(contractId, funcIdx, Base58Encode(funcData), attachment)
	transaction.SenderPublicKey = acc.PublicKey()
	transaction.Signature = acc.SignData(transaction.BuildTxData())
	return transaction
}

// BuildSendTokenTransaction build SendToken transaction
func (acc *Account) BuildSendTokenTransaction(tokenId string, recipient string, amount int64, isSplitSupported bool, attachment string) *Transaction {
	a := &Contract{
		ContractId: TokenId2ContractId(tokenId),
		Amount:     amount,
		Recipient:  recipient,
	}
	funcData := a.BuildSendData()
	funcIdx := FuncidxSend
	if isSplitSupported {
		funcIdx = FuncidxSendSplit
	}
	transaction := NewExecuteTransaction(a.ContractId, int16(funcIdx), Base58Encode(funcData), attachment)
	transaction.SenderPublicKey = acc.PublicKey()
	transaction.Signature = acc.SignData(transaction.BuildTxData())
	return transaction
}

func (acc *Account) BuildRegisterLockContractTransaction(tokenId string, desc string) *Transaction {
	lockContract := LockContract{
		TokenId: tokenId,
	}
	transaction := NewRegisterTransaction(TokenContractLock, Base58Encode(lockContract.BuildRegisterData()), desc)
	transaction.SenderPublicKey = acc.PublicKey()
	transaction.Signature = acc.SignData(transaction.BuildTxData())
	return transaction
}

func (acc *Account) BuildRegisterNFTContractTransaction(contract string, desc string) *Transaction {
	nftContract := NFTContract{}
	transaction := NewRegisterTransaction(contract, Base58Encode(nftContract.BuildRegisterData()), desc)
	transaction.SenderPublicKey = acc.PublicKey()
	transaction.Signature = acc.SignData(transaction.BuildTxData())
	return transaction
}