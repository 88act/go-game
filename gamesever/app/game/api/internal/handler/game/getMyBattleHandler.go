package game

import (
	"net/http"

	"go-game/app/game/api/internal/logic/game"
	"go-game/app/game/api/internal/svc"
	"go-game/common/result"
)

func GetMyBattleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := game.NewGetMyBattleLogic(r.Context(), svcCtx)
		resp, err := l.GetMyBattle()
		result.HttpResult(r, w, resp, err, "")
	}
}
