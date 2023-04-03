package system

import (
	"gwa_server/global"
	"gwa_server/models/system"
	"strconv"
)

type MenuService struct{}

// GetMenuTreeMap
// 根据用户权限id获取菜单列表
func (menuService *MenuService) GetMenuTreeMap(authorityId uint) (treeMap map[string][]system.SysMenu, err error) {
	var allMenus []system.SysMenu
	var baseMenu []system.SysBaseMenu
	treeMap = make(map[string][]system.SysMenu)

	var SysAuthorityMenus []system.SysAuthorityMenu
	err = global.GWA_DB.Where("sys_authority_authority_id = ?", authorityId).Find(&SysAuthorityMenus).Error
	if err != nil {
		return
	}

	var MenuIds []string

	for i := range SysAuthorityMenus {
		MenuIds = append(MenuIds, SysAuthorityMenus[i].MenuId)
	}

	err = global.GWA_DB.Where("id in (?)", MenuIds).Order("sort").Find(&baseMenu).Error
	if err != nil {
		return
	}
	for i := range baseMenu {
		allMenus = append(allMenus, system.SysMenu{
			SysBaseMenu: baseMenu[i],
			AuthorityId: authorityId,
			MenuId:      strconv.Itoa(int(baseMenu[i].ID)),
		})
	}
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

// GetMenu
// 获取所有菜单
func (menuService *MenuService) GetMenu(authorityId uint) (treeMap map[string][]system.SysMenu, err error) {
	var allMenus []system.SysMenu
	var baseMenu []system.SysBaseMenu
	treeMap = make(map[string][]system.SysMenu)

	//var SysAuthorityMenus []system.SysAuthorityMenu
	//err = global.GWA_DB.Where("sys_authority_authority_id = ?", authorityId).Find(&SysAuthorityMenus).Error
	if err != nil {
		return
	}

	//var MenuIds []string
	//
	//for i := range SysAuthorityMenus {
	//	MenuIds = append(MenuIds, SysAuthorityMenus[i].MenuId)
	//}

	err = global.GWA_DB.Order("sort").Find(&baseMenu).Error
	if err != nil {
		return
	}
	for i := range baseMenu {
		allMenus = append(allMenus, system.SysMenu{
			SysBaseMenu: baseMenu[i],
			AuthorityId: authorityId,
			MenuId:      strconv.Itoa(int(baseMenu[i].ID)),
		})
	}
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}
