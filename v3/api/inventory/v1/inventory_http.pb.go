package v1

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http/binding"
	"github.com/gorilla/mux"

	"github.com/yourusername/inventory-service/internal/service"
)

func RegisterInventoryHTTPServer(srv *http.Server, invSvc service.InventoryService) {
	router := mux.NewRouter()
	router.Use(middleware.Logging())
	router.Use(middleware.Recovery())

	h := &inventoryHTTPHandler{
		invSvc: invSvc,
	}

	router.HandleFunc("/v1/inventory/{sku}", h.GetInventory).Methods(http.MethodGet)
	router.HandleFunc("/v1/inventory/{sku}", h.UpdateInventory).Methods(http.MethodPut)

	srv.Handler = router
}

type inventoryHTTPHandler struct {
	invSvc service.InventoryService
}

func (h *inventoryHTTPHandler) GetInventory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sku := vars["sku"]

	ctx := r.Context()
	resp, err := h.invSvc.GetInventory(ctx, &GetInventoryRequest{Sku: sku})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (h *inventoryHTTPHandler) UpdateInventory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sku := vars["sku"]

	var req UpdateInventoryRequest
	if err := binding.Bind(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	req.Sku = sku

	ctx := r.Context()
	resp, err := h.invSvc.UpdateInventory(ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func init() {
	encoding.RegisterCodec(jsonCodec{})
}

type jsonCodec struct{}

func (jsonCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (jsonCodec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (jsonCodec) Name() string {
	return "json"
}
