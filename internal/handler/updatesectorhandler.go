package handler

import (
	"net/http"

	"github.com/r35krag0th/zero-gtfo-overlay/internal/logic"
	"github.com/r35krag0th/zero-gtfo-overlay/internal/svc"
	"github.com/r35krag0th/zero-gtfo-overlay/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateSectorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSectorRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateSectorLogic(r.Context(), svcCtx)
		resp, err := l.UpdateSector(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
