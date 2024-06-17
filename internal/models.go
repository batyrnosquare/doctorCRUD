package internal

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type Doctor struct {
	Name       string `json:"name" validate:"required" `
	Surname    string `json:"surname" validate:"required" `
	Position   string `json:"position" validate:"required" `
	Age        uint8  `json:"age" validate:"required" `
	Experience uint8  `json:"experience" validate:"required" `
}
