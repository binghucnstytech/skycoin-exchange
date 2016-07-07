package bitcoin_interface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Vout struct {
	Addr  string `json:"addr"`
	Value int    `json:"value"`
}

type TData struct {
	Address string `json:"address"`
	Key     string `json:"key"`
	OutAddr []Vout `json:"vout"`
	Fee     int64  `json:"fee"`
}

func TestNewRawTransaction(t *testing.T) {
	d, err := ioutil.ReadFile("test.json")
	assert.Nil(t, err)
	td := TData{}
	err = json.Unmarshal(d, &td)
	assert.Nil(t, err)

	utxos := GetUnspentOutputsBlkChnInfo(td.Address)
	outAddr := make([]OutAddr, len(td.OutAddr))
	for i, o := range td.OutAddr {
		outAddr[i].Addr = o.Addr
		outAddr[i].Value = int64(o.Value)
	}

	utks := make([]BlkChnUtxoWithkey, len(utxos))
	for i, utxo := range utxos {
		bk := BlkChnUtxoWithkey{
			BlkChnUtxo: utxo,
			Privkey:    td.Key,
		}
		utks[i] = bk
	}

	tx, err := NewTransaction(utks, outAddr)
	assert.Nil(t, err)
	b := bytes.Buffer{}
	tx.Serialize(&b)
	fmt.Printf("%x\n", b.Bytes())
}
