package conf

import "time"

// 		DBstring:       "mongodb://root:helloKitty@127.0.0.1:27017",
func NewConfig() Config {
	return Config{
		//DBstring:       "mongodb://hsnhub_dev:hsn@172.38.8.89:27890/hsnhub_db_dev",
		//DBName :"hsnhub_db_dev",
		//DBstring:       "mongodb://hsnhub_dev_test:hsn_test@172.38.8.89:27890/hsnhub_db_dev_test", //test
		//DBName :"hsnhub_db_dev_test",//test

		DBstring:       "127.0.0.1:27017/test", //test
		DBName:         "test",                //test
		GenesisAddress: "http://172.38.8.89:26678/genesis?",
		Remote: RpcLcd{
			Rpc: "http://172.38.8.89:26678",
			Lcd: "http://172.38.8.89:1317",
		},
		Param: Params{ /* 爬虫间隔时间 deng*/
			HTTPGetTimeOut:         30,
			StartHeight:            0,
			DefaultBLockTime:       6,
			BlockInterval:          5,
			ConsensusInterval:      1,
			StatusInterval:         7,
			SingingInfoInterval:    1800,
			ProposalInterval:       5,
			MissedBlockInterval:    60,
			DelegationInterval:     300,
			ValidatorsSetsInterval: 30,
			TransactionsInterval:   50,
			DelegatorInterval:      600,
		},
		Public: Publics{
			ChainId:            "hsn",
			ChainName:          "hsn",
			CoinGeckoId:        "hsn",
			CoinToVoitingPower: 1000000.0,
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
	StakingDenom       string
	MintingDenom       string
	StakingFraction    int
	GasPrice           float32
	CoinGeckoId        string  // get price for coinGeckid
	CoinToVoitingPower float32 // 代币转换votingPower比例
	ValidatorsSetLimit int     // 验证100个区块的验证者集合
}
type Params struct {
	HTTPGetTimeOut         time.Duration
	StartHeight            time.Duration
	DefaultBLockTime       time.Duration
	BlockInterval          time.Duration
	ConsensusInterval      time.Duration
	StatusInterval         time.Duration
	SingingInfoInterval    time.Duration
	ProposalInterval       time.Duration
	MissedBlockInterval    time.Duration
	DelegationInterval     time.Duration
	ValidatorsSetsInterval time.Duration
	TransactionsInterval   time.Duration
	DelegatorInterval      time.Duration
}
