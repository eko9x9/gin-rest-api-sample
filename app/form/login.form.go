package form

type LoginRequestBody struct {
	Username string `josn:"usernmae" binding:"required" validate:"required"`
	Password string `josn:"usernmae" binding:"required" validate:"required"`
}
