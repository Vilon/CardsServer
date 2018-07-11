package internal

import (
	"reflect"
	"server/msg"

	"github.com/name5566/leaf/log"

	"github.com/name5566/leaf/gate"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.Login{}, handleAuth)
}

func handleAuth(args []interface{}) {
	log.Debug("1111111")
	m := args[0].(*msg.Login)
	a := args[1].(gate.Agent)

	if len(m.Account) < 2 || len(m.Account) > 12 {
		a.WriteMsg(&msg.LoginFaild{Code: msg.LoginFaild_AccIDInvalid})
		return
	}
	a.WriteMsg(&msg.LoginSuccessfull{})
	log.Debug("222222222222")
}
