package skycoin_exchange

import (
	"github.com/skycoin/skycoin/src/cipher"
	"net/http"
)

/*
	Load privatekey from json file
	- this is the clients identity
*/



	
//get balance of client and deposit addresses
func (self *client) GetStatus() {

}

func (fself )
type Client struct {
	Seckey cipher.SecKey
	Server string //remote server to querry
}

//handle events
func (self *Server) eventHandler(w http.ResponseWriter, r *http.Request) {



func (self *Client) RunWebserver() {

	host := "localhost"
	fmt.Printf("host: %v\n", host)

	port := int(8081)
	fmt.Printf("port: %v\n", port)
	addr := fmt.Sprintf("%s:%d", host, port)

	mux := http.NewServeMux()

	mux.Handle("/event", http.HandlerFunc(eventHandler))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "")
	})

	err := http.ListenAndServe(addr, mux)
	fmt.Println(err.Error())
}
