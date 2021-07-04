package form

type RegisterRequestBody struct {
	DisplayName string `json:"name" binding:"required" validate:"required"`
	Username    string `json:"username" binding:"required" validate:"required"`
	Password    string `json:"password" binding:"required" validate:"required"`
}
