// MemUser
package baseModel

import (
	"context"

	"go-game/common/myconfig"
	"go-game/common/utils"
	"strings"

	"github.com/gogf/gf/util/gconv"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type BaseModelSev struct {
	Db  *gorm.DB
	Rds *redis.Redis
	//Cache           cache.ClusterConf
	CacheKeyListPre string
	CacheKeyPre     string
}

// func (m *BaseModelSev) GetPathByGuid(ctx context.Context, guid string) (path string, err error) {

// 	//basicFS := new(businessSev.BasicFileService)
// 	// 先从 redis 获取 ，没有则读取数据库 ，缓存2小时
// 	if utils.IsEmptyStr(guid) {
// 		return "", nil
// 	}
// 	if len(guid) < 32 {
// 		logx.Errorf(" GetPathByGuid无效guid=" + guid)
// 		return "", nil
// 	}
// 	///  byteData := []byte("ssss112112qweqweqwe2")
// 	//	GetCache().Set("ssss", byteData)
// 	cacheKey := m.GetCacheKey("file", guid)
// 	cacheData, err := cache.GetCache().Get(cacheKey)
// 	if err == nil {
// 		//global.LOG.Info("from cacheKey=" + cacheKey)
// 		path = cacheData.(string)
// 		//global.LOG.Info("from path=" + path)
// 		return path, err
// 	}
// 	basicFile := BasicFile{}
// 	err =  m.Db.WithContext(ctx).Model(basicFile).Where("guid = ?", guid).First(&basicFile).Error
// 	if err != nil {
// 		//cacheData = []byte("")
// 		cache.GetCache().Set(cacheKey, "")
// 		return "", err
// 	}
// 	path = basicFile.Path
// 	if !utils.IsEmpty(path) {
// 		path = ctx.Config.LocalRes.BaseUrl + path
// 	}
// 	//cacheData = []byte(path)
// 	cache.GetCache().Set(cacheKey, path)
// 	return path, err
// }

func (m *BaseModelSev) GetPathByGuid(ctx context.Context, guid string) (path string, err error) {
	// 如果是 http开头 ，直接返回
	if strings.HasPrefix(guid, "http") {
		return guid, nil
	}

	//basicFS := new(businessSev.BasicFileService)
	// 先从 redis 获取 ，没有则读取数据库 ，缓存2小时
	if utils.IsEmptyStr(guid) {
		return "", nil
	}
	// if len(guid) < 32 {
	// 	logx.WithContext(ctx).Infow("GetPathByGuid无效guid=" + guid)
	// 	return "", nil
	// }
	///  byteData := []byte("ssss112112qweqweqwe2")
	//	GetCache().Set("ssss", byteData)
	cacheKey := m.GetCacheKey("file", guid)
	path, err = m.Rds.Get(cacheKey)
	if err == nil {
		return path, err
	}
	basicFile := BasicFile{}
	// mapData := make(map[string]interface{})
	// mapData["guid"] = guid
	// mapData["status"] = 1
	err = m.Db.WithContext(ctx).Where("guid = ?  ", guid).Select("path,guid").First(&basicFile).Error
	if err != nil {
		//cacheData = []byte("")
		m.Rds.Set(cacheKey, "")
		return "", err
	}
	path = basicFile.Path
	if !utils.IsEmpty(path) {
		path = myconfig.HttpRoot + path
	}
	//cacheData = []byte(path)
	m.Rds.Set(cacheKey, path)
	return path, err
}

func (m *BaseModelSev) DelCache(table string, ids []int64) {
	for _, v := range ids {
		cacheKey := m.GetCacheKey(table, v)
		m.Rds.Del(cacheKey)
	}
}
func (m *BaseModelSev) DelCache32(table string, ids []int) {
	for _, v := range ids {
		cacheKey := m.GetCacheKey(table, v)
		m.Rds.Del(cacheKey)
	}
}
func (m *BaseModelSev) GetCacheKey(table string, id interface{}) (key string) {
	key = table + "_" + gconv.String(id)
	return key
}
