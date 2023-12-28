package user

import (
	"net/http"

	"go-game/common/result"

	"go-game/app/usercenter/api/internal/logic/user"
	"go-game/app/usercenter/api/internal/svc"
	"go-game/app/usercenter/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		result.HttpResult(r, w, resp, err, req)
	}
}
