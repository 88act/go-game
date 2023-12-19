package listen

import (
	"context"
	"go-game/app/basic/rpc/internal/config"
	kqMq "go-game/app/basic/rpc/internal/mqs/kq"
	"go-game/app/basic/rpc/internal/svc"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

// 发布/订阅 pub sub use kq (kafka)
func KqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.ImageZipConf, kqMq.NewImageZipMq(ctx, svcContext)),
		kq.MustNewQueue(c.SendEmailConf, kqMq.NewSendEmailMq(ctx, svcContext)),
		kq.MustNewQueue(c.SendSmsConf, kqMq.NewSendSmsMq(ctx, svcContext)),
		//.....
	}

}
