package basic

import (
	"net/http"

	"go-game/common/result"

	"go-game/app/basic/api/internal/logic/basic"
	"go-game/app/basic/api/internal/svc"
	"go-game/app/basic/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CaptchaReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		l := basic.NewCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.Captcha(&req)
		result.HttpResult(r, w, resp, err, req)
	}
}
