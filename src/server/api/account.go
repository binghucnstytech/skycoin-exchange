package api

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/skycoin/skycoin-exchange/src/pp"
	"github.com/skycoin/skycoin-exchange/src/server/engine"
)

func getRequest(c *gin.Context, out interface{}) error {
	d := c.MustGet("rawdata").([]byte)
	return json.Unmarshal(d, out)
}

// CreateAccount create account with specific pubkey,
func CreateAccount(ee engine.Exchange) gin.HandlerFunc {
	return func(c *gin.Context) {
		errRlt := &pp.EmptyRes{}
		for {
			req := pp.CreateAccountReq{}
			if err := getRequest(c, &req); err != nil {
				glog.Error(err)
				errRlt = pp.MakeErrResWithCode(pp.ErrCode_WrongRequest)
				break
			}

			pubkey := pp.BytesToPubKey(req.GetPubkey())
			if err := pubkey.Verify(); err != nil {
				glog.Error(err)
				errRlt = pp.MakeErrResWithCode(pp.ErrCode_WrongAccountId)
				break
			}

			// create account with pubkey.
			if _, err := ee.CreateAccountWithPubkey(pubkey); err != nil {
				glog.Error(err)
				errRlt = pp.MakeErrResWithCode(pp.ErrCode_WrongRequest)
				break
			}

			res := pp.CreateAccountRes{
				AccountId: req.Pubkey,
				CreatedAt: pp.PtrInt64(time.Now().Unix()),
			}

			reply(c, res)
			return
		}

		c.JSON(200, *errRlt)
	}
}
