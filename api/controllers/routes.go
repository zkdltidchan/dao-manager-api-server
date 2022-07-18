package controllers

import "github.com/zkdltidchan/dao-manager-api-server/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//ManagerUsers routes
	// middlewares.SetMiddlewareAuthentication => for verify auth token
	s.Router.HandleFunc("/manager_users", middlewares.SetMiddlewareJSON(s.CreateManagerUser)).Methods("POST")
	s.Router.HandleFunc("/manager_users", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetManagerUsers))).Methods("GET")
	s.Router.HandleFunc("/manager_users/{id}", middlewares.SetMiddlewareJSON(s.GetManagerUser)).Methods("GET")
	s.Router.HandleFunc("/manager_users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateManagerUser))).Methods("PUT")
	s.Router.HandleFunc("/manager_users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteManagerUser)).Methods("DELETE")

	//Users routes
	// middlewares.SetMiddlewareAuthentication => for verify auth token
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetUsers))).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetUser))).Methods("GET")

	// //Posts routes
	// s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	// s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")
}
