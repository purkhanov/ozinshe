package schemas

type User struct {
	Id          int     `json:"id" db:"id"`
	FirstName   *string `json:"first_name" db:"first_name"`
	Email       string  `json:"email" db:"email" binding:"required"`
	IsAdmin     bool    `json:"is_admin" db:"is_admin"`
	PhoneNumber *string `json:"phone_number" db:"phone_number"`
	YearOfBirth *string `json:"year_of_birth" db:"year_of_birth"`
	CreatedAt   string  `json:"created_at" db:"created_at"`
	Password    string  `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	FirstName   *string `json:"first_name"`
	Email       *string `json:"email"`
	PhoneNumber *string `json:"phone_number"`
	YearOfBirth *string `json:"year_of_birth"`
	Password    *string `json:"password"`
}

func (u *UpdateUserInput) ToMap() map[string]any {
	userMap := make(map[string]any, 5)

	if u.FirstName != nil {
		userMap["first_name"] = *u.FirstName
	}

	if u.Email != nil {
		userMap["email"] = *u.Email
	}

	if u.PhoneNumber != nil {
		userMap["phone_number"] = *u.PhoneNumber
	}

	if u.YearOfBirth != nil {
		userMap["year_of_birth"] = *u.YearOfBirth
	}

	if u.Password != nil {
		userMap["password_hash"] = *u.Password
	}

	return userMap
}
