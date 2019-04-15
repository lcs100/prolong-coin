package main

import (
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"io/ioutil"
	"path/filepath"
	"fmt"
)


func main() {

	btcdHomeDir := btcutil.AppDataDir("btcd", false)
	certs, err := ioutil.ReadFile(filepath.Join(btcdHomeDir, "rpc.cert"))

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

 	data := make(map[string]int)
	address := make(map[string]string)
	address["SgyPztP3aGG9KRKrnhGK67Z7wbsaGwqQRE"] = "VM"
	address["SNTsvpxn3ryxhvP749t3ykBHZCUhDaCeeb"] = "HuaWei"
	address["SZ68SFfPvWfAxycFGWiAoc6PkDgfzkbfoJ"] = "Tencent"
	address["SVcRE7hfK6db2VZhzGPjiSn9VjofkUSG8t"] = "Linode"
	address["SanecaPK6oW8ywAJGi5ZccHnV3FvgYthM3"] = "Vultr(140)"
	// Query the RPC server for the current block count and display it.
	blockCount ,_ := client.GetBlockCount()
	for blockId := int64(1); blockId<= blockCount;blockId ++ {
		blockHash,_ := client.GetBlockHash(blockId)
		block,_ := client.GetBlockVerbose(blockHash)
	//	nonce := block.Nonce
	//	bits := block.Bits
	//	timeStamp := block.Time
		txHash,_ := chainhash.NewHashFromStr(block.Tx[0])
		tx,_:= client.GetRawTransactionVerbose(txHash)
		miner := tx.Vout[0].ScriptPubKey.Addresses[0]
		//fmt.Println(nonce,bits,timeStamp,miner)
		data[address[miner]] +=1
	}
	fmt.Println(data)
	defer client.Shutdown()
}
