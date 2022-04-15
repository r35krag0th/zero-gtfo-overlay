package logic

import (
	"context"

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
	// for _, rundown := range l.svcCtx.R
	// resp = &types.ListRundownResponse{
	//
	// }
	// resp = &types.ListRundownResponse{
	// 	Rundowns: nil,
	// }
	// for _, rundown := range l.svcCtx.Rundowns {
	// 	resp.Rundowns = append(resp.Rundowns, &types.RundownIndexItem{
	// 		ID:          strconv.Atoi(rundown.ID),
	// 		Name:        "",
	// 		Expeditions: nil,
	// 	})
	// }
	return
}
