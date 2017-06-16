package hello

// User from request
type User struct {
	ID    int    `json:"id" form:"id" query:"id"`
	Name  string `json:"name" form:"name" query:"name" validate:"required"`
	Email string `json:"email" form:"email" query:"email" validate:"email"`
}
