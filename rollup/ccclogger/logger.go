package logger

import (
	"math/big"
	"time"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/vm"
)

var _ vm.EVMLogger = (*CCCLogger)(nil)

type CCCLogger struct {
	// todo: keep track of circuit resources
}

func (l *CCCLogger) CaptureStart(env *vm.EVM, from common.Address, to common.Address, create bool, input []byte, gas uint64, value *big.Int) {

}

func (l *CCCLogger) CaptureState(pc uint64, op vm.OpCode, gas, cost uint64, scope *vm.ScopeContext, rData []byte, depth int, err error) {

}

func (l *CCCLogger) CaptureStateAfter(pc uint64, op vm.OpCode, gas, cost uint64, scope *vm.ScopeContext, rData []byte, depth int, err error) {

}

func (l *CCCLogger) CaptureEnter(typ vm.OpCode, from common.Address, to common.Address, input []byte, gas uint64, value *big.Int) {

}

func (l *CCCLogger) CaptureExit(output []byte, gasUsed uint64, err error) {

}

func (l *CCCLogger) CaptureFault(pc uint64, op vm.OpCode, gas, cost uint64, scope *vm.ScopeContext, depth int, err error) {

}

func (l *CCCLogger) CaptureEnd(output []byte, gasUsed uint64, t time.Duration, err error) {

}

// ShouldRevertAndCommit returns true if last executed transaction triggered an overflow
// Caller should revert the last transaction and close the block
func (l *CCCLogger) ShouldRevertAndClose() bool {
	return false
}
