package system

import (
	"gwa_server/global"
	"gwa_server/models/system"
)

type ApiService struct{}

// GetApiList
// 获取API列表
func (apiService ApiService) GetApiList() (api []system.SysApi, err error) {
	var apiList []system.SysApi
	err = global.GWA_DB.Find(&apiList).Error
	global.GWA_Log.Warn(apiList)
	if err != nil {
		return
	}
	return apiList, err
}
