package system

import "gwa_server/global"

type JwtBlacklist struct {
	global.GWA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
