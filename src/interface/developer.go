package _interface

import (
	"context"
	"game-library-management-system/src/model"
)

type DeveloperRepositorer interface {
	GetAllDevelopers(ctx context.Context) ([]model.Developer, error)
	GetDeveloperById(ctx context.Context, id string) (*model.Developer, error)
	AddDeveloper(ctx context.Context, developer model.Developer) (*model.Developer, error)
	UpdateDeveloper(ctx context.Context, id string, developer model.Developer) (*model.Developer, error)
	DeleteDeveloper(ctx context.Context, id string) error
}

type DeveloperServicer interface {
	GetAllDevelopers(ctx context.Context) ([]model.Developer, error)
	GetDeveloperById(ctx context.Context, id string) (*model.Developer, error)
	AddDeveloper(ctx context.Context, developer model.Developer) (*model.Developer, error)
	UpdateDeveloper(ctx context.Context, id string, developer model.Developer) (*model.Developer, error)
	DeleteDeveloper(ctx context.Context, id string) error
}
