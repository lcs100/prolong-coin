名称
======
基于CPU物理资源利用率的比特币算法

配置环境：
----
    go 1.11
    gcc
    ubuntu 16.04及以上

编译步骤:
-----
首先

    cd prolong-coin/
    bash pre-install.sh

之后，用户自行配置环境变量

    export GOPATH=$HOME/go
    export GOROOT=/usr/local/go
    export GOBIN=$GOROOT/bin
    export GOTOOLS=$GOROOT/pkg/tool
    export PATH=$PATH:$GOBIN:$GOTOOLS

之后执行

    source ~/.bashrc

切换目录到prolong-coin/github.com/btcsuite/btcd，执行

    bash download.sh


单节点上运行
=====

1、新建钱包

    btcwallet --simnet --create

新建密码并牢记，这样会新建一个钱包

2、开启btcd

    btcd --simnet --rpcuser=rpcuser --rpcpass=rpcpass

3、保持终端，另外开启一个终端，让钱包和btcd进行连接

    btcwallet --simnet --username=rpcuser --password=rpcpass

4、打开第三个终端，执行以下命令，解锁钱包：

    btcctl --simnet --rpcuser=rpcuser --rpcpass=rpcpass --wallet walletpassphrase "lcs" 180

创建新账户：

    btcctl --simnet --rpcuser=rpcuser --rpcpass=rpcpass --wallet createnewaccount lcsnew

创建新地址：

    btcctl --simnet --rpcuser=rpcuser --rpcpass=rpcpass --wallet getnewaddress lcsnew

记住新建的地址，比较重要

5、利用新地址进行挖矿

    btcd --simnet --rpcuser=rpcuser --rpcpass=rpcpass --miningaddr= --txindex

6、查看某个账户的地址

    btcctl --simnet --rpcuser=rpcuser --rpcpass=rpcpass --wallet getaddressesbyaccount default
