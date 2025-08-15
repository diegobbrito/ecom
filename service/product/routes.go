package product

import (
	"net/http"

	"github.com/diegobbrito/ecom/types"
	"github.com/diegobbrito/ecom/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handlerGetProduct).Methods(http.MethodGet)
	router.HandleFunc("/products", h.handlerCreateProduct).Methods(http.MethodPost)
}

func (h *Handler) handlerGetProduct(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusOK, ps)
}

func (h *Handler) handlerCreateProduct(w http.ResponseWriter, r *http.Request) {
	var p types.Product

	if err := utils.ParseJSON(r, &p); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.store.CreateProduct(p); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, p)
}
