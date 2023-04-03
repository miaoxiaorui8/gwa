package request

// User login structure
type Login struct {
	Username string `form:"username"` // 用户名
	Password string `form:"password"` // 密码
}
