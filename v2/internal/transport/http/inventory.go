package http

import (
	"encoding/json"
	"net/http"

	"github.com/yourusername/inventory-service/api/inventory/v1"
	"github.com/yourusername/inventory-service/internal/application"
	"github.com/gorilla/mux"
)

type InventoryHandler struct {
	inventorySvc *application.InventoryService
}

func NewInventoryHandler(inventorySvc *application.InventoryService) *InventoryHandler {
	return &InventoryHandler{
		inventorySvc: inventorySvc,
	}
}

func (h *InventoryHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/inventory", h.GetInventory).Methods("GET")
	router.HandleFunc("/inventory", h.UpdateInventory).Methods("PUT")
	router.HandleFunc("/inventories", h.ListInventories).Methods("GET")
}

func (h *InventoryHandler) GetInventory(w http.ResponseWriter, r *http.Request) {
	var req v1.GetInventoryRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	inventory, err := h.inventorySvc.GetInventory(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(inventory)
}

func (h *InventoryHandler) UpdateInventory(w http.ResponseWriter, r *http.Request) {
	var req v1.UpdateInventoryRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	inventory, err := h.inventorySvc.UpdateInventory(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(inventory)
}

func (h *InventoryHandler) ListInventories(w http.ResponseWriter, r *http.Request) {
	inventories, err := h.inventorySvc.ListInventories(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(inventories)
}
