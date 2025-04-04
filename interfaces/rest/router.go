package rest

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"

	"go-rest-proxy/interfaces/rest/middlewares/cors"
	jwt "go-rest-proxy/interfaces/rest/middlewares/iam"
	"go-rest-proxy/internal/viewmodels"
	cartController "go-rest-proxy/module/cart/controller"
	iamController "go-rest-proxy/module/iam/controller"
	productController "go-rest-proxy/module/product/controller"
)

// InitializeRouter initializes router and routes
func InitializeRouter() *chi.Mux {
	// initialize controllers
	iamCommandController := iamController.IAMCommandController{}
	cartQueryController := cartController.CartQueryController{}
	productQueryController := productController.ProductQueryController{}

	// create router
	r := chi.NewRouter()

	// global and recommended middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(cors.Init().Handler)

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
	r.Group(func(r chi.Router) {
		r.Use(middleware.BasicAuth(os.Getenv("API_NAME"), map[string]string{
			"sudo": os.Getenv("OPENAPI_DOCS_PASSWORD"),
		}))

		workDir, _ := os.Getwd()
		docsDir := http.Dir(filepath.Join(workDir, "docs"))
		FileServer(r, "/docs", docsDir)
	})

	// API routes
	r.Group(func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)

			// iam endpoints
			r.Group(func(r chi.Router) {
				r.Route("/iam", func(r chi.Router) {
					r.Post("/generate-token", iamCommandController.GenerateToken)
				})
			})

			r.Group(func(r chi.Router) {
				r.Use(jwtauth.Verifier(tokenAuth))
				r.Use(jwt.JWTAuthMiddleware)

				// cart endpoints
				r.Route("/cart", func(r chi.Router) {
					r.Get("/list", cartQueryController.GetCarts)
					r.Get("/{id}", cartQueryController.GetCartByID)
				})

				// product endpoints
				r.Route("/product", func(r chi.Router) {
					r.Get("/list", productQueryController.GetProducts)
					r.Get("/{id}", productQueryController.GetProductByID)
				})
			})
		})
	})

	return r
}

// Serve will start serving and listening to router
func Serve(port int) {
	log.Printf("[PROXY-SERVER] REST server running on :%d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), InitializeRouter())
	if err != nil {
		log.Fatalf("[PROXY-SERVER] REST server failed %v", err)
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
