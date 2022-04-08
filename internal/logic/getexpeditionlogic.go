package logic

import (
	"context"

	"github.com/r35krag0th/zero-gtfo-overlay/internal/svc"
	"github.com/r35krag0th/zero-gtfo-overlay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExpeditionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetExpeditionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExpeditionLogic {
	return &GetExpeditionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetExpeditionLogic) GetExpedition(req *types.GetExpeditionRequest) (resp *types.ExpeditionResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
