package game

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-cms/app/game/cmd/api/internal/logic/game"
	"go-cms/app/game/cmd/api/internal/svc"
	"go-cms/app/game/cmd/api/internal/types"
	"go-cms/common/result"
)

func ChangeScoreHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChangeScoreReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := game.NewChangeScoreLogic(r.Context(), svcCtx)
		resp, err := l.ChangeScore(&req)
		result.HttpResult(r, w, resp, err, req)
	}
}
