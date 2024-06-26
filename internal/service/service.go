package service

import (
	"context"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

type Storage interface {
	CheckDuplicateUser(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, email, passwordHash, passwordSalt string, isRecruiter bool) (int64, error)
	LoginUser(ctx context.Context, email string) (models.LoginUserResponse, error)
	CreateVacancy(ctx context.Context, vacancy models.VacancyRequest) error
	GetVacanciesByFilter(ctx context.Context, filter models.VacancyFilterRequest) ([]models.VacancyShortInfo, error)
	GetVacancyByID(ctx context.Context, vacancyID int64) (models.VacancyFullInfo, error)
	EditVacancy(ctx context.Context, vacancyID int64, vacancy models.VacancyRequest) error
	ArchiveVacancy(ctx context.Context, vacancyID int64) error
	VacancyApply(ctx context.Context, vacancyID, userID int64) error
	ResumeCreate(ctx context.Context, userID int64, resume models.Resume) error
	ResumeGet(ctx context.Context, userID int64) (models.Resume, error)
	ResumesGetByFilter(ctx context.Context, filter models.ResumeFilterRequest) ([]models.Resume, error)
	StagesGet(ctx context.Context, request models.GetUsersStagesRequest) ([]models.ResumeForStages, error)
}

// service provides business-logic
type service struct {
	storage Storage
}

// New creates new instance of actions
func New(storage Storage) *service {
	return &service{
		storage: storage,
	}
}
