package logic

import (
	"context"

	"github.com/r35krag0th/zero-gtfo-overlay/rundown/internal/svc"
	"github.com/r35krag0th/zero-gtfo-overlay/rundown/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListExpeditionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListExpeditionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListExpeditionsLogic {
	return &ListExpeditionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListExpeditionsLogic) ListExpeditions(req *types.ListRundownExpeditionsRequest) (resp *types.ListRundownExpeditionsResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
