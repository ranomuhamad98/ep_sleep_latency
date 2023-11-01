package controllers
import "asia-southeast2-bei-eipo-trial-apigee.cloudfunctions.net/ep_sleep_latency/api/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
}
