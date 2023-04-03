package response

import "gwa_server/models/system"

type SysUserResponse struct {
	User system.SysAdmin `json:"admin"`
}

type LoginResponse struct {
	Admin     system.SysAdmin `json:"admin"`
	Token     string          `json:"token"`
	ExpiresAt int64           `json:"expiresAt"`
}
