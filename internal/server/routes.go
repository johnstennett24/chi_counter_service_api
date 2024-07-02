package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)

	r.Get("/health", s.healthHandler)

	r.Get("/ingredient/{id}", s.getIngredientByIdHandler)

	r.Get("/ingredients", s.getIngredientsHandler)

	r.Get("/stores/{id}", s.getStoreById)

	r.Get("/employees/{id}", s.getEmployeeById)

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}

func (s *Server) getIngredientsHandler(w http.ResponseWriter, r *http.Request) {
	ingredients, err := s.db.GetIngredients()
	if err != nil {
		log.Fatalf("error getting ingredients: %v", err)
	}
	jsonResp, _ := json.Marshal(ingredients)
	_, _ = w.Write(jsonResp)
}

func (s *Server) getIngredientByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ingredient, err := s.db.GetIngredientById(id)
	if err != nil {
		log.Fatalf("error getting ingredient by id: %v", err)
		jsonResp, err := json.Marshal(map[string]string{"error": "ingredient not found"})
		response, err := w.Write(jsonResp)
		w.WriteHeader(http.StatusInternalServerError)
		if err != nil {
			log.Fatalf("error writing response: %v", response)
		}
		_, _ = w.Write(jsonResp)
	}
	jsonResp, _ := json.Marshal(ingredient)
	fmt.Fprintf(log.Writer(), "Ingredient: %v\n", ingredient)
	_, _ = w.Write(jsonResp)
}

func (s *Server) getStoreById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	store, err := s.db.GetStoreById(id)
	if err != nil {
		log.Fatalf("error getting store by id: %v", err)
		jsonResp, err := json.Marshal(map[string]string{"error": "store not found"})
		response, err := w.Write(jsonResp)
		w.WriteHeader(http.StatusInternalServerError)
		if err != nil {
			log.Fatalf("error writing response: %v", response)
		}
		_, _ = w.Write(jsonResp)
	}
	jsonResp, _ := json.Marshal(store)
	fmt.Fprintf(log.Writer(), "Store: %v\n", store)
	_, _ = w.Write(jsonResp)
}

func (s *Server) getEmployeeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	employee, err := s.db.GetEmployeeById(id)
	if err != nil {
		log.Fatalf("error getting employee by id: %v", err)
		jsonResp, err := json.Marshal(map[string]string{"error": "employee not found"})
		response, err := w.Write(jsonResp)
		w.WriteHeader(http.StatusInternalServerError)
		if err != nil {
			log.Fatalf("error writing response: %v", response)
		}
		_, _ = w.Write(jsonResp)
	}
	jsonResp, _ := json.Marshal(employee)
	fmt.Fprintf(log.Writer(), "Employee: %v\n", employee)
	_, _ = w.Write(jsonResp)
}
