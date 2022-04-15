package handler

import (
	"net/http"

	"github.com/r35krag0th/zero-gtfo-overlay/overlay/internal/logic"
	"github.com/r35krag0th/zero-gtfo-overlay/overlay/internal/svc"
	"github.com/r35krag0th/zero-gtfo-overlay/overlay/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateExpeditionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateExpeditionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateExpeditionLogic(r.Context(), svcCtx)
		resp, err := l.UpdateExpedition(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
