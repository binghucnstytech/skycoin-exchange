package main

import (
	"flag"
	"log"

	"github.com/skycoin/skycoin-exchange/src/server"
	"github.com/skycoin/skycoin/src/cipher"
)

var sk = "38d010a84c7b9374352468b41b076fa585d7dfac67ac34adabe2bbba4f4f6257"

func registerFlags(cfg *server.Config) {
	flag.IntVar(&cfg.Port, "port", 6060, "server listen port")
	flag.IntVar(&cfg.Fee, "fee", 10000, "transaction fee in satoish")
	flag.StringVar(&cfg.DataDir, "dataDir", "", "data directory")
	flag.StringVar(&cfg.WalletName, "wltName", "server.wlt", "server's wallet file name")
	flag.StringVar(&cfg.Seed, "s", "seed", "wallet's seed")
	flag.IntVar(&cfg.UtxoPoolSize, "poolsize", 1000, "utxo pool size")
}

func main() {
	cfg := server.Config{}
	registerFlags(&cfg)
	flag.Parse()

	key, err := cipher.SecKeyFromHex(sk)
	if err != nil {
		log.Fatal(err)
	}
	cfg.Seckey = key

	s := server.New(cfg)
	s.Run()
}
