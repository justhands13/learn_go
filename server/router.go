package server

func (s *Server) routes() {
	s.HandleFunc("/shopping-list", s.listShoppingCart()).Methods("GET")
	s.HandleFunc("/shopping-list", s.addShoppingItem()).Methods("POST")
	s.HandleFunc("/shopping-list/{id}", s.removeShoppingItem()).Methods("DELETE")
	s.HandleFunc("/shopping-list/{id}", s.changeShoppingItem()).Methods("PATCH")
}
