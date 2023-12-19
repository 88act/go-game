package game

import (
	"net/http"

	"go-game/app/game/api/internal/logic/game"
	"go-game/app/game/api/internal/svc"
	"go-game/app/game/api/internal/types"
	"go-game/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UseGoodsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := game.NewUseGoodsLogic(r.Context(), svcCtx)
		resp, err := l.UseGoods(&req)
		result.HttpResult(r, w, resp, err, req)
	}
}
