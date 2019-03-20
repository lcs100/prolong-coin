package main

import (
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"io/ioutil"
	"log"
	"path/filepath"
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "root"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "blockchain"
)
func Clear(db *sql.DB) {
	stmt, _ := db.Prepare("delete from block")
	stmt.Exec()
}

func Insert(db *sql.DB,blockId int64,blockHash string,miner string, nonce uint32, bits string,timeStamp int64) {
	stmt, err := db.Prepare("INSERT INTO block(blockId, blockHash,miner,nonce,bits,createTime) VALUES(?, ?,?,?,?,?);")
	res, err := stmt.Exec(blockId,blockHash,miner,nonce,bits,timeStamp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
func main() {

	db,_ := sql.Open("mysql", "root:root@/blockchain")
	Clear(db)
	btcdHomeDir := btcutil.AppDataDir("btcd", false)
	certs, err := ioutil.ReadFile(filepath.Join(btcdHomeDir, "rpc.cert"))
	if err != nil {
		//log.Fatal(err)
	}

	// Create a new RPC client using websockets.  Since this example is
	// not long-lived, the connection will be closed as soon as the program
	// exits.
	connCfg := &rpcclient.ConnConfig{
		Host:         "localhost:18556",
		Endpoint:     "ws",
		User:         "hht",
		Pass:         "hht",
		Certificates: certs,
	}
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		//log.Fatal(err)
	}
	defer client.Shutdown()

	// Query the RPC server for the current block count and display it.
	//blockCount ,_ := client.GetBlockCount()
	for blockId := int64(0); blockId< 0;blockId ++ {
		blockHash,_ := client.GetBlockHash(blockId)
		block,_ := client.GetBlockVerbose(blockHash)
		nonce := block.Nonce
		bits := block.Bits
		timeStamp := block.Time
		txHash,_ := chainhash.NewHashFromStr(block.Tx[0])
		txOut,_ := client.GetTxOut(txHash,0,true)
		fmt.Println(txOut)
		miner := txOut.ScriptPubKey.Addresses[0]
		Insert(db,blockId,block.Hash,miner,nonce,bits,timeStamp)
		fmt.Println(blockId,block.Hash,miner,nonce,bits,timeStamp)
	}


}