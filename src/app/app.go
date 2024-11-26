package app

import (
	"game-library-management-system/configs"
	"game-library-management-system/logger"
	"game-library-management-system/src/handler"
	"game-library-management-system/src/interface"
	"game-library-management-system/src/repository"
	"game-library-management-system/src/service"
	"go.uber.org/zap"
)

type App struct {
	config *configs.Config
	server *Server
	logger *zap.Logger
}

func NewApp() (*App, error) {
	config, err := configs.Load()
	if err != nil {
		return nil, err
	}

	lgr, err := logger.InitLogger()
	if err != nil {
		return nil, err
	}

	return &App{
		config: config,
		server: NewServer(config.Port),
		logger: lgr,
	}, nil
}

func (a *App) createDeveloperRepository() (_interface.DeveloperRepositorer, error) {
	developerInterface, err := repository.NewDeveloperRepository(a.config.DBUri, a.config.DBName)
	if err != nil {
		return nil, err
	}
	return developerInterface, nil
}

func (a *App) createGameRepository() (_interface.GameRepositorer, error) {
	gameInterface, err := repository.NewGameRepository(a.config.DBUri, a.config.DBName)
	if err != nil {
		return nil, err
	}
	return gameInterface, err
}

func (a *App) createDeveloperService(developerRepository _interface.DeveloperRepositorer, gameRepository _interface.GameRepositorer) (*service.DeveloperService, error) {
	developerService, err := service.NewDeveloperService(developerRepository, gameRepository, a.logger)
	if err != nil {
		return nil, err
	}
	return developerService, nil
}

func (a *App) createGameService(gameRepository _interface.GameRepositorer) (*service.GameService, error) {
	gameService, err := service.NewGameService(gameRepository, a.logger)
	if err != nil {
		return nil, err
	}
	return gameService, nil
}

func (a *App) setUpRoutes(handler *handler.Handler) {
	devsEndpoints := handler.RegisterRoutesForDevelopers()
	for _, endpoint := range devsEndpoints {
		a.server.Router.HandleFunc(endpoint.Path, endpoint.Handler).Methods(endpoint.Method)
	}

	gamesEndpoints := handler.RegisterRoutesForGames()
	for _, endpoint := range gamesEndpoints {
		a.server.Router.HandleFunc(endpoint.Path, endpoint.Handler).Methods(endpoint.Method)
	}
}

func (a *App) Run() error {
	developerRepository, err := a.createDeveloperRepository()
	if err != nil {
		return err
	}

	gameRepository, err := a.createGameRepository()
	if err != nil {
		return err
	}

	gameService, err := a.createGameService(gameRepository)
	if err != nil {
		return err
	}

	developerService, err := a.createDeveloperService(developerRepository, gameRepository)
	if err != nil {
		return err
	}

	a.setUpRoutes(handler.NewHandler(developerService, gameService))

	err = a.server.Start()
	if err != nil {
		return err
	}

	return nil
}
