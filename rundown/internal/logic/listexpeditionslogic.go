package logic

import (
	"context"
	"fmt"

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
	data, found := l.svcCtx.Rundowns[req.RundownID]
	if !found {
		err = fmt.Errorf("unknown expedition")
		return
	}

	// ID, Name, URL
	resp = &types.ListRundownExpeditionsResponse{
		// Expeditions: *l.svcCtx.Rundowns[req.RundownID],
		Expeditions: nil,
	}

	for _, expedition := range data.Expeditions {
		resp.Expeditions = append(resp.Expeditions, types.ExpeditionURLReference{
			ID:   expedition.ID,
			Name: expedition.Name,
			URL:  "",
		})
	}
	return
}
