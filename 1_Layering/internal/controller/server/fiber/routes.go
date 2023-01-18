package fiber

func (s *Server) SetupBookserviceRoutes() {
	api := s.App.Group("/api")
	api.Get("/books/:id", s.GetBook)
	api.Get("/books", s.GetBooks)
	api.Post("/books", s.CreateBook)
	api.Patch("/books", s.UpdateBook)
	api.Delete("/books", s.DeleteBook)
}
