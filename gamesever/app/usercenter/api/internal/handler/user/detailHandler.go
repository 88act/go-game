package user

import (
	"net/http"

	"go-game/app/usercenter/api/internal/logic/user"
	"go-game/app/usercenter/api/internal/svc"
	"go-game/common/ctxdata"
	"go-game/common/result"
)

func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewDetailLogic(r.Context(), svcCtx)
		userId := ctxdata.GetUidFromCtx(r.Context())
		resp, err := l.Detail(userId)
		result.HttpResult(r, w, resp, err, userId)
	}
}
