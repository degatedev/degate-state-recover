package data

import (
	"errors"
	"staterecovery/conf"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gogather/com/log"
	"github.com/jinzhu/gorm"
)

type DBDataResource struct {
	queryDB   *gorm.DB
}

type BlockTransaction struct {
	BlockId int64
	Input []byte
	BatchId int64
}

func initDB(config *conf.MySql, debug bool) (db *gorm.DB) {
	var err error
	db, err = gorm.Open("mysql", config.DSN)
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(config.Idle)
	db.DB().SetMaxOpenConns(config.Active)
	db.DB().SetConnMaxLifetime(config.IdleTimeout)

	if debug {
		db = db.Debug()
	}

	if err = db.DB().Ping(); err != nil {
		panic(err)
	}
	return
}

func newDBDataResource(conf *conf.Config) *DBDataResource {
	return &DBDataResource{
		queryDB: initDB(conf.MySql, conf.MySql.Debug),
	}
}

func (dr *DBDataResource) GetDataByBlockId(blockId int) (data *BlockData, err error) {
	blockData := &BlockTransaction{}
	err = dr.queryDB.Table("block_transactions").Where("block_id=? and status=3", blockId).First(blockData).Error
	if err != nil {
		log.Println("in getDataByBlockId query error:", err.Error())
		return nil, err
	}
	return &BlockData{
		BlockId: blockData.BlockId,
		Data: blockData.Input,
	}, nil
}

func (dr *DBDataResource) GetTokenIDByAddress(tokenAddr string) (tokenID string, err error) {
	return "0", errors.New("in DBDataResource GetTokenIDByAddress not implemented")
}