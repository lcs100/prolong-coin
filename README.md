名称
基于CPU物理资源利用率的比特币算法

环境：
go 1.11

拉下来之后，执行pre-install.sh
在执行完pre-install.sh之后，用户自行配置环境变量：

export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export GOBIN=$GOROOT/bin
export GOTOOLS=$GOROOT/pkg/tool
export PATH=$PATH:$GOBIN:$GOTOOLS

之后执行：

source ~/.bashrc

之后执行/github.com/btcsuite/btcd/download.sh
