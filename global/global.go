package global

import (
	"github.com/sirupsen/logrus"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"gwa_server/config"
)

var (
	GWA_Config              *config.Config
	GWA_DB                  *gorm.DB
	GWA_Log                 *logrus.Logger
	GWA_Concurrency_Control = &singleflight.Group{} //并发控制

	BlackCache local_cache.Cache
)
