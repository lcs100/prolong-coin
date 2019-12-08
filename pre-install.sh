# 在一台新服务器上部署，最好是外网服务器
# 如果服务器已经安装go和git，请注释go和git的下载部分
sudo apt install gcc
sudo apt install libc6-dev

cd $HOME
wget https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz
tar xzf go1.11.5.linux-amd64.tar.gz -C /usr/local
export PATH=$PATH:/usr/local/go/bin

cd prolong-coin/github.com/btcsuite/btcd

sh download.sh
