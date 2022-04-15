package handler

import (
	"net/http"

	"github.com/r35krag0th/zero-gtfo-overlay/rundown/internal/logic"
	"github.com/r35krag0th/zero-gtfo-overlay/rundown/internal/svc"
	"github.com/r35krag0th/zero-gtfo-overlay/rundown/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func listExpeditionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListRundownExpeditionsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewListExpeditionsLogic(r.Context(), svcCtx)
		resp, err := l.ListExpeditions(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
