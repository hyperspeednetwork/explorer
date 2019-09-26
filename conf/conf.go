package conf

import "time"

func NewConfig() Config {
	return Config{
		DBstring:       "mongodb://hsnnet:jiulian666@hsn_dev_mongo:27017/admin",
		DBName:         "hsnNet",
		GenesisAddress: "http://127.0.0.1:26678/genesis?",
		Remote: RpcLcd{
			Rpc: "http://127.0.0.1:26678",
			Lcd: "http://127.0.0.1:1317",
		},

		//DBstring:       "mongodb://hsnhub_dev_test:hsn_test@172.38.8.89:27890/hsnhub_db_dev_test", //test
		//DBName :"hsnhub_db_dev_test",//test
		//
		//GenesisAddress: "http://172.38.8.89:26678/genesis?",
		//Remote: RpcLcd{
		//	Rpc: "http://172.38.8.89:26678",
		//	Lcd: "http://172.38.8.89:1317",
		//},
		//DBstring:       "127.0.0.1:27017/test5", //test
		//DBName:         "test5",                //test
		//GenesisAddress: "http://172.38.8.89:26678/genesis?",
		//Remote: RpcLcd{
		//	Rpc: "http://172.38.8.89:26678",
		//	Lcd: "http://172.38.8.89:1317",
		//},

		Param: Params{ /* 爬虫间隔时间 deng*/
			HTTPGetTimeOut:         30,
			BlockInterval:          5,
			DelegationInterval:     300,
			ValidatorsSetsInterval: 300,
			TransactionsInterval:   120,
			DelegatorInterval:      600,
			CanNotGetErrorInterval: 10,
		},
		Public: Publics{
			ChainId:            "hsn",
			ChainName:          "hsn",
			CoinGeckoId:        "hsn",
			CoinToVoitingPower: 1.0,
			ValidatorsSetLimit: 100,
		},
	}

}

type Config struct {
	DBstring       string
	DBName         string
	GenesisAddress string
	Remote         RpcLcd
	Param          Params
	Public         Publics
}

type RpcLcd struct {
	Rpc string
	Lcd string
}

type Publics struct {
	ChainName string
	ChainId   string

	//Bech32PrefixAccAddr  string
	//Bech32PrefixAccPub   string
	//Bech32PrefixValAddr  string
	//Bech32PrefixValPub   string
	//Bech32PrefixConsAddr string
	//Bech32PrefixConPub   string
	//GasPrice           float32
	CoinGeckoId        string  // get price for coinGeckid
	CoinToVoitingPower float32 // 代币转换votingPower比例
	ValidatorsSetLimit int     // 验证100个区块的验证者集合
}
type Params struct {
	HTTPGetTimeOut         time.Duration
	BlockInterval          time.Duration
	DelegationInterval     time.Duration
	ValidatorsSetsInterval time.Duration
	TransactionsInterval   time.Duration
	DelegatorInterval      time.Duration
	CanNotGetErrorInterval time.Duration
}
