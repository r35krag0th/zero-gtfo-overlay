package logic

import (
	"context"
	"fmt"
	"strings"

	"github.com/r35krag0th/zero-gtfo-overlay/internal/data"
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
	overlayData, err := data.GetOverlayData(l.svcCtx.ConsulClient)
	if err != nil {
		return nil, err
	}

	currentRundown, ok := l.svcCtx.Rundowns[l.svcCtx.Config.CurrentRundown]
	if !ok {
		return nil, fmt.Errorf("current rundown (%s) has no data", l.svcCtx.Config.CurrentRundown)
	}

	expedition := currentRundown.GetExpedition(overlayData.Expedition.ID)
	if expedition == nil {
		return nil, fmt.Errorf("unknown expedition in Rundown %s: %v", l.svcCtx.Config.CurrentRundown, overlayData.Expedition.ID)
	}

	if sector := expedition.GetSector(req.Name); sector == nil {
		return nil, fmt.Errorf(
			"invalid sector '%v', available sectors are: %s",
			req.Name,
			strings.Join(expedition.AvailableSectors(), ","),
		)
	}

	overlayData.Sector = req.Name
	err = data.UpdateOverlayData(overlayData, l.svcCtx.ConsulClient)
	if err != nil {
		return nil, err
	}
	return &types.SectorResponse{
		Name: overlayData.Sector,
	}, nil
}
