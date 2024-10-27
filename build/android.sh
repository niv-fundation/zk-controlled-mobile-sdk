set -e

go get -u golang.org/x/mobile/bind

export PATH=$PATH:~/go/bin
gomobile init

gomobile bind -target=android -o Ethereum.aar -androidapi 24