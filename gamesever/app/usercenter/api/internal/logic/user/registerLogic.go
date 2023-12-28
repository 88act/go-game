package user

import (
	"context"

	"go-game/app/usercenter/api/internal/svc"
	"go-game/app/usercenter/api/internal/types"
	"go-game/app/usercenter/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.LoginReq) (resp *types.LoginResp, err error) {

	rpcReq := pb.RegisterReq{Username: req.Username, Password: req.Password, LoginType: 1}
	if rpcResp, err := l.svcCtx.UsercenterRpc.Register(l.ctx, &rpcReq); err == nil {
		var myResp types.LoginResp
		_ = copier.Copy(&myResp, rpcResp)
		return &myResp, nil
	} else {
		return nil, err
	}
}

// func GenerateToken(accessExpire int64, accessSecret string, user model.MemUser) (*types.LoginResp, error) {
// 	now := time.Now().Unix()

// 	accessToken, err := GetJwtToken(accessSecret, now, accessExpire, user.Id, user.UserType, user.CuId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &types.LoginResp{
// 		AccessToken:  accessToken,
// 		AccessExpire: now + accessExpire,
// 		RefreshAfter: now + accessExpire/2,
// 	}, nil
// }

// func GetJwtToken(secretKey string, iat, seconds, userId int64, userType int, cuId int64) (string, error) {

// 	claims := make(jwt.MapClaims)
// 	claims["exp"] = iat + seconds
// 	claims["iat"] = iat
// 	claims[ctxdata.CtxKeyJwtUserId] = userId
// 	claims[ctxdata.CtxKeyJwtUserType] = userType
// 	claims[ctxdata.CtxKeyCuId] = cuId
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	token.Claims = claims
// 	return token.SignedString([]byte(secretKey))
// }
