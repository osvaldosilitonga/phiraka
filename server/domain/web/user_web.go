package web

type RegisterReq struct {
	Username string `json:"username" validate:"required,min=4,max=128"`
	Password string `json:"password" validate:"required,min=5,max=8"`
}

type LoginReq struct {
	Username string `json:"username" validate:"required,min=4,max=128"`
	Password string `json:"password" validate:"required,min=5,max=8"`
}

type DeleteReq struct {
	Username string `json:"username" validate:"required,min=4,max=128"`
}

type UpdateReq struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// ----------------------------------------------------------------------------

type FindAllUserResp struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	CreateTime string `json:"create_time"`
}
