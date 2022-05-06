package logic

import (
	"context"
	"fmt"

	"github.com/r35krag0th/zero-gtfo-overlay/rundown/internal/svc"
	"github.com/r35krag0th/zero-gtfo-overlay/rundown/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRundownsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListRundownsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRundownsLogic {
	return &ListRundownsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListRundownsLogic) ListRundowns(req *types.ListRundownRequest) (resp *types.ListRundownResponse, err error) {
	resp = &types.ListRundownResponse{}
	for _, rundown := range l.svcCtx.Rundowns {
		var expeditions []types.ExpeditionURLReference
		for _, expedition := range rundown.Expeditions {
			expeditions = append(expeditions, types.ExpeditionURLReference{
				ID:   expedition.ID,
				Name: expedition.Name,
				// URL: fmt.Sprintf("%s://%s:%d/api/v1/rundowns/%s/expedition/%s", )
				URL: fmt.Sprintf("/api/v1/rundowns/%d/expedition/%s", rundown.ID, expedition.ID),
			})
		}

		resp.Rundowns = append(resp.Rundowns, types.RundownIndexItem{
			ID:          rundown.ID,
			Name:        rundown.Name,
			Expeditions: expeditions,
		})
	}
	return
}
