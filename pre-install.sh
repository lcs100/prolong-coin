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

