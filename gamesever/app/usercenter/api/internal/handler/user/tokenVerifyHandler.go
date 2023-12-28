package user

import (
	"net/http"

	"go-game/app/usercenter/api/internal/logic/user"
	"go-game/app/usercenter/api/internal/svc"
	"go-game/common/result"
)

func TokenVerifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewTokenVerifyLogic(r.Context(), svcCtx)
		resp, err := l.TokenVerify()
		result.HttpResult(r, w, resp, err, "")
	}
}
