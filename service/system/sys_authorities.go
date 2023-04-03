package system

import (
	"gwa_server/global"
	"gwa_server/models/system"
)

type AuthoritiesService struct{}

// GetAuthoritiesList
// 获取角色列表
func (authoritieService AuthoritiesService) GetAuthoritiesList() (authorities []system.SysAuthority, err error) {
	var authoritiesList []system.SysAuthority
	err = global.GWA_DB.Find(&authoritiesList).Error
	if err != nil {
		return
	}
	return authoritiesList, err
}
