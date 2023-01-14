package fiber

func (s *Server) SetupBookserviceRoutes() {
	s.App.Get("/api/book", s.GetBook)
	s.App.Get("/api/books", s.GetBooks)
	s.App.Post("/api/book", s.CreateBook)
	s.App.Patch("/api/book", s.UpdateBook)
	s.App.Delete("/api/book", s.DeleteBook)
}
