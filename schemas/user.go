package schemas

type User struct {
	Id          int     `json:"id" db:"id"`
	FirstName   *string `json:"first_name" db:"first_name"`
	Email       *string `json:"email" db:"email"`
	IsAdmin     bool    `json:"is_admin" db:"is_admin"`
	PhoneNumber *string `json:"phone_number" db:"phone_number"`
	YearOfBirth *string `json:"year_of_birth" db:"year_of_birth"`
	CreatedAt   *string `json:"created_at" db:"created_at"`
	Password    string  `json:"password" db:"password_hash"`
}

type UserInput struct {
	FirstName   *string `json:"first_name" db:"first_name"`
	Email       *string `json:"email" db:"email"`
	PhoneNumber *string `json:"phone_number" db:"phone_number"`
	YearOfBirth *string `json:"year_of_birth" db:"year_of_birth"`
	Password    *string `json:"password" db:"password_hash"`
}

type UserResponse struct {
	Id          *int    `json:"id" db:"id"`
	FirstName   *string `json:"first_name" db:"first_name"`
	Email       *string `json:"email" db:"email"`
	IsAdmin     *bool   `json:"is_admin" db:"is_admin"`
	PhoneNumber *string `json:"phone_number" db:"phone_number"`
	YearOfBirth *string `json:"year_of_birth" db:"year_of_birth"`
	CreatedAt   *string `json:"created_at" db:"created_at"`
}

// swagger -------------
type UserSignIn struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

type UserCreatedSwaggerResponse struct {
	Id int `json:"id"`
}

type UserTokenSwaggerResponse struct {
	Token string `json:"token"`
}
