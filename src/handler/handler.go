package handler

import (
	"encoding/json"
	_interface "game-library-management-system/src/interface"
	"game-library-management-system/src/model"
	"github.com/gorilla/mux"
	"net/http"
)

type Endpoint struct {
	Path    string
	Handler http.HandlerFunc
	Method  string
}

type Handler struct {
	developerService _interface.DeveloperRepositorer
	gameService      _interface.GameServicer
}

// NewHandler creates a new Handler instance.
func NewHandler(developerService _interface.DeveloperRepositorer, gameService _interface.GameServicer) *Handler {
	return &Handler{
		developerService: developerService,
		gameService:      gameService,
	}
}

// GetDevelopers handles the HTTP request to retrieve all developers.
func (h *Handler) GetDevelopers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	developers, err := h.developerService.GetAllDevelopers(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(developers)
	if err != nil {
		return
	}
}

// GetDeveloper handles the HTTP request to retrieve a developer by ID.
func (h *Handler) GetDeveloper(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id := vars["id"]

	developer, err := h.developerService.GetDeveloperById(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(developer)
	if err != nil {
		return
	}
}

// CreateDeveloper handles the HTTP request to create a new developer.
func (h *Handler) CreateDeveloper(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var developer model.Developer
	if err := json.NewDecoder(r.Body).Decode(&developer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := h.developerService.AddDeveloper(ctx, developer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// UpdateDeveloper handles the HTTP request to update an existing developer.
func (h *Handler) UpdateDeveloper(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id := vars["id"]

	var developer model.Developer
	if err := json.NewDecoder(r.Body).Decode(&developer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := h.developerService.UpdateDeveloper(ctx, id, developer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteDeveloper handles the HTTP request to delete a developer by ID.
func (h *Handler) DeleteDeveloper(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.developerService.DeleteDeveloper(ctx, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GetGames handles the HTTP request to retrieve all games.
func (h *Handler) GetGames(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	games, err := h.gameService.GetAllGames(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(games)
	if err != nil {
		return
	}
}

// GetGame handles the HTTP request to retrieve a game by ID.
func (h *Handler) GetGame(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id := vars["id"]

	game, err := h.gameService.GetGameById(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(game)
	if err != nil {
		return
	}
}

// CreateGame handles the HTTP request to create a new game.
func (h *Handler) CreateGame(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var game model.Game
	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := h.gameService.AddGame(ctx, game); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// UpdateGameAvailability handles the HTTP request to update a game's availability.
func (h *Handler) UpdateGameAvailability(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id := vars["id"]

	if _, err := h.gameService.UpdateAvailability(ctx, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteGame handles the HTTP request to delete a game by ID.
func (h *Handler) DeleteGame(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.gameService.DeleteGame(ctx, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// FindGamesByDeveloper handles the HTTP request to retrieve all games by a developer.
func (h *Handler) FindGamesByDeveloper(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id := vars["id"]

	games, err := h.gameService.FindGamesByDeveloper(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(games)
	if err != nil {
		return
	}
}

// RegisterRoutesForDevelopers registers the routes for developers.
func (h *Handler) RegisterRoutesForDevelopers() []Endpoint {
	return []Endpoint{
		{Path: "/developers", Handler: h.GetDevelopers, Method: "GET"},
		{Path: "/developers/{id}", Handler: h.GetDeveloper, Method: "GET"},
		{Path: "/developers", Handler: h.CreateDeveloper, Method: "POST"},
		{Path: "/developers/{id}", Handler: h.UpdateDeveloper, Method: "PUT"},
		{Path: "/developers/{id}", Handler: h.DeleteDeveloper, Method: "DELETE"},
	}
}

// RegisterRoutesForGames registers the routes for games.
func (h *Handler) RegisterRoutesForGames() []Endpoint {
	return []Endpoint{
		{Path: "/games", Handler: h.GetGames, Method: "GET"},
		{Path: "/games/{id}", Handler: h.GetGame, Method: "GET"},
		{Path: "/games", Handler: h.CreateGame, Method: "POST"},
		{Path: "/games/{id}", Handler: h.UpdateGameAvailability, Method: "PUT"},
		{Path: "/games/{id}", Handler: h.DeleteGame, Method: "DELETE"},
		{Path: "/games/developer/{developer}", Handler: h.FindGamesByDeveloper, Method: "GET"},
	}
}
