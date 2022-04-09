package logic

import (
	"context"
	"fmt"

	"github.com/r35krag0th/zero-gtfo-overlay/internal/data"
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
	overlayData, err := data.GetOverlayData(l.svcCtx.ConsulClient)
	if err != nil {
		return nil, err
	}

	currentRundown, ok := l.svcCtx.Rundowns[l.svcCtx.Config.CurrentRundown]
	if !ok {
		return nil, fmt.Errorf("current rundown (%s) has no data", l.svcCtx.Config.CurrentRundown)
	}

	expedition := currentRundown.GetExpedition(req.ID)
	if expedition == nil {
		return nil, fmt.Errorf("unknown expedition in Rundown %s: %v", l.svcCtx.Config.CurrentRundown, req.ID)
	}

	overlayData.Expedition.ID = req.ID
	overlayData.Expedition.Name = expedition.Name

	err = data.UpdateOverlayData(overlayData, l.svcCtx.ConsulClient)
	if err != nil {
		return nil, err
	}
	return &types.ExpeditionResponse{
		ID:   overlayData.Expedition.ID,
		Name: overlayData.Expedition.Name,
	}, nil
}
