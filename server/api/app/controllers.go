// Code generated by goagen v1.1.0, command line:
// $ main
//
// API "fieldkit": Application Controllers
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// ExpeditionController is the controller interface for the Expedition actions.
type ExpeditionController interface {
	goa.Muxer
	Add(*AddExpeditionContext) error
	Get(*GetExpeditionContext) error
	List(*ListExpeditionContext) error
}

// MountExpeditionController "mounts" a Expedition resource controller on the given service.
func MountExpeditionController(service *goa.Service, ctrl ExpeditionController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/projects/:project/expedition", ctrl.MuxHandler("preflight", handleExpeditionOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/projects/:project/expeditions/:expedition", ctrl.MuxHandler("preflight", handleExpeditionOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/projects/:project/expeditions", ctrl.MuxHandler("preflight", handleExpeditionOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAddExpeditionContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*AddExpeditionPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Add(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleExpeditionOrigin(h)
	service.Mux.Handle("POST", "/projects/:project/expedition", ctrl.MuxHandler("Add", h, unmarshalAddExpeditionPayload))
	service.LogInfo("mount", "ctrl", "Expedition", "action", "Add", "route", "POST /projects/:project/expedition", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetExpeditionContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleExpeditionOrigin(h)
	service.Mux.Handle("GET", "/projects/:project/expeditions/:expedition", ctrl.MuxHandler("Get", h, nil))
	service.LogInfo("mount", "ctrl", "Expedition", "action", "Get", "route", "GET /projects/:project/expeditions/:expedition", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListExpeditionContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleExpeditionOrigin(h)
	service.Mux.Handle("GET", "/projects/:project/expeditions", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Expedition", "action", "List", "route", "GET /projects/:project/expeditions", "security", "jwt")
}

// handleExpeditionOrigin applies the CORS response headers corresponding to the origin.
func handleExpeditionOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "https://*.fieldkit.org") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
			}
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "https://fieldkit.org") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalAddExpeditionPayload unmarshals the request body into the context request data Payload field.
func unmarshalAddExpeditionPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &addExpeditionPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// ProjectController is the controller interface for the Project actions.
type ProjectController interface {
	goa.Muxer
	Add(*AddProjectContext) error
	Get(*GetProjectContext) error
	List(*ListProjectContext) error
	ListCurrent(*ListCurrentProjectContext) error
}

// MountProjectController "mounts" a Project resource controller on the given service.
func MountProjectController(service *goa.Service, ctrl ProjectController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/project", ctrl.MuxHandler("preflight", handleProjectOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/projects/:project", ctrl.MuxHandler("preflight", handleProjectOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/projects", ctrl.MuxHandler("preflight", handleProjectOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/user/projects", ctrl.MuxHandler("preflight", handleProjectOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAddProjectContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*AddProjectPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Add(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleProjectOrigin(h)
	service.Mux.Handle("POST", "/project", ctrl.MuxHandler("Add", h, unmarshalAddProjectPayload))
	service.LogInfo("mount", "ctrl", "Project", "action", "Add", "route", "POST /project", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetProjectContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleProjectOrigin(h)
	service.Mux.Handle("GET", "/projects/:project", ctrl.MuxHandler("Get", h, nil))
	service.LogInfo("mount", "ctrl", "Project", "action", "Get", "route", "GET /projects/:project", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListProjectContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleProjectOrigin(h)
	service.Mux.Handle("GET", "/projects", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Project", "action", "List", "route", "GET /projects", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListCurrentProjectContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.ListCurrent(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleProjectOrigin(h)
	service.Mux.Handle("GET", "/user/projects", ctrl.MuxHandler("ListCurrent", h, nil))
	service.LogInfo("mount", "ctrl", "Project", "action", "ListCurrent", "route", "GET /user/projects", "security", "jwt")
}

// handleProjectOrigin applies the CORS response headers corresponding to the origin.
func handleProjectOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "https://*.fieldkit.org") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
			}
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "https://fieldkit.org") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalAddProjectPayload unmarshals the request body into the context request data Payload field.
func unmarshalAddProjectPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &addProjectPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/swagger.yaml", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger.json", "api/public/swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "api/public/swagger/swagger.json", "route", "GET /swagger.json")

	h = ctrl.FileHandler("/swagger.yaml", "api/public/swagger/swagger.yaml")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.yaml", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "api/public/swagger/swagger.yaml", "route", "GET /swagger.yaml")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			}
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "https://*.fieldkit.org") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
			}
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "https://fieldkit.org") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// TeamController is the controller interface for the Team actions.
type TeamController interface {
	goa.Muxer
	Add(*AddTeamContext) error
	Get(*GetTeamContext) error
	List(*ListTeamContext) error
}

// MountTeamController "mounts" a Team resource controller on the given service.
func MountTeamController(service *goa.Service, ctrl TeamController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/projects/:project/expeditions/:expedition/team", ctrl.MuxHandler("preflight", handleTeamOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/projects/:project/expeditions/:expedition/teams/:team", ctrl.MuxHandler("preflight", handleTeamOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/projects/:project/expeditions/:expedition/teams", ctrl.MuxHandler("preflight", handleTeamOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAddTeamContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*AddTeamPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Add(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleTeamOrigin(h)
	service.Mux.Handle("POST", "/projects/:project/expeditions/:expedition/team", ctrl.MuxHandler("Add", h, unmarshalAddTeamPayload))
	service.LogInfo("mount", "ctrl", "Team", "action", "Add", "route", "POST /projects/:project/expeditions/:expedition/team", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetTeamContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleTeamOrigin(h)
	service.Mux.Handle("GET", "/projects/:project/expeditions/:expedition/teams/:team", ctrl.MuxHandler("Get", h, nil))
	service.LogInfo("mount", "ctrl", "Team", "action", "Get", "route", "GET /projects/:project/expeditions/:expedition/teams/:team", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListTeamContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleTeamOrigin(h)
	service.Mux.Handle("GET", "/projects/:project/expeditions/:expedition/teams", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Team", "action", "List", "route", "GET /projects/:project/expeditions/:expedition/teams", "security", "jwt")
}

// handleTeamOrigin applies the CORS response headers corresponding to the origin.
func handleTeamOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "https://*.fieldkit.org") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
			}
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "https://fieldkit.org") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalAddTeamPayload unmarshals the request body into the context request data Payload field.
func unmarshalAddTeamPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &addTeamPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// UserController is the controller interface for the User actions.
type UserController interface {
	goa.Muxer
	Add(*AddUserContext) error
	Get(*GetUserContext) error
	GetCurrent(*GetCurrentUserContext) error
	List(*ListUserContext) error
	Login(*LoginUserContext) error
	Logout(*LogoutUserContext) error
	Refresh(*RefreshUserContext) error
}

// MountUserController "mounts" a User resource controller on the given service.
func MountUserController(service *goa.Service, ctrl UserController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/user", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/user/:username", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/users", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/login", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/logout", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/refresh", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAddUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*AddUserPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Add(rctx)
	}
	h = handleUserOrigin(h)
	service.Mux.Handle("POST", "/user", ctrl.MuxHandler("Add", h, unmarshalAddUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Add", "route", "POST /user")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleUserOrigin(h)
	service.Mux.Handle("GET", "/user/:username", ctrl.MuxHandler("Get", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Get", "route", "GET /user/:username", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetCurrentUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.GetCurrent(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleUserOrigin(h)
	service.Mux.Handle("GET", "/user", ctrl.MuxHandler("GetCurrent", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "GetCurrent", "route", "GET /user", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleUserOrigin(h)
	service.Mux.Handle("GET", "/users", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "List", "route", "GET /users", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewLoginUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*AddUserPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Login(rctx)
	}
	h = handleUserOrigin(h)
	service.Mux.Handle("POST", "/login", ctrl.MuxHandler("Login", h, unmarshalLoginUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Login", "route", "POST /login")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewLogoutUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Logout(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleUserOrigin(h)
	service.Mux.Handle("POST", "/logout", ctrl.MuxHandler("Logout", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Logout", "route", "POST /logout", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRefreshUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*RefreshUserPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Refresh(rctx)
	}
	h = handleUserOrigin(h)
	service.Mux.Handle("POST", "/refresh", ctrl.MuxHandler("Refresh", h, unmarshalRefreshUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Refresh", "route", "POST /refresh")
}

// handleUserOrigin applies the CORS response headers corresponding to the origin.
func handleUserOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "https://*.fieldkit.org") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
			}
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "https://fieldkit.org") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalAddUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalAddUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &addUserPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalLoginUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalLoginUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &addUserPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalRefreshUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalRefreshUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &refreshUserPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
