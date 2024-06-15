package service

import "context"

func (s *service) VacancyApply(ctx context.Context, vacancyID, userID int64) error {
	return s.storage.VacancyApply(ctx, vacancyID, userID)
}
