package _interface

import (
	"context"
	"game-library-management-system/src/model"
)

type GameRepositorer interface {
	GetAllGames(ctx context.Context) ([]model.Game, error)
	GetGameById(ctx context.Context, id string) (*model.Game, error)
	AddGame(ctx context.Context, game model.Game) (*model.Game, error)
	UpdateAvailability(ctx context.Context, id string) (*model.Game, error)
	DeleteGame(ctx context.Context, id string) error
	FindGamesByDeveloper(ctx context.Context, developerName string) ([]model.Game, error)
	DeleteManyGamesByDeveloper(ctx context.Context, developerID string) error
}
