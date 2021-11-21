package vsys

type NetType byte

const (
	Protocol          = "v.systems"
	Api               = 3
	addrVersion uint8 = 5

	// Fee
	VsysPrecision   int64 = 1e8
	ContractExecFee int64 = 3e7
	DefaultTxFee    int64 = 1e7
	DefaultFeeScale int16 = 100

	// Network
	Testnet NetType = 'T'
	Mainnet NetType = 'M'

	// TX_TYPE
	TxTypePayment          = 2
	TxTypeLease            = 3
	TxTypeCancelLease      = 4
	TxTypeMinting          = 5
	TxTypeContractRegister = 8
	TxTypeContractExecute  = 9

	//contract funcIdx variable
	ActionInit      = "init"
	ActionSupersede = "supersede"
	ActionIssue     = "issue"
	ActionDestroy   = "destroy"
	ActionSplit     = "split"
	ActionSend      = "send"
	ActionTransfer  = "transfer"
	ActionDeposit   = "deposit"
	ActionWithdraw  = "withdraw"

	// function index
	FuncidxSupersede     = 0
	FuncidxIssue         = 1
	FuncidxDestroy       = 2
	FuncidxSplit         = 3
	FuncidxSend          = 3
	FuncidxSendSplit     = 4
	FuncidxWithdraw      = 6
	FuncidxWithdrawSplit = 7
	FuncidxDeposit       = 5
	FuncidxDepositSplit  = 6

	// function index v2
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

	// function index for NFT
	NFTFuncidxSupersede    = 0
	NFTFuncidxIssue        = 1
	NFTFuncidxSend         = 2
	NFTFuncidxTransfer     = 3
	NFTFuncidxDeposit      = 4
	NFTFuncidxWithdraw     = 5

	// function index for NFT v2
	NFTFuncidx2Supersede    = 0
	NFTFuncidx2Issue        = 1
	NFTFuncidx2UpdateList   = 2
	NFTFuncidx2Send         = 3
	NFTFuncidx2Transfer     = 4
	NFTFuncidx2Deposit      = 5
	NFTFuncidx2Withdraw     = 6


	TokenContract               = "3GQnJtxDQc3zFuUwXKbrev1TL7VGxk5XNZ7kUveKK6BsneC1zTSTRjgBTdDrksHtVMv6nwy9Wy6MHRgydAJgEegDmL4yx7tdNjdnU38b8FrCzFhA1aRNxhEC3ez7JCi3a5dgVPr93hS96XmSDnHYvyiCuL6dggahs2hKXjdz4SGgyiUUP4246xnELkjhuCF4KqRncUDcZyWQA8UrfNCNSt9MRKTj89sKsV1hbcGaTcX2qqqSU841HyokLcoQSgmaP3uBBMdgSYVtovPLEFmpXFMoHWXAxQZDaEtZcHPkrhJyG6CdTgkNLUQKWtQdYzjxCc9AsUGMJvWrxWMi6RQpcqYk3aszbEyAh4r4fcszHHAJg64ovDgMNUDnWQWJerm5CjvN76J2MVN6FqQkS9YrM3FoHFTj1weiRbtuTc3mCR4iMcu2eoxcGYRmUHxKiRoZcWnWMX2mzDw31SbvHqqRbF3t44kouJznTyJM6z1ruiyQW6LfFZuV6VxsKLX3KQ46SxNsaJoUpvaXmVj2hULoGKHpwPrTVzVpzKvYQJmz19vXeZiqQ2J3tVcSFH17ahSzwRkXYJ5HP655FHqTr6Vvt8pBt8N5vixJdYtfx7igfKX4aViHgWkreAqBK3trH4VGJ36e28RJP8Xrt6NYG2icsHsoERqHik7GdjPAmXpnffDL6P7NBfyKWtp9g9C289TDGUykS8CNiW9L4sbUabdrqsdkdPRjJHzzrb2gKTf2vB56rZmreTUbJ53KsvpZht5bixZ59VbCNZaHfZyprvzzhyTAudAmhp8Nrks7SV1wTySZdmfLyw7vsNmTEi3hmuPmYqExp4PoLPUwT4TYt2doYUX1ds3CesnRSjFqMhXnLmTgYXsAXvvT2E6PWTY5nPCycQv5pozvQuw1onFtGwY9n5s2VFjxS9W6FkCiqyyZAhCXP5o44wkmD5SVqyqoL5HmgNc8SJL7uMMMDDwecy7Sh9vvt3RXirH7F7bpUv3VsaepVGCHLfDp9GMG59ZiWK9Rmzf66e8Tw4unphu7gFNZuqeBk2YjCBj3i4eXbJvBEgCRB51FATRQY9JUzdMv9Mbkaq4DW69AgdqbES8aHeoax1UDDBi3raM8WpP2cKVEqoeeCGYM2vfN6zBAh7Tu3M4NcNFJmkNtd8Mpc2Md1kxRsusVzHiYxnsZjo"
	TokenContractWithSplit      = "3dPGAWbTw4srh5hmMiRUhHtcxmXXLUooKGAfnmz11j5NruyJpBzZpgvADMdZS7Mevy5MAHqFbfHYdfqaAe1JEpLWt1pJWLHZBV62zUhLGmVLXUP5UDvSs24jsBRHqZMC71ciE1uYtgydKxCoFJ3rAgsYqp7GDeTU2PXS5ygDmL6WXmbAYPS8jE4sfNUbJVwpvL1cTw4nnjnJvmLET8VmQybxFt415RemV3MFPeYZay5i5gMmyZa63bjzK1uMZAVWA9TpF5YQ1NTZjPaRPvQGYVY4kY9L4LFJvUG2bib1QaNh7wUAQnTzJfRYJoy1aegFGFZFnBGp9GugH4fHAY69vGmZQnhDw3jU45G9odFyXo3T5Ww4R5szegbjCUKdUGpXf9vY2cKEMJ7i8eCkFVG1dDFZeVov1KMjkVNV8rDBDYfcp3oSGNWQQvGSUT5iGUvDRN8phy1UpR3A9uMVebvjLnVzPx9RyqQ8HaXLM8vPhLuWLoh5hk1Zi1n9nwz55XvKDYjP6eeB55yK5vpg8xjaYDnw9bjYV7ZmS7LAsHvXfnwi8y2W6vk2hGvs4rtR1vNRZSQMPGRRSuwCRJL1yngH6uHWwm2ajWxc684jApuoLdyjZomfCtdpabSyU3kp9Lrn8zT8BVY332sJPQU6gTQi8ke9s9dBxCae4cfSQM6HhuBmFc5KKWHCVG4bm4KZRYbMtidw8ZZnjaAMtcGq7k3Se6GXaTxdS3GcuttB3VB7njypyzuqAcfCdYb9ht8Y1WuTCZ1aLsXsL6eydfk2WLJVrqYpbTk6AchV5gMAEopvc3qXvzrDCedjtNsDmA56Lh6PxrrKr8aV8Wzz8aMaQ88YsVBpE8J4cDkxzo31AojhzEGVBKLmpb3bjmsaw9VkpB6yL8ngYs8eJMSPdM289TSMaEmG4eHt1jezpHTKxkuB9cwqcvhGNLWuv8KXQkik5pRMXV67Qs2FvjpzeJ81z2hnVh1wCtsa6M6qAG1gsqLHa1AVMRzsowafC99uDexwWMBS2RqsZWZBXJcUiNVULjApSnoBREYfHYEpjJ152hnTYZCAwpZMWEkVdBQpZ3zk8gbfLxB4fWMfKgJJucbKPGp1K56u7P8MHQu9aNb9dEof2mwX8rTHjk8jSQ7kXVX4Mf1JqMRWWftkV3GmU1nqYhxRGu4FjDNAomwTr5epHpcMF6P5oiXcLWh5BFQVmGYKz129oizAyUJBsZdxr2WZEGDieLxUg8cve25g28oTuCVENST4z1ZsFAN9wTa1"
	TokenContractLock  			= "4Qgfi31k6qfLxTguJg8AeYzmmgaCTJCEPQyAdoRUUSrFDc91PhkdU6C8QQSsNCFc2xEud2XnuQ4YNJ51HgdNtBdnxZcU5Rnqdzyop41Ck81v4nRKkHpTdTrfD8vTur2w4mTFeTFKVzGvGjpHXUVvT47vZiKLBHSB7FHHpGf69bu8DQGXWu6xnZZkn9v2Rfc9mByhwVLSNghNdRhrQwRWPFJ9Qt7Yb8N8WdmcUCAC6PrC3Ha3Z9w7dyf6CsKcCMS6JmB2gvNQitm9jqAfjRxDdqPBUR6TtyjSdmHP9BZRGgiVCaQH7X8fbJZVWSib4RXvFoSrqY4SfVftDY3PU4hXASaRWbaheB8m4VgM4mA8nKDbZvRWZtZ4cHdWeNFyVPs6HxHQZHrQ3GZGNPjmBSyAkGRFS7i5dK8aYWQDEYu1Xijk63UFAWuf6tRdR44ZgRjWGUZJtdQBDFB38XaU8LSFEj2eaC1yNqZ6nnGeRXDzS1q3YKsGyJTqaDDMHvPHiHonGn76JQHAZN7eGU7biaSLxoikW4MaTPSfmcTmDyPGJyJNHjc8MrpV8aQSaGGyDkf1a9MpoJcyEjsPFQbxYzSJVqFEFg2oUL7Z8VUtJK2kYcWDz7w8UiiQqe3uuQnKDGb1nJ5Ad3W8ZPfVP6YHbJrnBKZXMMypNoveokVvxZMCkSNYDsoBxJzrwFvm5DcDJbePQU6VbeZ5SzQw9XTAw4DZpxkQm9RwRE9PXPqogpp9P6LhaiUa6ZD1cWUAHypjWLJ2Rds96oap3biBp5aESunuh99HByoXg5Aa7EQ3FrEvmeq9TLVFYpJraZyW"
	TokenContractV2WhiteList    = "7BekqFZ2yZqjiQFFnsxL4CDRFWCjHdZvFXQd6sxAgEktxwn5kkR6vkV27SFC7VmhuMysVfunZWTtHAqPjg4gGz72pha6TMUerSUSXSn7BHaVexyQJoUfqDT5bdr3XVpok1mU2gT29mwtJ6BibizpAEgZZncZauDnvqrWWdkCmRP8VXpPBiPEaUZuq9eRusrUcc5YHshhN6BVkArN84tarVQH3pTRmiekdQveuxFw4r4weXUxwEGCkYX3Zqeqc4mmRsajVCQwV5DuGTEwaBVWiAAfHLGPFgJF6w6aP3d22tdBRLqZ2Y4G5WHdhMunNDEZ2E79w7gbwqDXtz3eVfGtyET5NZEJGmM2S8pZSn2MPjvfPAYZMa9Zd4WXnPLZng1pxjYvrpqPDy27VQu1rhvxXMNPVMdP9QyCQSoExZUot1FmskS1NcmzKfguwsSWR1Z1py58iVDKm8t7x7RnaP7avcjtvixJQkPGg7qaxBKfRQ26vFePWeNdkbJwQJvqComvjEg3hEYjQrysk3j3M9QWEgXQzRqTPTFEVCTJSbdpL2GyYXYC4cLcB81UzJuWf2zoERNPdfpHwumoaaaSutfg7dccbWRaqogrBf6u9PfANQm9TsFca37UHhxvsq8WZdu71NQCY1V7w9NKKLbHF7MjjyCs6w2TM4Ej9Tyj8hFR4qo3MosgSbmQt298aEB3qQHVF8FshVwGg2vqAK7PNBHE7KgBgXQJiVRc4X1XZvWQt4uASvMowRECURoMZ17z2s3LnDrQYVqYedfzjJXxwsWXQkoQp51WWkFfp7QStBtfEhdUx15wtD8sjDdNrda8n3P6sNrN8J7NXxH4JPE7DzLLCjPSbn5Yc2jzomULSRiQN2yzC5qE43XiHB89VFqTHTduCFbP3Pom3uc5iBgjW9ky8LyPBMcsqQZSv99adjgbKpeaGPtJN6iUQ9mae1ddw6SBKTxZVZvqK6k7dJBjJ5UsFDyXLWkm8jogkRCFBfXPxmxyB5ihqk2wnsWNEbKEz6sg6RJqy5SR9A8r3QEx8FZt5z4DJpHyUAoi6KKVHEJfRvdjtjSDrayG2WUrBCgTTHsyGZEnuXLRXpy7XmdzFSwKSr4p7NPbAqt44yHdgjycn2MY5X1P9rneBdh4LukH3syRAarjmTSZr67QexRE4cca5fnxUZJ2zYNWRynqWmZy6aCBLBQziP81bHHbN5WP9MMseovCvzTpMso9TB3QLSRkCphJpyvv9qLN4tpFB9r9g3UGhTqqJFvxJDcLwR485AqLymM91kMjTvodniJ4coymUeE3MjGf2P67z4UiBDBxnzWbkCzmaPpkWFY9125hg9SovQrJnn9zzpF5smp7oiHhjrkzyi2G4qWVidtaWi6TipZFXwb8z6TSSjZkaj4SWexgnE2bUKeJS9P1xYwVSX39At735bqhfKCNP29n7UzX7bMwQiTWWK8bCiCpYSXfcfUpxtbYXdHgGMEZzpzawS9H5UeFiw31rS5Caps7QQJmMeetAuDa8tsiMJ9QauABLfJ4G6Hjkn5GM9jH9yXJWj2boH1U4ErVQXbr9KvmSsSsLeLLc3XeKQaczMtLroQax4D5estuP3Cy1gfqhbTsEWL2HkF7dUKDnuLmzsjv3kZXF9PMhcVR1Qj9j8KaYWYqKYV5TxXkrPrzSVa1yYEjU71A6ZYW327vgFJYFUJmx9vqTGym3yRiSoJiaYVfgf8iLwqS1EKSTMiisxE8hCHfKiew4YmiCTxPkq7pc5tHrKkogoRX7GdDnX93BsxGACu9nEbXwDZERLFLexrnRKpWDjqR2Z6CLWhXNPDJYMcUQ5rfGAhgu4ZK16q1"
	TokenContractV2BlackList  	= "2wsw3fMnDpB5PpXoJxJeuE9RkRNzQqZrV35hBa366PhG9Sb3sPeBNeYQo8CuExtT8GpKuc84PLMsevNoodw7YGVf24PKstuzhM96H2gQoawx4BVNZwy3UFyWn156SyZakSvJPXz521p1nzactXZod1Qnn7BWYXFYCU3JFe1LGy35Sg6aXwKz6swFmBtPg1vBeQsUq1TJ5GXkDksaUYjB8ix9ScNNG8faB1mCCMWwfrcr6PyBA7YeHsTLD86zuviak6HQEQQi9kqVr4XhnDJnZyiTKGcNDo49KZyTyvkPmkFyDEhLf9DYrJM3niePqtDQ9unJj52Bku7f47hrxo83eSh3UPncyq8Hti2Ffhgb8ZFCFdnPyRDEZ1YbKFGAsJL3h3GdPFoVdnYySmnVJWrm6fVUdGgkA5ijMeqEUpXte1m7MFYCJ1wQchjebpLk3NnZzrT8FysUJVUgUzmkoSniF2UPEPXuF9cyWFWGGoZjfDWqarPMi7miqdCPQMMw4QRvSWkB3gVyeZykAvKYzXm8wYGV6HDbipZeVoyZ1UVeR6E5C4VZQmjs4GupAR9EuT5mt1ALFT4HyAMX6RCRxjeHoSgnnUJcEiRHapAYSene174RvVkRGLTtonWTYnsXUrtPD6xks4GdpQWQv89EdNWFEtmMfyVvUEFuTPGXUS5TuqYxCzg8Gor5WjPip2wDmoMYQ3wikJoRpYSfRVw88RHQPBmkHrpeHYWkAx6N7Yk4WwgBF9SVVtEWnWmPVVbuH2bQrvks4iGL8DnmEiLMs6JuFsg3a3cMHqbdvQgfu72XYKFqQzzDbDhaqFKpR3bxgMMiJvGbPuydPk9DCsG5KpqZepkkD6RGhWTQzga9G6y6ryctoGZPBHpFRwirALkksarQSEuGryhatvnjqG9U14zyW2KvJYrErMyUVy3wNK5wRqAKMjE6hFPdoH9Cn6TYQLebVTBoYTfimn5gBmgnKqBtXSfUxiwrjWujQPGxgtbNCL1RXRNRJ8nrtcpphQyRVZ8JVeubYq1zM7G1AUurEyAQi64rcbsimGptcXMAvt9TbwDjpUGRWvF6dyw1XijcukfZBQh1fG5C8peumkGnP8PemmYWKP7qsifNc44PqnNG5qYVivwtK4sz2h3B6pwneX8XNYtGSjVJCb6gJ7oDG45shocvALKNu7LwfJxXT7MPAdx7CjbHU5B3qs71wJphwkc4yWa6hHTamPTGRFGuhJa4kFfeGMctE1WZrFe47L32fKZkSxaX1sguoi5w9UPHw6udJiKPYENSSbASYpfS9q8suCs1bbq8jdMhCwoGMDZaA4MNAW1Q6sLSX6ezZ436AMbVnXZLQW8jdBaX8rvRSMJu8fdYU9PHq4MkoczxNz5jNvRiTX9jTpN1Z1P5rtgnf6XN9vzTLdqsvwZcXqvSdBwdTVgk7qn9uNjuFZEgSmA6rnPhSu6TMxJLmjKP93uqiNmXsj1NKtqBZiHjrRaUzA4pAFEyfZTdo8oaDH7umSBU2s9ff5Cruds7cYFopLm2KavHH33S7BczL7FMXAcqrESiPUzhUhHbkBKHGiCAUMVE8zxo6Eo85W2PGn6D39MaUfahEmzq8zxmrDQdmagx5EQZUev3fNCFzTzU4zpY1sra5ZPknXJkyKKfj4r9xy9Kfd8s5hsiKFyX6V1Kc2T1Ehpdkobwb7Wc8V1n1GaeL7jRgvhVg1inPaWZ3zyqNBjxnzqtLpZor3VdXLo6SikzWNahCMLNMXaoBvmJDEJUazC9qGxin7SC3YWCTAyoskJRhVMp592ehmpruu2azeCHBF2rzP6LabikVfkBSeAzGQKVeiEkU3devRNpjNM4YDXQDm9wbkPKWrqBK4SRdo44PRYG3XwNhu2gpNX8b9AuirrbRPiaJ1tJ7rzodHzLheMyUMXRB9nYx8JgrhkZzPZa4oUxo8JUNuKZnn7Ku7fEt5y"
	TokenContractNFT  		  	= "VJodouhmnHVDwtkBZ2NdgahT7NAgNE9EpWoZApzobhpua2nDL9D3sbHSoRRk8bEFeme2BHrXPdcq5VNJcPdGMUD54Smwatyx74cPJyet6bCWmLciHE2jGw9u5TmatjdpFSjGKegh76GvJstK3VaLagvsJJMaaKM9MNXYtgJyDr1Zw7U9PXV7N9TQnSsqz6EHMgDvd8aTDqEG7bxxAotkAgeh4KHqnk6Ga117q5AJctJcbUtD99iUgPmJrC8vzX85TEXgHRY1psW7D6daeExfVVrEPHFHrU6XfhegKv9vRbJBGL861U4Qg6HWbWxbuitgtKoBazSp7VofDtrZebq2NSpZoXCAZC8DRiaysanAqyCJZf7jJ8NfXtWej8L9vg8PVs65MrEmK8toadcyCA2UGzg6pQKrMKQEUahruBiS7zuo62eWwJBxUD1fQ1RGPk9BbMDk9FQQxXu3thSJPnKktq3aJhD9GNFpvyEAaWigp5nfjgH5doVTQk1PgoxeXRAWQNPztjNvZWv6iD85CoZqfCWdJbAXPrWvYW5FsRLW1xJ4ELRUfReMAjCGYuFWdA3CZyefpiDEWqVTe5SA6J6XeUppRyXKpKQTc6upesoAGZZ2NtFDryq22izC6D5p1i98YpC6Dk1qcKevaANKHH8TfFoQT717nrQEY2aLoWrA1ip2t5etdZjNVFmghxXEeCAGy3NcLDFHmAfcBZhHKeJHp8H8HbiMRtWe3wmwKX6mPx16ahnd3dMGCsxAZfjQcy4J1HpuCm7rHMULkixUFYRYqx85c7UpLcijLRybE1MLRjEZ5SEYtazNuiZBwq1KUcNipzrxta9Rpvt2j4WyMadxPf5r9YeAaJJp42PiC6SGfyjHjRQN4K3pohdQRbbG4HQ95NaWCy7CAwbpXRCh9NDMMQ2cmTfB3KFW2M"
	TokenContractV2NFTWhiteList = "RVbUxLpK4Vi15qHaDvy1e4kjNmQqLrcdDqXBjMBHB1RH3xPQS7tcuEEoxHZ72mUVttTa3EHMmD6gRUcKSXW1kQkNpBNEFzgfM2qSS2BRvdHdke9xNPgu6i2m8KbViKeBiq9ydsHMMkyZL5ShfpfBD8BfJmTiSERjgx7voeusYwhTWT2VDg4E2k3krqDmbCVTtvxm6nydPxxwDH9RBBWJwLjkMHRdaoDFXrDBA5KXc6YgSnT9P6JZKRAbTYAzAEF1514oGEvRjgz8HUhCzjVjmTUgjHXeYT416f8soj1xzmQW4LCZszSGeo1hMh5HZWshg9NwJGkJW9C2HVXr31RvSvKcLK6NgDEnxaci6H9mjBagXW5MNTp7KKKmEuurKVHw1GydPE2Zx91cvFzQ5Q4xiUWhRtMCyMmDp5AH4STRgcdyK4DsFR2unaxhQgEHzxCVR3EPe6cddnW6ZdGDvSV6KzDz1RBPUu6Ex7PUJ6L8GAdbunppVJrmLYER7aySP67Z2j7mqaRwCTDYSMeh47a3aCpuhG9dGoTMF9UZhxBZu5irg7fwzifPWJKQBKPp2JvefKanEwxPqLUbFq5a1Gu8dqFxaa6cC248EHHK1WqHtkZcUdFN65V2rhGCxPvNKADwH7STsn7awHuowUTJJe1EnTHvyJ9Yzd8Kg8JFn1Q5tTEAQTFujCMnriXxbqmjNYhQghqxyxjVL447NDmvYBqey7jcp4CBde9myjbFdxDvbtVDawwxoGvGnzosn14RsBBBZjxcqKDHAMbkJLdWkjL5n8yjXxiAZfbEt4Bk5uDK4YP8YgbGSHKVhGWYWbmKxEDGAziRnGTdwy79RPtTsb6zx8fCGfU3gCD4skY6ny6bM28Ue81YBMPt89TKm6Gt8GLhXHft8vSp9cUiV6dDauVHmyu1vgACcN7kpu1yMZExZEazUDSBf9SuiyEZWXsDkXjm4ayauX7oTakMuFFyRACtrAVowB9thQGt1jWeLdg1kVhucrdLJ2fj2NWRX87Q6UqAjmtjyVBKyntpheWpJTg8GzLbH2ASVF556pgpC1jNQXo3HxEBoTGnNzf4v3E8xqibBGxE3wY7hvpvEXv7ww3Yw1TxLjxtMDjqLuZvXrMWoqYxQEanHBfJkz3bCMeDajJbYwyqsVkgTXCpGNgXdUGYD1w5TV5DhbbPjxFnJ73aVJ7ANVmNi6UYNxCLWPkVmNNnRMWpGc2sBw9cK7GpChQedwK1u5GaRd1yR3JVTCn7GdwELFF8BWSiWCShh8aVNz1EThqz8uMUeU6iUr5W1LTVTFiJ1kfEDVZyLEKZbZWSJSV83c9bHW4jFm3rRraEwQiMaRvBkq3ELeFpZKMwKtKte9UWUniZU9QbAZwMvNAEQqkUavNwKQS7haUxCUR5jv6iBj6hZ25qVZU9CduqH3YZAmonAQsm9WVTbh91qeFVtjBAmVvfUJT1y6AGSWyHhpyzSrGqAeornSkwdcaZtTDwsVMqy6Y1tha2493HmQ7dE8Cty8VH"
	TokenContractV2NFTBlackList = "RVbUxLpK4Vi15qHaDvy1e4kjNmQqLrcdDqXBjMBHB1RH3xPQS7tcuEEoxHZ72mUVttTa3EHMmD6gRUcKSXW1kQkNpBNEFzgfM2qSS2BRvdHdke9xNPgu6i2m8KbViKeBiq9ydsHMMkyZL5ShfpfBD8BfJmTiSERjgx7voeusYwhTWT2VDg4E2k3krqDmbCVTtvxm6nydPxxwDH9RBBWJwLjkMHRdaoDFXrDBA5KXc6YgSnT9P6JZKRAbTYAzAEF1514oGEvRjgz8HUhCzjVjmTUgjHXdpR6yuM3CwR8EYGbH9HNEcnUoTEcjLwGNCtUb2QBRJLzpu2VzhqJTZkXrrjuREfFpkBPN8WCYgi4LguF7avFQA4atahzVKU98j7UG9byW4ERkdHFduxSWX6nun8NYxgw4k1LyGe7A6NEo6vfAeBryYb5V2CowTiXb4xxhzjytuNPEckZYJcMLkgxtfzmcbfnqx7ff4hgWX2L5AWh5y9K21BDhrCWzjnb81atVvwBivMSBvaoFGNRj5RJ8Qz4r7cGyZ34PqcdZsbTvnTJTzHNLUxSyMQNqNBa3vwzifuuwPBUfZ7xKXGbQVo19D7BWzgVhmye1CKNxsW3QRNRYqrWZaiUqRmb1ids8BfAoyNEk9HzpJ3zYq28rjmZ4nZp8TCnaH7jZxUDeVhBfzQEmEbSSHDaPSnDNtLQ9VKQa9ov3ZTC2muUw7P8hUR3N5casgZpG3a23uYCLT4TKSRxQU9JZ3kPe2r8wVUvhmeJP9EJiUi4CmArEGQoGhQZLZuCneVZdWbreqMhnQLVjVekv5NPWwEf4UNgFH5RS6XQ2mcHqHd5KHAG9xa5qaPDQ6YZ4Fagh2bw92fw4CiDYiCodsPVYFqYDrWRJmbdY9xQmRZ4M4w4UKQU8HmwYhqCEtSLrH3dmuia2JkvEv8f47QZq78rLaUmcY5Fch9zjEZTe63StcdB1GycZoHK7iCoHrt1AjJ8Ex5J9xgJxbQ5nZEAS3vrqy9znznEuhawtL9PYn5VYEbFX6VdCTxbYcjehKXQRCxouyYAS35BqmoYuVYM54qtLUbgrXPUTefQ9XGGGAazeaSDkS3hASwx3DuLU7MYDmpYgyoefxEVVi8D2GVH6TaiU34qFTiTEK1348ZzruUugT5DZRwFgf5t197iYngD9AKb71TBsA2ZnJueNVFEFXmXgFKMU2eZvAZBFcoH1EAUttCgSztBwXMS9Cwqa5kXXwr1cxNUESzNoDgZWaEAHPiCB2PAjBdsgWGTbkepSBMbQTj5aKV4LybHoWs8JNKpHMMwSBeRLNmk3ibGkNu2qe7ZcZZJJqNz9vZjvhKJ8Ws1HYrqaPb7ysBHW7fU1mrh95y7AkSZGsKEbGNpyijpT66Q4wxpuysU6L3wDMdapdQLjdyBs3rjQtXRbhSiRyYZLShXUantffsBkMWmwD2WZg6Dp6hpZWBEFGq3kD1ysVbi47HqDTsD1aWTuE5hCY9XHmveE3WmRd8p56YZxmrXKh9Ns"
	ContractDescriptor          = "1JTuFtDUkasyJs9FCrgJ3CRDU8J9sh5oJbG9LqVpdBbEmDwjM1hbDnmhcBBdWHgaLDDiExdgnRmQDJwY3yfpm441KsPQpiPf9JJzNH8fLNoYGT5S8P8Qki4nsP9DmJ6qDLj4ktW2fJeEosWaGqULwcpCFiBFeQiA6DN8tJE6nEvz75ku4U66mMbtJGi2g9pJyHRoxT67Fzmnvpa4UhXyB6J3GgyM6od68o3vzXd2kzCGnXGK3jXNUTBKXwWQTiRq8oQ2XEKp3vNXSrYktjYbSLh7g4Zq6XrkwxwSpv2m3AehXwLbZ8wA7CYaKcgXngVdc2rgMMwMUygXnAxrcZNKRFyP7cjZ6213poGqHe7gkqBEpr93uQPRm5PCW3gQVr37KTafHpUfr1JSbmT4GjwEADP2L1KGmnVJJxCsnLgR2QnhiopaHcni2rYPe2TVStJ2kHmJwNBQ7YBWMqu45TSWySfehcwqXeKktJhrWrPzSKr1ca8Q9PXptKQVWaFn2pfqXWrdfRci8LLZD7FgPdp7Q4wr7PpEwMBkA7tUdw2PNYiHFLYaFduFx1ya1o"
	ContractWithSplitDescriptor = "1RypGiL5eNbDESxn2SVM8HrLF6udvXV6YmwvFsp4fLJfRcr7nQuVFMvXn6KmWJeq8c53tdrutZcsQA2zyHb8Wj1tQUjGmitP6kLzcnpQXEq7AUZpMT6j7LCrhJvs3oLCCr7SSpz3h4iJJJg9WuL7Acbsw1x2AK4tRSZWXyrnLgqWhgqbTdfmxFGHjD58XrScBibJ9AUwEWCAeAna3NFofSZaSDxFJAK2adrrHhJdktQCQobMJMmC164HtJKF569naoMREkncYedQwXWk4uyPzGTUKsfXFwLaR77wv8gtNEjqwvGtpdFJELyJ3RC2F7exhqiiVxTaoGrAanuv1bianVbKqPAygPaGrhA1H3JmQWksNhg6q7dtPvBuqWDqDs4DkhV35JhNFeiER18o49pxX8zR1n1jvis6QrU2cD1Cn3yXwSZaW8TXWMKZ7ULRo1UcJykQvQCLq3EBVfzf6iULhuRagTnJ3Sq4tFSxgnNPhATLDreQpEe1BA3SfRWKRskLFjXV5aMeYxgFLfqYEFJ37BaRVyFZDSUgrKLMnNzrZZG2P81t7MhT6GpDApLZkNtjdGRMQGFsRN2azGruQReFnXeB3mScaxgfhGxcu9B"
)

