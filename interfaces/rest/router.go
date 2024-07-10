package rest

import (
	"fmt"
	"go-rest-proxy/internal/viewmodels"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// InitializeRouter initializes router and routes
func InitializeRouter() *chi.Mux {
	// create router
	r := chi.NewRouter()

	// global and recommended middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	// default route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response := viewmodels.HTTPResponseVM{
			Status:  http.StatusOK,
			Success: true,
			Message: "alive",
		}

		response.JSON(w)
	})

	// docs routes
	workDir, _ := os.Getwd()
	docsDir := http.Dir(filepath.Join(workDir, "docs"))
	FileServer(r, "/docs", docsDir)

	// API routes
	r.Group(func(r chi.Router) {
		r.Route("/api", func(r chi.Router) {
			r.Route("/v1", func(r chi.Router) {
				// place your routes here
			})
		})
	})

	return r
}

// Serve will start serving and listening to router
func Serve(port int) {
	log.Printf("[SERVER] REST server running on :%d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), InitializeRouter())
	if err != nil {
		log.Fatalf("[SERVER] REST server failed %v", err)
	}
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
