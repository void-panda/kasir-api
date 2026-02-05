package handler

import (
	"encoding/json"
	"net/http"

	"kasir-api/model"
	"kasir-api/service"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(service *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	default:
		model.Error(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

// Checkout godoc
// @Summary Checkout products
// @Description Membuat transaksi baru dan mengurangi stok produk
// @Tags transactions
// @Accept json
// @Produce json
// @Param request body model.CheckoutRequest true "Checkout Request"
// @Success 200 {object} model.Response
// @Router /api/checkout [post]
func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var req model.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		model.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	transaction, err := h.service.Checkout(req.Items)
	if err != nil {
		model.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	model.Success(w, http.StatusOK, "checkout success", transaction)
}

// GetTodaySummary godoc
// @Summary Get today's sales summary
// @Description Mengambil ringkasan penjualan hari ini
// @Tags reports
// @Accept json
// @Produce json
// @Success 200 {object} model.Response
// @Router /api/report/hari-ini [get]
func (h *TransactionHandler) GetTodaySummary(w http.ResponseWriter, r *http.Request) {
	summary, err := h.service.GetTodaySummary()
	if err != nil {
		model.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	model.Success(w, http.StatusOK, "successfully get summary", summary)
}
