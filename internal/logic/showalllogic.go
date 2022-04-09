package logic

import (
	"context"

	"github.com/r35krag0th/zero-gtfo-overlay/internal/data"
	"github.com/r35krag0th/zero-gtfo-overlay/internal/svc"
	"github.com/r35krag0th/zero-gtfo-overlay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowAllLogic {
	return &ShowAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowAllLogic) ShowAll(req *types.ShowRequest) (resp *types.ShowResponse, err error) {
	overlayData, err := data.GetOverlayData(l.svcCtx.ConsulClient)
	if err != nil {
		return nil, err
	}

	return &types.ShowResponse{
		Expedition: overlayData.Expedition,
		Sector:     types.SectorResponse{Name: overlayData.Sector},
	}, nil
}
