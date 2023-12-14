package svc

import (
	"context"
	"go-cms/app/game/cmd/api/internal/config"
	"go-cms/app/usercenter/cmd/rpc/usercenter"
	"go-cms/common"
	"go-cms/common/mycache"
	"go-cms/common/myconfig"
	"os"
	"time"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	//BasicRpc      basic.Basic
	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	logc.MustSetup(c.LogConf)
	logc.Info(context.Background(), c.Name, " 服务启动...", c.Host, " port=", c.Port)
	mycache.InitObj(c.Redis.Host, c.Redis.Pass, 0)
	myconfig.HttpRoot = c.LocalRes.BaseUrl
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		//BasicRpc:      basic.NewBasic(zrpc.MustNewClient(c.BasicRpcConf)),
	}

}

func GormMysql(mode string, dbConf common.DbConf) *gorm.DB {
	ormLogger := common.NewGormLogger(mode)
	dialector := mysql.New(mysql.Config{
		DSN:                       dbConf.DNS, // data source name
		DefaultStringSize:         256,        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,      // 根据当前 MySQL 版本自动配置
	})
	option := &gorm.Config{
		//禁用默认全局事务
		SkipDefaultTransaction: true,
		//开启预编译sql
		PrepareStmt: true,
		Logger:      ormLogger,
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix: "ucenter_", // 表名前缀，`User` 对应的表名是 `tb_users`
		//},
	}

	db, err := gorm.Open(dialector, option)
	if err != nil {
		logx.Error("game api MySQL启动异常", err.Error())
		os.Exit(0)
		return nil
	}
	sqlDb, err := db.DB()
	if err != nil {
		logx.Error("game api MySQL启动异常2", err.Error())
		os.Exit(0)
		return nil
	}
	sqlDb.SetMaxOpenConns(dbConf.MaxOpenConns)
	sqlDb.SetMaxIdleConns(dbConf.MaxIdleConns)
	sqlDb.SetConnMaxIdleTime(time.Second * time.Duration(dbConf.ConnMaxIdleTime))
	sqlDb.SetConnMaxLifetime(time.Second * time.Duration(dbConf.ConnMaxLifetime))
	//sCtx.Db = db
	logx.Infof("%+v", sqlDb.Stats())
	logx.Info("game api  MySQL启动 success")
	return db

}
