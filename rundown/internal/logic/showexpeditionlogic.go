package logic

import (
	"context"

	"github.com/r35krag0th/zero-gtfo-overlay/rundown/internal/svc"
	"github.com/r35krag0th/zero-gtfo-overlay/rundown/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowExpeditionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowExpeditionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowExpeditionLogic {
	return &ShowExpeditionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowExpeditionLogic) ShowExpedition(req *types.ShowExpeditionRequest) (resp *types.ShowExpeditionResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
