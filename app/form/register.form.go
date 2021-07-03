package form

type RegisterRequestBody struct {
	DisplayName string `json:"name" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
}
