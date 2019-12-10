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
