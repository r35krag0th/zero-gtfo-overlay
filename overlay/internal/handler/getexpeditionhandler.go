package handler

import (
	"net/http"

	"github.com/r35krag0th/zero-gtfo-overlay/overlay/internal/logic"
	"github.com/r35krag0th/zero-gtfo-overlay/overlay/internal/svc"
	"github.com/r35krag0th/zero-gtfo-overlay/overlay/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetExpeditionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetExpeditionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetExpeditionLogic(r.Context(), svcCtx)
		resp, err := l.GetExpedition(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
