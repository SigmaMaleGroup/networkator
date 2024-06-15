package models

// ContextKey is type for context keys
type ContextKey string

const (
	CtxUserIDKey ContextKey = "userID"
	CtxEmailKey  ContextKey = "email"
	CtxRoleKey   ContextKey = "role"
)

const (
	Recruiter = "RECRUITER"
	Applicant = "APPLICANT"
)

func GetUserType(isRecruiter bool) string {
	if isRecruiter {
		return Recruiter
	}

	return Applicant
}

// RegisterCredentials provides users register credentials request
type RegisterCredentials struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsRecruiter bool   `json:"is_recruiter"`
}

// LoginCredentials provides users login credentials request
type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginUserResponse provides response for user login request
type LoginUserResponse struct {
	UserID       int64
	IsRecruiter  bool
	PasswordHash string
	PasswordSalt string
}
