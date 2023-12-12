package game

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-cms/app/game/cmd/api/internal/logic/game"
	"go-cms/app/game/cmd/api/internal/svc"
	"go-cms/app/game/cmd/api/internal/types"
	"go-cms/common/result"
)

func BuyGoodsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := game.NewBuyGoodsLogic(r.Context(), svcCtx)
		resp, err := l.BuyGoods(&req)
		result.HttpResult(r, w, resp, err, req)
	}
}
