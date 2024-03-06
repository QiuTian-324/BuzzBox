// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Mobile   string `json:"mobile"`   // 手机号
	Password string `json:"password"` // 密码
}

type LoginResponse struct {
	UserID  int64  `json:"user_id"` // 用户ID
	Token   Token  `json:"token"`   // 访问令牌
	Message string `json:"message"` // 登录成功响应消息
}

type RegisterRequest struct {
	Username         string `json:"username"`          // 用户名
	Mobile           string `json:"mobile"`            // 手机号
	Password         string `json:"password"`          // 密码
	VerificationCode string `json:"verification_code"` // 验证码
}

type RegisterResponse struct {
	UserID  int64  `json:"user_id"` // 用户ID
	Message string `json:"message"` // 注册成功响应消息
}

type Token struct {
	AccessToken  string `json:"access_token"`  // 访问令牌
	AccessExpire int64  `json:"access_expire"` // 访问令牌过期时间
}

type UserInfoResponse struct {
	UserID    int64  `json:"user_id"`    // 用户ID
	Username  string `json:"username"`   // 用户名
	Avatar    string `json:"avatar"`     // 头像链接
	Mobile    string `json:"mobile"`     // 手机号
	CreatedAt string `json:"created_at"` // 创建时间
	UpdatedAt string `json:"updated_at"` // 更新时间
}

type VerificationRequest struct {
	Mobile string `json:"mobile"` // 手机号
}

type VerificationResponse struct {
	Message string `json:"message"` // 验证结果
}
