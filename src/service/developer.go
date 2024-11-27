package service

import (
	"context"
	"game-library-management-system/src/interface"
	"game-library-management-system/src/model"
	"go.uber.org/zap"
)

type DeveloperService struct {
	developerRepository _interface.DeveloperRepositorer
	gameRepository      _interface.GameRepositorer
	logger              *zap.Logger
}

// NewDeveloperService creates a new DeveloperService
// It returns a pointer to a DeveloperService and an error
func NewDeveloperService(developerRepository _interface.DeveloperRepositorer, gameRepository _interface.GameRepositorer, logger *zap.Logger) (*DeveloperService, error) {
	return &DeveloperService{
		developerRepository: developerRepository,
		gameRepository:      gameRepository,
		logger:              logger,
	}, nil
}

// GetAllDevelopers gets all developers
func (s *DeveloperService) GetAllDevelopers(ctx context.Context) ([]model.Developer, error) {
	developers, err := s.developerRepository.GetAllDevelopers(ctx)
	if err != nil {
		s.logger.Error("Error getting all developers", zap.Error(err))
		return nil, err
	}
	return developers, nil
}

// GetDeveloperById gets a developer by ID
func (s *DeveloperService) GetDeveloperById(ctx context.Context, id string) (*model.Developer, error) {
	developer, err := s.developerRepository.GetDeveloperById(ctx, id)
	if err != nil {
		s.logger.Error("Error getting developer by ID", zap.String("id", id), zap.Error(err))
		return nil, err
	}
	return developer, nil
}

// AddDeveloper adds a developer
func (s *DeveloperService) AddDeveloper(ctx context.Context, developer model.Developer) (*model.Developer, error) {
	newDeveloper, err := s.developerRepository.AddDeveloper(ctx, developer)
	if err != nil {
		s.logger.Error("Error adding developer", zap.Error(err))
		return nil, err
	}
	return newDeveloper, nil
}

// UpdateDeveloper updates a developer
func (s *DeveloperService) UpdateDeveloper(ctx context.Context, id string, developer model.Developer) (*model.Developer, error) {
	updatedDeveloper, err := s.developerRepository.UpdateDeveloper(ctx, id, developer)
	if err != nil {
		s.logger.Error("Error updating developer", zap.String("id", id), zap.Error(err))
		return nil, err
	}
	return updatedDeveloper, nil
}

// DeleteDeveloper deletes a developer
func (s *DeveloperService) DeleteDeveloper(ctx context.Context, id string) error {
	err := s.gameRepository.DeleteManyGamesByDeveloper(ctx, id)
	if err != nil {
		s.logger.Error("Error deleting games by developer", zap.String("id", id), zap.Error(err))
		return err
	}

	err = s.developerRepository.DeleteDeveloper(ctx, id)
	if err != nil {
		s.logger.Error("Error deleting developer", zap.String("id", id), zap.Error(err))
		return err
	}

	return nil
}
