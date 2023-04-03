package system

import (
	uuid "github.com/satori/go.uuid"
	"gwa_server/global"
)

type SysAdmin struct {
	global.GWA_MODEL
	UUID        uuid.UUID      `json:"uuid" gorm:"index;comment:用户UUID"`                                                                                               // 用户UUID
	Username    string         `json:"userName" gorm:"index;comment:用户登录名"`                                                                                            // 用户登录名
	Password    string         `json:"-"  gorm:"comment:用户登录密码"`                                                                                                       // 用户登录密码
	NickName    string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                                                                      // 用户昵称
	HeaderImg   string         `json:"headerImg" gorm:"default:https://i0.hdslb.com/bfs/face/6775818308a81520555a398d10eba857f061af80.jpg@150w_150h.jpg;comment:用户头像"` // 用户头像
	AuthorityId uint           `json:"authorityId" gorm:"default:1;comment:用户角色ID"`                                                                                    // 用户角色ID
	Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;comment:用户角色"`
	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Enable      int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
}

func (SysAdmin) TableName() string {
	return "sys_admin"
}