var wordList = []string{"abandon", "ability", "able", "about", "above", "absent", "absorb", "abstract", "absurd", "abuse", "access",
	"accident", "account", "accuse", "achieve", "acid", "acoustic", "acquire", "across", "act", "action",
	"actor", "actress", "actual", "adapt", "add", "addict", "address", "adjust", "admit", "adult", "advance",
	"advice", "aerobic", "affair", "afford", "afraid", "again", "age", "agent", "agree", "ahead", "aim", "air",
	"airport", "aisle", "alarm", "album", "alcohol", "alert", "alien", "all", "alley", "allow", "almost",
	"alone", "alpha", "already", "also", "alter", "always", "amateur", "amazing", "among", "amount", "amused",
	"analyst", "anchor", "ancient", "anger", "angle", "angry", "animal", "ankle", "announce", "annual",
	"another", "answer", "antenna", "antique", "anxiety", "any", "apart", "apology", "appear", "apple",
	"approve", "april", "arch", "arctic", "area", "arena", "argue", "arm", "armed", "armor", "army", "around",
	"arrange", "arrest", "arrive", "arrow", "art", "artefact", "artist", "artwork", "ask", "aspect", "assault",
	"asset", "assist", "assume", "asthma", "athlete", "atom", "attack", "attend", "attitude", "attract",
	"auction", "audit", "august", "aunt", "author", "auto", "autumn", "average", "avocado", "avoid", "awake",
	"aware", "away", "awesome", "awful", "awkward", "axis", "baby", "bachelor", "bacon", "badge", "bag",
	"balance", "balcony", "ball", "bamboo", "banana", "banner", "bar", "barely", "bargain", "barrel", "base",
	"basic", "basket", "battle", "beach", "bean", "beauty", "because", "become", "beef", "before", "begin",
	"behave", "behind", "believe", "below", "belt", "bench", "benefit", "best", "betray", "better", "between",
	"beyond", "bicycle", "bid", "bike", "bind", "biology", "bird", "birth", "bitter", "black", "blade", "blame",
	"blanket", "blast", "bleak", "bless", "blind", "blood", "blossom", "blouse", "blue", "blur", "blush",
	"board", "boat", "body", "boil", "bomb", "bone", "bonus", "book", "boost", "border", "boring", "borrow",
	"boss", "bottom", "bounce", "box", "boy", "bracket", "brain", "brand", "brass", "brave", "bread", "breeze",
	"brick", "bridge", "brief", "bright", "bring", "brisk", "broccoli", "broken", "bronze", "broom", "brother",
	"brown", "brush", "bubble", "buddy", "budget", "buffalo", "build", "bulb", "bulk", "bullet", "bundle",
	"bunker", "burden", "burger", "burst", "bus", "business", "busy", "butter", "buyer", "buzz", "cabbage",
	"cabin", "cable", "cactus", "cage", "cake", "call", "calm", "camera", "camp", "can", "canal", "cancel",
	"candy", "cannon", "canoe", "canvas", "canyon", "capable", "capital", "captain", "car", "carbon", "card",
	"cargo", "carpet", "carry", "cart", "case", "cash", "casino", "castle", "casual", "cat", "catalog", "catch",
	"category", "cattle", "caught", "cause", "caution", "cave", "ceiling", "celery", "cement", "census",
	"century", "cereal", "certain", "chair", "chalk", "champion", "change", "chaos", "chapter", "charge",
	"chase", "chat", "cheap", "check", "cheese", "chef", "cherry", "chest", "chicken", "chief", "child",
	"chimney", "choice", "choose", "chronic", "chuckle", "chunk", "churn", "cigar", "cinnamon", "circle",
	"citizen", "city", "civil", "claim", "clap", "clarify", "claw", "clay", "clean", "clerk", "clever", "click",
	"client", "cliff", "climb", "clinic", "clip", "clock", "clog", "close", "cloth", "cloud", "clown", "club",
	"clump", "cluster", "clutch", "coach", "coast", "coconut", "code", "coffee", "coil", "coin", "collect",
	"color", "column", "combine", "come", "comfort", "comic", "common", "company", "concert", "conduct",
	"confirm", "congress", "connect", "consider", "control", "convince", "cook", "cool", "copper", "copy",
	"coral", "core", "corn", "correct", "cost", "cotton", "couch", "country", "couple", "course", "cousin",
	"cover", "coyote", "crack", "cradle", "craft", "cram", "crane", "crash", "crater", "crawl", "crazy",
	"cream", "credit", "creek", "crew", "cricket", "crime", "crisp", "critic", "crop", "cross", "crouch",
	"crowd", "crucial", "cruel", "cruise", "crumble", "crunch", "crush", "cry", "crystal", "cube", "culture",
	"cup", "cupboard", "curious", "current", "curtain", "curve", "cushion", "custom", "cute", "cycle", "dad",
	"damage", "damp", "dance", "danger", "daring", "dash", "daughter", "dawn", "day", "deal", "debate",
	"debris", "decade", "december", "decide", "decline", "decorate", "decrease", "deer", "defense", "define",
	"defy", "degree", "delay", "deliver", "demand", "demise", "denial", "dentist", "deny", "depart", "depend",
	"deposit", "depth", "deputy", "derive", "describe", "desert", "design", "desk", "despair", "destroy",
	"detail", "detect", "develop", "device", "devote", "diagram", "dial", "diamond", "diary", "dice", "diesel",
	"diet", "differ", "digital", "dignity", "dilemma", "dinner", "dinosaur", "direct", "dirt", "disagree",
	"discover", "disease", "dish", "dismiss", "disorder", "display", "distance", "divert", "divide", "divorce",
	"dizzy", "doctor", "document", "dog", "doll", "dolphin", "domain", "donate", "donkey", "donor", "door",
	"dose", "double", "dove", "draft", "dragon", "drama", "drastic", "draw", "dream", "dress", "drift", "drill",
	"drink", "drip", "drive", "drop", "drum", "dry", "duck", "dumb", "dune", "during", "dust", "dutch", "duty",
	"dwarf", "dynamic", "eager", "eagle", "early", "earn", "earth", "easily", "east", "easy", "echo", "ecology",
	"economy", "edge", "edit", "educate", "effort", "egg", "eight", "either", "elbow", "elder", "electric",
	"elegant", "element", "elephant", "elevator", "elite", "else", "embark", "embody", "embrace", "emerge",
	"emotion", "employ", "empower", "empty", "enable", "enact", "end", "endless", "endorse", "enemy", "energy",
	"enforce", "engage", "engine", "enhance", "enjoy", "enlist", "enough", "enrich", "enroll", "ensure",
	"enter", "entire", "entry", "envelope", "episode", "equal", "equip", "era", "erase", "erode", "erosion",
	"error", "erupt", "escape", "essay", "essence", "estate", "eternal", "ethics", "evidence", "evil", "evoke",
	"evolve", "exact", "example", "excess", "exchange", "excite", "exclude", "excuse", "execute", "exercise",
	"exhaust", "exhibit", "exile", "exist", "exit", "exotic", "expand", "expect", "expire", "explain", "expose",
	"express", "extend", "extra", "eye", "eyebrow", "fabric", "face", "faculty", "fade", "faint", "faith",
	"fall", "false", "fame", "family", "famous", "fan", "fancy", "fantasy", "farm", "fashion", "fat", "fatal",
	"father", "fatigue", "fault", "favorite", "feature", "february", "federal", "fee", "feed", "feel", "female",
	"fence", "festival", "fetch", "fever", "few", "fiber", "fiction", "field", "figure", "file", "film",
	"filter", "final", "find", "fine", "finger", "finish", "fire", "firm", "first", "fiscal", "fish", "fit",
	"fitness", "fix", "flag", "flame", "flash", "flat", "flavor", "flee", "flight", "flip", "float", "flock",
	"floor", "flower", "fluid", "flush", "fly", "foam", "focus", "fog", "foil", "fold", "follow", "food",
	"foot", "force", "forest", "forget", "fork", "fortune", "forum", "forward", "fossil", "foster", "found",
	"fox", "fragile", "frame", "frequent", "fresh", "friend", "fringe", "frog", "front", "frost", "frown",
	"frozen", "fruit", "fuel", "fun", "funny", "furnace", "fury", "future", "gadget", "gain", "galaxy",
	"gallery", "game", "gap", "garage", "garbage", "garden", "garlic", "garment", "gas", "gasp", "gate",
	"gather", "gauge", "gaze", "general", "genius", "genre", "gentle", "genuine", "gesture", "ghost", "giant",
	"gift", "giggle", "ginger", "giraffe", "girl", "give", "glad", "glance", "glare", "glass", "glide",
	"glimpse", "globe", "gloom", "glory", "glove", "glow", "glue", "goat", "goddess", "gold", "good", "goose",
	"gorilla", "gospel", "gossip", "govern", "gown", "grab", "grace", "grain", "grant", "grape", "grass",
	"gravity", "great", "green", "grid", "grief", "grit", "grocery", "group", "grow", "grunt", "guard", "guess",
	"guide", "guilt", "guitar", "gun", "gym", "habit", "hair", "half", "hammer", "hamster", "hand", "happy",
	"harbor", "hard", "harsh", "harvest", "hat", "have", "hawk", "hazard", "head", "health", "heart", "heavy",
	"hedgehog", "height", "hello", "helmet", "help", "hen", "hero", "hidden", "high", "hill", "hint", "hip",
	"hire", "history", "hobby", "hockey", "hold", "hole", "holiday", "hollow", "home", "honey", "hood", "hope",
	"horn", "horror", "horse", "hospital", "host", "hotel", "hour", "hover", "hub", "huge", "human", "humble",
	"humor", "hundred", "hungry", "hunt", "hurdle", "hurry", "hurt", "husband", "hybrid", "ice", "icon", "idea",
	"identify", "idle", "ignore", "ill", "illegal", "illness", "image", "imitate", "immense", "immune",
	"impact", "impose", "improve", "impulse", "inch", "include", "income", "increase", "index", "indicate",
	"indoor", "industry", "infant", "inflict", "inform", "inhale", "inherit", "initial", "inject", "injury",
	"inmate", "inner", "innocent", "input", "inquiry", "insane", "insect", "inside", "inspire", "install",
	"intact", "interest", "into", "invest", "invite", "involve", "iron", "island", "isolate", "issue", "item",
	"ivory", "jacket", "jaguar", "jar", "jazz", "jealous", "jeans", "jelly", "jewel", "job", "join", "joke",
	"journey", "joy", "judge", "juice", "jump", "jungle", "junior", "junk", "just", "kangaroo", "keen", "keep",
	"ketchup", "key", "kick", "kid", "kidney", "kind", "kingdom", "kiss", "kit", "kitchen", "kite", "kitten",
	"kiwi", "knee", "knife", "knock", "know", "lab", "label", "labor", "ladder", "lady", "lake", "lamp",
	"language", "laptop", "large", "later", "latin", "laugh", "laundry", "lava", "law", "lawn", "lawsuit",
	"layer", "lazy", "leader", "leaf", "learn", "leave", "lecture", "left", "leg", "legal", "legend", "leisure",
	"lemon", "lend", "length", "lens", "leopard", "lesson", "letter", "level", "liar", "liberty", "library",
	"license", "life", "lift", "light", "like", "limb", "limit", "link", "lion", "liquid", "list", "little",
	"live", "lizard", "load", "loan", "lobster", "local", "lock", "logic", "lonely", "long", "loop", "lottery",
	"loud", "lounge", "love", "loyal", "lucky", "luggage", "lumber", "lunar", "lunch", "luxury", "lyrics",
	"machine", "mad", "magic", "magnet", "maid", "mail", "main", "major", "make", "mammal", "man", "manage",
	"mandate", "mango", "mansion", "manual", "maple", "marble", "march", "margin", "marine", "market",
	"marriage", "mask", "mass", "master", "match", "material", "math", "matrix", "matter", "maximum", "maze",
	"meadow", "mean", "measure", "meat", "mechanic", "medal", "media", "melody", "melt", "member", "memory",
	"mention", "menu", "mercy", "merge", "merit", "merry", "mesh", "message", "metal", "method", "middle",
	"midnight", "milk", "million", "mimic", "mind", "minimum", "minor", "minute", "miracle", "mirror", "misery",
	"miss", "mistake", "mix", "mixed", "mixture", "mobile", "model", "modify", "mom", "moment", "monitor",
	"monkey", "monster", "month", "moon", "moral", "more", "morning", "mosquito", "mother", "motion", "motor",
	"mountain", "mouse", "move", "movie", "much", "muffin", "mule", "multiply", "muscle", "museum", "mushroom",
	"music", "must", "mutual", "myself", "mystery", "myth", "naive", "name", "napkin", "narrow", "nasty",
	"nation", "nature", "near", "neck", "need", "negative", "neglect", "neither", "nephew", "nerve", "nest",
	"net", "network", "neutral", "never", "news", "next", "nice", "night", "noble", "noise", "nominee",
	"noodle", "normal", "north", "nose", "notable", "note", "nothing", "notice", "novel", "now", "nuclear",
	"number", "nurse", "nut", "oak", "obey", "object", "oblige", "obscure", "observe", "obtain", "obvious",
	"occur", "ocean", "october", "odor", "off", "offer", "office", "often", "oil", "okay", "old", "olive",
	"olympic", "omit", "once", "one", "onion", "online", "only", "open", "opera", "opinion", "oppose",
	"option", "orange", "orbit", "orchard", "order", "ordinary", "organ", "orient", "original", "orphan",
	"ostrich", "other", "outdoor", "outer", "output", "outside", "oval", "oven", "over", "own", "owner",
	"oxygen", "oyster", "ozone", "pact", "paddle", "page", "pair", "palace", "palm", "panda", "panel", "panic",
	"panther", "paper", "parade", "parent", "park", "parrot", "party", "pass", "patch", "path", "patient",
	"patrol", "pattern", "pause", "pave", "payment", "peace", "peanut", "pear", "peasant", "pelican", "pen",
	"penalty", "pencil", "people", "pepper", "perfect", "permit", "person", "pet", "phone", "photo", "phrase",
	"physical", "piano", "picnic", "picture", "piece", "pig", "pigeon", "pill", "pilot", "pink", "pioneer",
	"pipe", "pistol", "pitch", "pizza", "place", "planet", "plastic", "plate", "play", "please", "pledge",
	"pluck", "plug", "plunge", "poem", "poet", "point", "polar", "pole", "police", "pond", "pony", "pool",
	"popular", "portion", "position", "possible", "post", "potato", "pottery", "poverty", "powder", "power",
	"practice", "praise", "predict", "prefer", "prepare", "present", "pretty", "prevent", "price", "pride",
	"primary", "print", "priority", "prison", "private", "prize", "problem", "process", "produce", "profit",
	"program", "project", "promote", "proof", "property", "prosper", "protect", "proud", "provide", "public",
	"pudding", "pull", "pulp", "pulse", "pumpkin", "punch", "pupil", "puppy", "purchase", "purity", "purpose",
	"purse", "push", "put", "puzzle", "pyramid", "quality", "quantum", "quarter", "question", "quick", "quit",
	"quiz", "quote", "rabbit", "raccoon", "race", "rack", "radar", "radio", "rail", "rain", "raise", "rally",
	"ramp", "ranch", "random", "range", "rapid", "rare", "rate", "rather", "raven", "raw", "razor", "ready",
	"real", "reason", "rebel", "rebuild", "recall", "receive", "recipe", "record", "recycle", "reduce",
	"reflect", "reform", "refuse", "region", "regret", "regular", "reject", "relax", "release", "relief",
	"rely", "remain", "remember", "remind", "remove", "render", "renew", "rent", "reopen", "repair", "repeat",
	"replace", "report", "require", "rescue", "resemble", "resist", "resource", "response", "result", "retire",
	"retreat", "return", "reunion", "reveal", "review", "reward", "rhythm", "rib", "ribbon", "rice", "rich",
	"ride", "ridge", "rifle", "right", "rigid", "ring", "riot", "ripple", "risk", "ritual", "rival", "river",
	"road", "roast", "robot", "robust", "rocket", "romance", "roof", "rookie", "room", "rose", "rotate",
	"rough", "round", "route", "royal", "rubber", "rude", "rug", "rule", "run", "runway", "rural", "sad",
	"saddle", "sadness", "safe", "sail", "salad", "salmon", "salon", "salt", "salute", "same", "sample", "sand",
	"satisfy", "satoshi", "sauce", "sausage", "save", "say", "scale", "scan", "scare", "scatter", "scene",
	"scheme", "school", "science", "scissors", "scorpion", "scout", "scrap", "screen", "script", "scrub", "sea",
	"search", "season", "seat", "second", "secret", "section", "security", "seed", "seek", "segment", "select",
	"sell", "seminar", "senior", "sense", "sentence", "series", "service", "session", "settle", "setup",
	"seven", "shadow", "shaft", "shallow", "share", "shed", "shell", "sheriff", "shield", "shift", "shine",
	"ship", "shiver", "shock", "shoe", "shoot", "shop", "short", "shoulder", "shove", "shrimp", "shrug",
	"shuffle", "shy", "sibling", "sick", "side", "siege", "sight", "sign", "silent", "silk", "silly", "silver",
	"similar", "simple", "since", "sing", "siren", "sister", "situate", "six", "size", "skate", "sketch", "ski",
	"skill", "skin", "skirt", "skull", "slab", "slam", "sleep", "slender", "slice", "slide", "slight", "slim",
	"slogan", "slot", "slow", "slush", "small", "smart", "smile", "smoke", "smooth", "snack", "snake", "snap",
	"sniff", "snow", "soap", "soccer", "social", "sock", "soda", "soft", "solar", "soldier", "solid",
	"solution", "solve", "someone", "song", "soon", "sorry", "sort", "soul", "sound", "soup", "source", "south",
	"space", "spare", "spatial", "spawn", "speak", "special", "speed", "spell", "spend", "sphere", "spice",
	"spider", "spike", "spin", "spirit", "split", "spoil", "sponsor", "spoon", "sport", "spot", "spray",
	"spread", "spring", "spy", "square", "squeeze", "squirrel", "stable", "stadium", "staff", "stage", "stairs",
	"stamp", "stand", "start", "state", "stay", "steak", "steel", "stem", "step", "stereo", "stick", "still",
	"sting", "stock", "stomach", "stone", "stool", "story", "stove", "strategy", "street", "strike", "strong",
	"struggle", "student", "stuff", "stumble", "style", "subject", "submit", "subway", "success", "such",
	"sudden", "suffer", "sugar", "suggest", "suit", "summer", "sun", "sunny", "sunset", "super", "supply",
	"supreme", "sure", "surface", "surge", "surprise", "surround", "survey", "suspect", "sustain", "swallow",
	"swamp", "swap", "swarm", "swear", "sweet", "swift", "swim", "swing", "switch", "sword", "symbol",
	"symptom", "syrup", "system", "table", "tackle", "tag", "tail", "talent", "talk", "tank", "tape", "target",
	"task", "taste", "tattoo", "taxi", "teach", "team", "tell", "ten", "tenant", "tennis", "tent", "term",
	"test", "text", "thank", "that", "theme", "then", "theory", "there", "they", "thing", "this", "thought",
	"three", "thrive", "throw", "thumb", "thunder", "ticket", "tide", "tiger", "tilt", "timber", "time", "tiny",
	"tip", "tired", "tissue", "title", "toast", "tobacco", "today", "toddler", "toe", "together", "toilet",
	"token", "tomato", "tomorrow", "tone", "tongue", "tonight", "tool", "tooth", "top", "topic", "topple",
	"torch", "tornado", "tortoise", "toss", "total", "tourist", "toward", "tower", "town", "toy", "track",
	"trade", "traffic", "tragic", "train", "transfer", "trap", "trash", "travel", "tray", "treat", "tree",
	"trend", "trial", "tribe", "trick", "trigger", "trim", "trip", "trophy", "trouble", "truck", "true",
	"truly", "trumpet", "trust", "truth", "try", "tube", "tuition", "tumble", "tuna", "tunnel", "turkey",
	"turn", "turtle", "twelve", "twenty", "twice", "twin", "twist", "two", "type", "typical", "ugly",
	"umbrella", "unable", "unaware", "uncle", "uncover", "under", "undo", "unfair", "unfold", "unhappy",
	"uniform", "unique", "unit", "universe", "unknown", "unlock", "until", "unusual", "unveil", "update",
	"upgrade", "uphold", "upon", "upper", "upset", "urban", "urge", "usage", "use", "used", "useful", "useless",
	"usual", "utility", "vacant", "vacuum", "vague", "valid", "valley", "valve", "van", "vanish", "vapor",
	"various", "vast", "vault", "vehicle", "velvet", "vendor", "venture", "venue", "verb", "verify", "version",
	"very", "vessel", "veteran", "viable", "vibrant", "vicious", "victory", "video", "view", "village",
	"vintage", "violin", "virtual", "virus", "visa", "visit", "visual", "vital", "vivid", "vocal", "voice",
	"void", "volcano", "volume", "vote", "voyage", "wage", "wagon", "wait", "walk", "wall", "walnut", "want",
	"warfare", "warm", "warrior", "wash", "wasp", "waste", "water", "wave", "way", "wealth", "weapon", "wear",
	"weasel", "weather", "web", "wedding", "weekend", "weird", "welcome", "west", "wet", "whale", "what",
	"wheat", "wheel", "when", "where", "whip", "whisper", "wide", "width", "wife", "wild", "will", "win",
	"window", "wine", "wing", "wink", "winner", "winter", "wire", "wisdom", "wise", "wish", "witness", "wolf",
	"woman", "wonder", "wood", "wool", "word", "work", "world", "worry", "worth", "wrap", "wreck", "wrestle",
	"wrist", "write", "wrong", "yard", "year", "yellow", "you", "young", "youth", "zebra", "zero", "zone", "zoo"}
