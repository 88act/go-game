package user

import (
	"net/http"

	"go-game/app/usercenter/api/internal/logic/user"
	"go-game/app/usercenter/api/internal/svc"
	"go-game/app/usercenter/api/internal/types"
	"go-game/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginScanHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginScanReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewLoginScanLogic(r.Context(), svcCtx)
		resp, err := l.LoginScan(&req)
		result.HttpResult(r, w, resp, err, req)
	}
}
