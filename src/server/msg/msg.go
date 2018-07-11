package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)

var (
	Processor = protobuf.NewProcessor()
)

func init() {
	Processor.SetByteOrder(true)
	Processor.Register(&StartFight{})
	Processor.Register(&FightResult{})
	Processor.Register(&EnterFight{})
	Processor.Register(&SignUpResponse{})
	Processor.Register(&TosChat{})
	Processor.Register(&TocChat{})
	Processor.Register(&Login{})
	Processor.Register(&PlayerBaseInfo{})
	Processor.Register(&LoginSuccessfull{})
	Processor.Register(&LoginFaild{})
}
