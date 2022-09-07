package message

import "github.com/waro163/wechat-message/utils"

// TransferCustomer 转发客服消息
type TransferCustomer struct {
	CommonMsg
}

// NewTransferCustomer 实例化TransferCustomer
func NewTransferCustomer() *TransferCustomer {
	tc := new(TransferCustomer)
	tc.SetMsgType(MsgTypeTransfer)
	tc.SetCreateTime(utils.GetCurrTS())
	return tc
}

// TransferAssignCustomer转发到指定客服
type TransferAssignCustomer struct {
	CommonMsg
	KfAccount CDATA `xml:"TransInfo>KfAccount"`
}

// NewTransferAssignCustomer 实例化
func NewTransferAssignCustomer(kfAccount string) *TransferAssignCustomer {
	tac := new(TransferAssignCustomer)
	tac.KfAccount = CDATA(kfAccount)
	tac.SetMsgType(MsgTypeTransfer)
	tac.SetCreateTime(utils.GetCurrTS())
	return tac
}
