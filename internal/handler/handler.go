package handler

import (
   "encoding/json"
   "fmt"
   "net/http"
   "errors"
   "l0/internal/service"
   "l0/internal/model"
   "github.com/gorilla/mux"
//    "strconv"
//    "github.com/sirupsen/logrus"
)

type Handler struct{
  service *service.OrderService
}

func NewOrderHandler(orderService *service.OrderService) *Handler{
	resultHandler := new(Handler)
	resultHandler.service = *&orderService
	return resultHandler
}

func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r) 
	if vars["order_uid"] == "" {
		WrapError(w, errors.New("missing id"))
		return
	}


	order, err := h.service.GetOrder(vars["order_uid"])
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{} {
		"result" : "OK",
		"data" : order,
	}

	WrapOK(w, m)
}

func (h *Handler) NewOrder(w http.ResponseWriter, r *http.Request){
	var newOrder model.Order
	err := json.NewDecoder(r.Body).Decode(&newOrder)

	if err != nil {
		WrapError(w, err)
		return
	}
	 
	err = h.service.CreateOrder(newOrder)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{} {
		"result" : "OK",
		"data" : "",
	}

	WrapOK(w, m)
}


func WrapError(w http.ResponseWriter, err error) {
	WrapErrorWithStatus(w, err, http.StatusBadRequest)
}

func WrapErrorWithStatus(w http.ResponseWriter, err error, httpStatus int) {
	var m = map[string]string {
		"result" : "error",
		"data" : err.Error(),
	}

	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")  
	w.WriteHeader(httpStatus) 
	fmt.Fprintln(w, string(res))
}

func WrapOK(w http.ResponseWriter, m map[string]interface{}) {
	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(res))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	WrapErrorWithStatus(w, errors.New("not found"), http.StatusNotFound)
}