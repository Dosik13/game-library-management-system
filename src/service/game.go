package service

import (
	"context"
	"game-library-management-system/src/interface"
	"game-library-management-system/src/model"
	"go.uber.org/zap"
)

type GameService struct {
	gameRepository _interface.GameRepositorer
	logger         *zap.Logger
}

// NewGameService creates a new GameService
// It returns a pointer to a GameService and an error
func NewGameService(gameRepository _interface.GameRepositorer, logger *zap.Logger) (_interface.GameServicer, error) {
	return &GameService{
		gameRepository: gameRepository,
		logger:         logger,
	}, nil
}

// GetAllGames gets all games
func (s *GameService) GetAllGames(ctx context.Context) ([]model.Game, error) {
	games, err := s.gameRepository.GetAllGames(ctx)
	if err != nil {
		s.logger.Error("Error getting all games", zap.Error(err))
		return nil, err
	}
	return games, nil
}

// GetGameById gets a game by ID
func (s *GameService) GetGameById(ctx context.Context, id string) (*model.Game, error) {
	game, err := s.gameRepository.GetGameById(ctx, id)
	if err != nil {
		s.logger.Error("Error getting game by ID", zap.String("id", id), zap.Error(err))
		return nil, err
	}
	return game, nil
}

// AddGame adds a game
func (s *GameService) AddGame(ctx context.Context, game model.Game) (*model.Game, error) {
	newGame, err := s.gameRepository.AddGame(ctx, game)
	if err != nil {
		s.logger.Error("Error adding game", zap.Error(err))
		return nil, err
	}
	return newGame, nil
}

// UpdateAvailability updates a game's availability
func (s *GameService) UpdateAvailability(ctx context.Context, id string) (*model.Game, error) {
	updatedGame, err := s.gameRepository.UpdateAvailability(ctx, id)
	if err != nil {
		s.logger.Error("Error updating game availability", zap.String("id", id), zap.Error(err))
		return nil, err
	}
	return updatedGame, nil
}

// DeleteGame deletes a game
func (s *GameService) DeleteGame(ctx context.Context, id string) error {
	err := s.gameRepository.DeleteGame(ctx, id)
	if err != nil {
		s.logger.Error("Error deleting game", zap.String("id", id), zap.Error(err))
		return err
	}
	return nil
}

// FindGamesByDeveloper finds games by developer
func (s *GameService) FindGamesByDeveloper(ctx context.Context, developer string) ([]model.Game, error) {
	games, err := s.gameRepository.FindGamesByDeveloper(ctx, developer)
	if err != nil {
		s.logger.Error("Error finding games by developer", zap.String("developer", developer), zap.Error(err))
		return nil, err
	}
	return games, nil
}

// DeleteManyGamesByDeveloper deletes many games by developer
func (s *GameService) DeleteManyGamesByDeveloper(ctx context.Context, developer string) error {
	err := s.gameRepository.DeleteManyGamesByDeveloper(ctx, developer)
	if err != nil {
		s.logger.Error("Error deleting games by developer", zap.String("developer", developer), zap.Error(err))
		return err
	}
	return nil
}
