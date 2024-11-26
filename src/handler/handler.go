package handler

import (
	"encoding/json"
	"game-library-management-system/src/model"
	"game-library-management-system/src/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Endpoint struct {
	Path    string
	Handler http.HandlerFunc
	Method  string
}

type Handler struct {
	developerService *service.DeveloperService
	gameService      *service.GameService
}

func NewHandler(developerService *service.DeveloperService, gameService *service.GameService) *Handler {
	return &Handler{
		developerService: developerService,
		gameService:      gameService,
	}
}

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

func (h *Handler) FindGamesByDeveloper(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id := vars["id"]

	games, err := h.gameService.FindGameByDeveloper(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(games)
	if err != nil {
		return
	}
}

func (h *Handler) RegisterRoutesForDevelopers() []Endpoint {
	return []Endpoint{
		{Path: "/developers", Handler: h.GetDevelopers, Method: "GET"},
		{Path: "/developers/{id}", Handler: h.GetDeveloper, Method: "GET"},
		{Path: "/developers", Handler: h.CreateDeveloper, Method: "POST"},
		{Path: "/developers/{id}", Handler: h.UpdateDeveloper, Method: "PUT"},
		{Path: "/developers/{id}", Handler: h.DeleteDeveloper, Method: "DELETE"},
	}
}

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
