package game

import (
	"net/http"

	"go-game/app/game/api/internal/logic/game"
	"go-game/app/game/api/internal/svc"
	"go-game/common/result"
)

func GetNotifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := game.NewGetNotifyLogic(r.Context(), svcCtx)
		resp, err := l.GetNotify()
		result.HttpResult(r, w, resp, err, "")
	}
}
