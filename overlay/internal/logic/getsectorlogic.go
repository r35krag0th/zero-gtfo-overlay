package logic

import (
	"context"

	"github.com/r35krag0th/zero-gtfo-overlay/internal"
	"github.com/r35krag0th/zero-gtfo-overlay/overlay/internal/svc"
	"github.com/r35krag0th/zero-gtfo-overlay/overlay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSectorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSectorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSectorLogic {
	return &GetSectorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSectorLogic) GetSector(req *types.GetSectorRequest) (resp *types.SectorResponse, err error) {
	overlayData, err := internal.GetOverlayData(l.svcCtx.ConsulClient)
	if err != nil {
		return nil, err
	}

	return &types.SectorResponse{
		Name: overlayData.Sector,
	}, nil
}
