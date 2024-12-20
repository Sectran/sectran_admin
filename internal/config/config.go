package config

import (
	"github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/plugins/casbin"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth               rest.AuthConf
	CROSConf           config.CROSConf
	UploadConf         UploadConf
	CasbinDatabaseConf config.DatabaseConf
	RedisConf          config.RedisConf
	CasbinConf         casbin.CasbinConf
	DatabaseConf       config.DatabaseConf
	I18nConf           i18n.Conf
}

type UploadConf struct {
	MaxImageSize int64 `json:",default=33554432,env=MAX_IMAGE_SIZE"`
	MaxVideoSize int64 `json:",default=1073741824,env=MAX_VIDEO_SIZE"`
	MaxAudioSize int64 `json:",default=33554432,env=MAX_AUDIO_SIZE"`
	MaxOtherSize int64 `json:",default=10485760,env=MAX_OTHER_SIZE"`
}
