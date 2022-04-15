package handler

import (
	"net/http"

	"github.com/r35krag0th/zero-gtfo-overlay/rundown/internal/logic"
	"github.com/r35krag0th/zero-gtfo-overlay/rundown/internal/svc"
	"github.com/r35krag0th/zero-gtfo-overlay/rundown/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func listRundownsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListRundownRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewListRundownsLogic(r.Context(), svcCtx)
		resp, err := l.ListRundowns(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
