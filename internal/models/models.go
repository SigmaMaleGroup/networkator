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

type LoginUserResponse struct {
	UserID       int64
	IsRecruiter  bool
	PasswordHash string
	PasswordSalt string
}

type VacancyRequest struct {
	RecruiterID    int64
	Name           string   `json:"name"`
	Experience     int64    `json:"experience,omitempty"`
	City           string   `json:"city,omitempty"`
	EmploymentType int64    `json:"employment_type,omitempty"`
	SalaryFrom     int64    `json:"salary_from,omitempty"`
	SalaryTo       int64    `json:"salary_to,omitempty"`
	CompanyName    string   `json:"company_name,omitempty"`
	Skills         []string `json:"skills,omitempty"`
	Address        string   `json:"address,omitempty"`
	Description    string   `json:"description,omitempty"`
}

type VacancyFilterRequest struct {
	Experience     int64  `json:"experience,omitempty"`
	City           string `json:"city,omitempty"`
	EmploymentType int64  `json:"employment_type,omitempty"`
	SalaryFrom     int64  `json:"salary_from,omitempty"`
	SalaryTo       int64  `json:"salary_to,omitempty"`
	CompanyName    string `json:"company_name,omitempty"`
	Archived       bool   `json:"archived,omitempty"`
}

type VacancyFilterResponse struct {
	Vacancies []VacancyShortInfo `json:"vacancies"`
}

type VacancyShortInfo struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	SalaryFrom     int64  `json:"salary_from,omitempty"`
	SalaryTo       int64  `json:"salary_to,omitempty"`
	City           string `json:"city,omitempty"`
	EmploymentType int64  `json:"employment_type,omitempty"`
	Description    string `json:"description,omitempty"`
}

type VacancyFullInfo struct {
	ID             int64    `json:"id"`
	Name           string   `json:"name"`
	City           string   `json:"city,omitempty"`
	SalaryFrom     int64    `json:"salary_from,omitempty"`
	SalaryTo       int64    `json:"salary_to,omitempty"`
	Skills         []string `json:"skills,omitempty"`
	Experience     int64    `json:"experience,omitempty"`
	Address        string   `json:"address,omitempty"`
	Description    string   `json:"description,omitempty"`
	EmploymentType int64    `json:"employment_type,omitempty"`
}
