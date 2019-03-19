import os
import re
from bitcoin.rpc import RawProxy
def RunCommand(command,arg):
    return os.popen("btcctl --simnet --wallet --rpcuser=hht --rpcpass=hht "+ command+" "+str(arg)).readlines()

def calPreZero(num):
    for j in range(256):
        if (1<<j) >= num:
            return 256-j


def getBlockCount():
    return int(RunCommand("getblockcount","")[0].strip())

def getBlock(data):
    block = {}
    print data
    for i in range(len(data)):
        value = str(data[i])
        value = re.sub(re.compile(r'"|,|\s*|'),"",value)
        if len(value) > 4:
            index = value.find(":")
            block[value[:index]] = value[index+1:]
    getTransaction(block['tx'])
    return block

def getTransaction(txid):
    tx = RunCommand(getrawtransaction,txid)
    decodeTx = Runcommand(decoderawtransaction,tx)
    print decodeTx
if __name__ == '__main__':
    '''
    data = ['{\n', '  "hash": "42c6edfaebd12bdf714080dd16926eb99bc3c87a7f0d4bedd6960dac19f7cd82",\n', '  "confirmations": 23,\n', '  "strippedsize": 285,\n', '  "size": 285,\n', '  "weight": 1140,\n', '  "height": 0,\n', '  "version": 1,\n', '  "versionHex": "00000001",\n', '  "merkleroot": "4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b",\n', '  "tx": [\n', '    "4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b"\n', '  ],\n', '  "time": 1401292357,\n', '  "nonce": 2,\n', '  "bits": "2000000f",\n', '  "difficulty": 1,\n', '  "previousblockhash": "0000000000000000000000000000000000000000000000000000000000000000",\n', '  "nextblockhash": "0000085a5c9b14b32e2fc188ce3398cdbffa0776bc44f9e2111f87ff74f308db"\n', '}\n']
    data = "".join(data)
    data = re.sub(r'"|\s*|\n',"",data)
    print data
    block_list = []
    blockCount = getBlockCount()
    for i in range(1):
        value = RunCommand("getblock",RunCommand("getblockhash",i)[0].strip())
        block_list.append(getBlock(value))
    print block_list
    '''
    p = RawProxy()
    info = p.getblockchaininfo()
    print(info['blocks'])