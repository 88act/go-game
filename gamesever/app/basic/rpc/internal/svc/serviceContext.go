package svc

import (
	"context"
	"go-game/app/basic/rpc/internal/config"
	"go-game/common/myconfig"
	"go-game/common/orm"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	DB     *orm.DB
	Redis  *redis.Redis
	//BasicFileSev *baseModel.BasicFileSev
	//BasicEmailSev *model.BasicEmailSev
	//BasicSmsSev   *model.BasicSmsSev
}

func NewServiceContext(c config.Config) *ServiceContext {
	logc.MustSetup(c.LogConf)
	logc.Info(context.Background(), c.Name, " RPC 服务器启动...")
	myconfig.HttpRoot = c.LocalRes.BaseUrl
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DNS,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.ConnMaxLifetime,
	})
	rds := redis.MustNewRedis(redis.RedisConf{
		Host: c.Redis.Host,
		Pass: c.Redis.Pass,
		Type: c.Redis.Type,
	})

	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  rds,
		//BasicFileSev: baseModel.NewBasicFileSev(db, rds),
	}

}
