package server

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/codahale/chacha20"
	"github.com/gin-gonic/gin"
	"github.com/skycoin/skycoin-exchange/src/server/account"
	"github.com/skycoin/skycoin-exchange/src/server/wallet"
	"github.com/skycoin/skycoin/src/cipher"
)

// CaseHandler represents one test case, which will be invoked by MockServer.
type CaseHandler func() (*httptest.ResponseRecorder, *http.Request)

// MockServer mock server state for various test cases,
// people can fake the server's state by fullfill the Server interface, and
// define various request cases by defining functions that match the signature of
// CaseHandler.
func MockServer(svr Server, fs CaseHandler) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	router := NewRouter(svr)
	w, r := fs()
	router.ServeHTTP(w, r)
	return w
}

// HttpRequestCase is used to create REST api test cases.
func HttpRequestCase(method string, url string, body io.Reader) CaseHandler {
	return func() (*httptest.ResponseRecorder, *http.Request) {
		w := httptest.NewRecorder()
		r, err := http.NewRequest(method, url, body)
		if err != nil {
			panic(err)
		}
		switch method {
		case "POST":
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}
		return w, r
	}
}

// FakeAccount for mocking various account state.
type FakeAccount struct {
	ID      string
	WltID   string
	Addr    string
	Nk      account.NonceKey
	Balance uint64
}

// FakeServer for mocking various server status.
type FakeServer struct {
	A account.Accounter
}

func (fa FakeAccount) GetWalletID() string {
	return fa.WltID
}

func (fa FakeAccount) GetAccountID() account.AccountID {
	d, err := cipher.PubKeyFromHex(fa.ID)
	if err != nil {
		panic(err)
	}
	return account.AccountID(d)
}

func (fa FakeAccount) GetNewAddress(ct wallet.CoinType) string {
	return fa.Addr
}

func (fa FakeAccount) GetBalance(ct wallet.CoinType) uint64 {
	return fa.Balance
}

func (fa *FakeAccount) SetNonceKey(nk account.NonceKey) {
	fa.Nk = nk
}

func (fa FakeAccount) GetNonceKey() account.NonceKey {
	return fa.Nk
}

func (fa FakeAccount) Encrypt(r io.Reader) ([]byte, error) {
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return []byte{}, err
	}

	data := make([]byte, len(d))
	c, err := chacha20.New(fa.Nk.Key, fa.Nk.Nonce)
	if err != nil {
		return []byte{}, err
	}
	c.XORKeyStream(data, d)
	return data, nil
}

func (fa FakeAccount) Decrypt(r io.Reader) ([]byte, error) {
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return []byte{}, err
	}

	data := make([]byte, len(d))
	c, err := chacha20.New(fa.Nk.Key, fa.Nk.Nonce)
	if err != nil {
		return []byte{}, err
	}
	c.XORKeyStream(data, d)
	return data, nil
}

func (fa FakeAccount) IsExpired() bool {
	d := time.Now().Unix() - fa.Nk.Expire_at.Unix()
	return d >= 0
}

func (fa FakeAccount) GenerateWithdrawTx(coins uint64, coinType wallet.CoinType) ([]byte, error) {
	return []byte{}, nil
}

func (fs *FakeServer) CreateAccountWithPubkey(pk cipher.PubKey) (account.Accounter, error) {
	if fs.A.GetWalletID() == "" {
		return nil, fmt.Errorf("create wallet failed")
	}
	return fs.A, nil
}

func (fs *FakeServer) GetAccount(id account.AccountID) (account.Accounter, error) {
	if fs.A != nil && fs.A.GetAccountID() == id {
		return fs.A, nil
	}
	return nil, errors.New("account not found")
}

func (fs *FakeServer) Run() {

}

func (fs FakeServer) GetNonceKeyLifetime() time.Duration {
	return time.Second * time.Duration(10*60)
}
