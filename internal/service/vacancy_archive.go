package service

import "context"

func (s *service) ArchiveVacancy(ctx context.Context, vacancyID int64) error {
	return s.storage.ArchiveVacancy(ctx, vacancyID)
}
