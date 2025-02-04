package transaction

import (
	"github.com/xiaoyueya/go-sdk/common/types"
	"github.com/xiaoyueya/go-sdk/types/msg"
	"github.com/xiaoyueya/go-sdk/types/tx"
)

type TransferTokenOwnershipResult struct {
	tx.TxCommitResult
}

func (c *client) TransferTokenOwnership(symbol string, newOwner types.AccAddress, sync bool, options ...Option) (*TransferTokenOwnershipResult, error) {
	fromAddr := c.keyManager.GetAddr()
	transferOwnershipMsg := msg.NewTransferOwnershipMsg(fromAddr, symbol, newOwner)
	commit, err := c.broadcastMsg(transferOwnershipMsg, sync, options...)
	if err != nil {
		return nil, err
	}
	return &TransferTokenOwnershipResult{*commit}, nil
}
