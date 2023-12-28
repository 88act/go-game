package basic

import (
	"net/http"

	"go-game/common/result"

	"go-game/app/basic/api/internal/logic/basic"
	"go-game/app/basic/api/internal/svc"
	"go-game/app/basic/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SendCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := basic.NewSendCodeLogic(r.Context(), svcCtx)
		resp, err := l.SendCode(&req)

		result.HttpResult(r, w, resp, err, req)
	}
}
