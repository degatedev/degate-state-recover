package conf

import (
	"flag"
	"log"
	"staterecovery/common/entity"
	"time"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	Conf     = &Config{}
)

// Config get config from toml .
type Config struct {
	DataFrom                          int
	MySql                             *MySql
	StateSavedFile                    string
	ProtocolAccount                   string
	WhiteAccounts                     []*entity.WhiteAccount `json:"whiteAccounts"`
	LoopInterval                      int
	SaveStateBlockInterval            int
	LastL2BlockID                     int
	WithdrawModeAccount				  string
	WithdrawModeToken				  string
	WithdrawModeFilePath          	  string
	// state check
	StateCheckLoopInterval 			  int
	StateFileGetUrl 				  string
	PrometheusPort 					  string
	ExchangeContract            	  string
	FirstL1BlockID 					  uint64
	ChainNode               		  string
	BlockFilePath               	  string
}

type MySql struct {
	DSN         string
	Active      int           // pool
	Idle        int           // pool
	IdleTimeout time.Duration // connect max life time.
	Debug       bool
}

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
}

// Init init conf
func Init() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	if err != nil {
		log.Println("toml decode error:", err.Error())
		return err
	}
	return nil
}
