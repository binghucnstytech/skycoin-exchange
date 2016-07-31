package order

import (
	"os"
	"path/filepath"

	"github.com/skycoin/skycoin/src/util"
)

type Type uint8

const (
	Bid Type = iota
	Ask
)

var (
	orderDir string = filepath.Join(util.UserHome(), ".skycoin-exchange/orderbook")
	orderExt string = "ods"
	idExt    string = "id"
)

type Order struct {
	ID          uint64 `json:"id"` // order id.
	AccountID   string `json:"account_id"`
	Type        Type   `json:"type"`         // order type.
	Price       uint64 `json:"price"`        // price of this order.
	Amount      uint64 `json:"amount"`       // total amount of this order.
	RestAmt     uint64 `json:"reset_amt"`    // rest amount.
	CreatedTime uint64 `json:"created_time"` // created time of the order.
}

type byPriceThenTimeDesc []Order
type byPriceThenTimeAsc []Order
type byOrderID []Order

func (bp byPriceThenTimeDesc) Len() int {
	return len(bp)
}

func (bp byPriceThenTimeDesc) Less(i, j int) bool {
	a := bp[i]
	b := bp[j]
	if a.Price > b.Price {
		return true
	} else if a.Price == b.Price {
		return a.CreatedTime > b.CreatedTime
	}
	return false
}

func (bp byPriceThenTimeDesc) Swap(i, j int) {
	bp[i], bp[j] = bp[j], bp[i]
}

func (bp byPriceThenTimeAsc) Len() int {
	return len(bp)
}

func (bp byPriceThenTimeAsc) Less(i, j int) bool {
	a := bp[i]
	b := bp[j]
	if a.Price < b.Price {
		return true
	} else if a.Price == b.Price {
		return a.CreatedTime > b.CreatedTime
	}
	return false
}

func (bp byPriceThenTimeAsc) Swap(i, j int) {
	bp[i], bp[j] = bp[j], bp[i]
}

func (bo byOrderID) Len() int {
	return len(bo)
}

func (bo byOrderID) Less(i, j int) bool {
	return bo[i].ID > bo[j].ID
}

func (bo byOrderID) Swap(i, j int) {
	bo[i], bo[j] = bo[j], bo[i]
}

func InitDir(path string) {
	if path == "" {
		path = orderDir
	} else {
		orderDir = path
	}
	// create the account dir if not exist.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			panic(err)
		}
	}
}
