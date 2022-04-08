package logic

import (
	"context"

	"github.com/r35krag0th/zero-gtfo-overlay/internal/svc"
	"github.com/r35krag0th/zero-gtfo-overlay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateExpeditionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateExpeditionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateExpeditionLogic {
	return &UpdateExpeditionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateExpeditionLogic) UpdateExpedition(req *types.UpdateExpeditionRequest) (resp *types.ExpeditionResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
