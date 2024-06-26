package models

import "time"

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

type AuthResponse struct {
	Token string `json:"token"`
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

type ResumeFilterRequest struct {
	SalaryFrom int64    `json:"salary_from,omitempty"`
	SalaryTo   int64    `json:"salary_to,omitempty"`
	Education  bool     `json:"education,omitempty"`
	Skills     []string `json:"skills,omitempty"`
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

type Resume struct {
	ID             int64        `json:"id,omitempty"`
	Fio            string       `json:"fio,omitempty"`
	Position       string       `json:"position,omitempty"`
	Gender         int64        `json:"gender,omitempty"`
	Address        string       `json:"address,omitempty"`
	BirthDate      time.Time    `json:"birth_date,omitempty"`
	Phone          string       `json:"phone,omitempty"`
	SalaryFrom     int64        `json:"salary_from,omitempty"`
	SalaryTo       int64        `json:"salary_to,omitempty"`
	Education      string       `json:"education,omitempty"`
	Skills         []string     `json:"skills,omitempty"`
	Nationality    string       `json:"nationality,omitempty"`
	Disabilities   bool         `json:"disabilities,omitempty"`
	WorkExperience []Experience `json:"work_experience,omitempty"`
}

type ResumeForStages struct {
	ID           int64     `json:"id,omitempty"`
	UserID       int64     `json:"user_id,omitempty"`
	Fio          string    `json:"fio,omitempty"`
	Position     string    `json:"position,omitempty"`
	Gender       int64     `json:"gender,omitempty"`
	Address      string    `json:"address,omitempty"`
	BirthDate    time.Time `json:"birth_date,omitempty"`
	Phone        string    `json:"phone,omitempty"`
	SalaryFrom   int64     `json:"salary_from,omitempty"`
	SalaryTo     int64     `json:"salary_to,omitempty"`
	Education    string    `json:"education,omitempty"`
	Skills       []string  `json:"skills,omitempty"`
	Nationality  string    `json:"nationality,omitempty"`
	Disabilities bool      `json:"disabilities,omitempty"`
	StageName    string    `json:"stage_name,omitempty"`
}

type Experience struct {
	CompanyName string    `json:"company_name,omitempty"`
	TimeFrom    time.Time `json:"time_from,omitempty"`
	TimeTo      time.Time `json:"time_to,omitempty"`
	Position    string    `json:"position,omitempty"`
	Description string    `json:"description,omitempty"`
}

type GetUsersStagesRequest struct {
	VacancyID int64  `json:"vacancy_id,omitempty"`
	StageName string `json:"stage_name,omitempty"`
}
