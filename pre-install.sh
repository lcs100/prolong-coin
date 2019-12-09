# based on Ubuntu 16.04
# 在一台新服务器上部署，最好是外网服务器
# 如果服务器已经安装go和gcc，请注释相关下载部分

# 安装gcc和相关库
sudo apt install gcc
sudo apt install libc6-dev

# 下载和安装go
cd $HOME
wget https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz
tar xzf go1.11.5.linux-amd64.tar.gz -C /usr/local

# 修改bashrc写入相关go环境变量
echo " #go Path setting" >> ~/.bashrc
echo "export GOPATH=$HOME/go" >> ~/.bashrc
echo "export GOROOT=/usr/local/go" >> ~/.bashrc
GOROOT=$HOME/go
echo "export GOBIN=$GOROOT/bin" >> ~/.bashrc
echo "export GOTOOLS=$GOROOT/pkg/tool" >> ~/.bashrc
GOTOOLS=$GOROOT/pkg/tool
echo "export PATH=$PATH:$GOBIN:$GOTOOLS" >> ~/.bashrc

source ~/.bashrc

# 执行另外一个脚本
cd prolong-coin/github.com/btcsuite/btcd
bash download.sh
