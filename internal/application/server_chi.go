package application

import (
	"database/sql"
	"net/http"

	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal/handler"
	md "github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal/middleware"
	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal/repository"
	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/handlers"
)

type ConfigurationServer struct {
	Addr     string
	MySQLDSN string
}

// NewServerChi creates a new instance of the server
func NewServerChi(cfg ConfigurationServer) *serverChi {
	// default config
	defaultCfg := ConfigurationServer{
		Addr:     ":8080",
		MySQLDSN: "",
	}
	if cfg.Addr != "" {
		defaultCfg.Addr = cfg.Addr
	}
	if cfg.MySQLDSN != "" {
		defaultCfg.MySQLDSN = cfg.MySQLDSN
	}

	return &serverChi{
		addr:     defaultCfg.Addr,
		mysqlDSN: defaultCfg.MySQLDSN,
	}
}

type serverChi struct {
	addr     string
	mysqlDSN string
}

func (s *serverChi) Run() (err error) {
	db, err := sql.Open("mysql", s.mysqlDSN)
	if err != nil {
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/api/v1", func(r chi.Router) {
		buildClientRouter(r, db)
		buildAffiliateRouter(r, db)
		buildLegalRepRouter(r, db)
		buildCredentialsRouter(r, db)
		buildLoginRouter(r, db)
	})

	corsHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	corsOrigins := handlers.AllowedOrigins([]string{"http://localhost:5173"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	err = http.ListenAndServe(s.addr, handlers.CORS(corsHeaders, corsOrigins, corsMethods)(router))
	return
}

func buildClientRouter(router chi.Router, db *sql.DB) {

	rp := repository.NewClientMysql(db)
	svc := service.NewClientDefault(rp)
	hd := handler.NewClientHandler(svc)
	rpAuth := repository.NewUserRepository(db)
	authService := service.NewAuthDefault(rpAuth)

	router.Group(func(r chi.Router) {
		r.Use(md.AuthMiddleware(authService))
		r.Get("/clients", hd.FindAll())
		r.Get("/clients/{id}", hd.FindByID())
		r.Get("/clients/nit/{nit}", hd.FindByNit())
		r.Post("/clients", hd.Create())
		r.Delete("/clients/{id}", hd.Delete())
		r.Put("/clients", hd.Update())
		r.Get("/clients/legal-rep/{idLegalRep}", hd.FindByLegalRepID())
	})
}

func buildAffiliateRouter(router chi.Router, db *sql.DB) {

	rp := repository.NewAffiliateMySql(db)
	svc := service.NewAffiliateDefault(rp)
	hd := handler.NewAffiliateHandler(svc)
	rpAuth := repository.NewUserRepository(db)
	authService := service.NewAuthDefault(rpAuth)

	router.Group(func(r chi.Router) {
		r.Use(md.AuthMiddleware(authService))
		r.Get("/affiliates", hd.FindAll())
		r.Get("/affiliates/{id}", hd.FindByID())
		r.Get("/affiliates/cc/{cc}", hd.FindByCc())
		r.Post("/affiliates", hd.Create())
		r.Delete("/affiliates/{id}", hd.Delete())
		r.Put("/affiliates", hd.Update())
		r.Get("/affiliates/client/{clientId}", hd.FindByClientId())
	})
}

func buildLegalRepRouter(router chi.Router, db *sql.DB) {

	rp := repository.NewLegalRepMySql(db)
	svc := service.NewLegalRepDefault(rp)
	hd := handler.NewLegalRepDefault(svc)
	rpAuth := repository.NewUserRepository(db)
	authService := service.NewAuthDefault(rpAuth)

	router.Group(func(r chi.Router) {
		r.Use(md.AuthMiddleware(authService))
		r.Get("/legal-reps", hd.FindAll())
		r.Get("/legal-reps/{id}", hd.FindByID())
		r.Get("/legal-reps/cc/{cc}", hd.FindByCc())
		r.Post("/legal-reps", hd.Create())
		r.Put("/legal-reps", hd.Update())
		r.Delete("/legal-reps/{id}", hd.Delete())
	})
}

func buildCredentialsRouter(router chi.Router, db *sql.DB) {

	rp := repository.NewCredentialsMySql(db)
	svc := service.NewCredentialsDefault(rp)
	hd := handler.NewCredentialsHandler(svc)
	rpAuth := repository.NewUserRepository(db)
	authService := service.NewAuthDefault(rpAuth)

	router.Group(func(r chi.Router) {
		r.Use(md.AuthMiddleware(authService))
		r.Get("/credentials", hd.FindAll())
		r.Get("/credentials/{id}", hd.FindByID())
		r.Post("/credentials", hd.Create())
		r.Delete("/credentials/{id}", hd.Delete())
		r.Put("/credentials", hd.Update())
		r.Get("/credentials/client/{idClient}", hd.FindByClient())
	})
}

func buildLoginRouter(router chi.Router, db *sql.DB) {
	rp := repository.NewUserRepository(db)
	svc := service.NewAuthDefault(rp)
	hd := handler.NewAuthHandler(svc)

	router.Post("/login", hd.Login())
}
