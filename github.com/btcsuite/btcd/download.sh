cd ../
git clone git@github.com:btcsuite/btclog.git
git clone git@github.com:btcsuite/btcutil.git
git clone git@github.com:btcsuite/go-socks.git
git clone git@github.com:btcsuite/goleveldb.git
git clone git@github.com:btcsuite/websocket.git
git clone git@github.com:btcsuite/snappy-go.git
git clone git@github.com:btcsuite/winsvc.git
git clone git@github.com:btcsuite/btcwallet.git
git clone git@github.com:btcsuite/golangcrypto.git

cd ..
git clone git@github.com:davecgh/go-spew.git davecgh/go-spew
git clone git@github.com:jessevdk/go-flags.git jessevdk/go-flags
git clone git@github.com:kkdai/bstream.git kkdai/bstream
git clone git@github.com:aead/siphash.git aead/siphash
git clone git@github.com:jrick/logrotate.git jrick/logrotate
git clone git@github.com:etcd-io/bbolt.git coreos/bbolt
cd ..
mkdir -p golang.org/x && cd golang.org/x
git clone git@github.com:golang/crypto.git crypto
git clone git@github.com:golang/net.git net
git clone git@github.com:golang/sys.git sys
git clone git@github.com:golang/text.git text
cd ../../github.com
git clone git@github.com:golang/protobuf.git golang/protobuf
git clone git@github.com:lightninglabs/gozmq.git lightninglabs/gozmq
git clone git@github.com:lightninglabs/neutrino.git lightninglabs/neutrino
git clone git@github.com:lightningnetwork/lnd.git lightningnetwork/lnd

cd ..
git clone git@github.com:grpc/grpc-go.git google.golang.org/grpc
git clone git@github.com:google/go-genproto.git google.golang.org/genproto

# compile btcwallet
cd github.com/btcsuite/btcwallet
echo "compiling btcwallet..."
go install . ./cmd/...

# compile btcd
cd ../btcd
echo "compiling btcd"
go install . ./cmd/...


# copy conf to HOME
cd $HOME
mkdir .btcd
cp $HOME/prolong-coin/github.com/btcsuite/btcd/sample-btcd.conf ~/.btcd/btcd.conf

cd $HOME
mkdir .btcwallet
cp $HOME/prolong-coin/github.com/btcsuite/btcwallet/sample-btcwallet.conf ~/.btcwallet/btcwallet.conf