package main

import (
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
	//"github.com/btcsuite/btcd/chaincfg/chainhash"
	"io/ioutil"
	"log"
	"path/filepath"
	"fmt"
	"database/sql"
	_"mysql"
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

func insert(db *sql.DB,blockId int64,blockHash string,miner string, nonce uint32, bits string,timeStamp int64) {
	fmt.Println(miner)
	stmt, err := db.Prepare("INSERT INTO block(blockId, blockHash,miner,nonce,bits,createTime) VALUES(?,?,?,?,?,?);")
	res, err := stmt.Exec(blockId,blockHash,miner,nonce,bits,timeStamp)
	fmt.Println(res,err)
	if err != nil {
		log.Fatal(err,res)
	}
	//fmt.Println(res)
}
func main() {

	db,_ := sql.Open("mysql", "root:root@/blockchain")
	//Clear(db)
	btcdHomeDir := btcutil.AppDataDir("btcd", false)
	certs, err := ioutil.ReadFile(filepath.Join(btcdHomeDir, "rpc.cert"))
	if err != nil {
		log.Fatal(err)
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


	// Query the RPC server for the current block count and display it.
	//blockCount ,_ := client.GetBlockCount()
	for blockId := int64(1); blockId<= 1;blockId ++ {
		blockHash,_ := client.GetBlockHash(blockId)
		block,_ := client.GetBlockVerbose(blockHash)
		nonce := block.Nonce
		bits := block.Bits
		timeStamp := block.Time
		//txHash,_ := chainhash.NewHashFromStr(block.Tx[0])
		//tx,_:= client.GetRawTransactionVerbose(txHash)
		//miner := tx.Vout[0].ScriptPubKey.Addresses[0]
		insert(db,blockId,block.Hash,block.Hash,nonce,bits,timeStamp)
	}

	defer client.Shutdown()
}
