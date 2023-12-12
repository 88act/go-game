package game

import (
	"net/http"

	"go-cms/app/game/cmd/api/internal/logic/game"
	"go-cms/app/game/cmd/api/internal/svc"
	"go-cms/common/result"
)

func GetRankHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := game.NewGetRankLogic(r.Context(), svcCtx)
		resp, err := l.GetRank()
		result.HttpResult(r, w, resp, err, "")
	}
}
