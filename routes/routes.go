package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

func Routes(conn *pgx.Conn) *chi.Mux {

	routes := []Route{
		{
			Path:    "/tasks",
			Method:  "GET",
			Handler: GetAllTasksHandler(conn),
		},
		{
			Path:    "/task/{id}",
			Method:  "GET",
			Handler: GetTaskByIDHandler(conn),
		},
		{
			Path:    "/tasks",
			Method:  "POST",
			Handler: PostTaskHandler(conn),
		},
	}

	r := chi.NewRouter()

	for _, route := range routes {
		switch route.Method {
		case "GET":
			r.Get(route.Path, route.Handler)
		case "POST":
			r.Post(route.Path, route.Handler)
		}
	}

	return r
}
