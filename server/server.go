package server

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/justhands13/learn_go/models"
)

type Server struct {
	*mux.Router
	shoppingCart []models.ShoppingItem
}

func NewServer() *Server {
	s := &Server{
		Router:       mux.NewRouter(),
		shoppingCart: []models.ShoppingItem{},
	}
	s.routes()
	return s
}

func (s *Server) addShoppingItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i models.ShoppingItem
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		i.ID = uuid.New()
		s.shoppingCart = append(s.shoppingCart, i)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) listShoppingCart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "applicatoin/json")
		if err := json.NewEncoder(w).Encode(s.shoppingCart); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (s *Server) removeShoppingItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr, _ := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		for i, item := range s.shoppingCart {
			if item.ID == id {
				s.shoppingCart = append(s.shoppingCart[:i], s.shoppingCart[i+1:]...)
				break
			}
		}
	}
}

func (s *Server) changeShoppingItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//parse newItem from request
		var newItem models.ShoppingItem
		if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//parse id from params
		idStr, _ := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		//change name in shopping cart
		for i := 0; i < len(s.shoppingCart); i++ {
			if s.shoppingCart[i].ID == id {
				s.shoppingCart[i].Name = newItem.Name
				break
			}
		}
	}
}
