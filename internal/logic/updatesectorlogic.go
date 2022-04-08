package logic

import (
	"context"

	"github.com/r35krag0th/zero-gtfo-overlay/internal/svc"
	"github.com/r35krag0th/zero-gtfo-overlay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSectorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSectorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSectorLogic {
	return &UpdateSectorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSectorLogic) UpdateSector(req *types.UpdateSectorRequest) (resp *types.SectorResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
