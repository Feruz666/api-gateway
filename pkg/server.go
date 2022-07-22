package pkg

import (
	"fmt"
	"net/http"
	"os"

	auth "github.com/Feruz666/api-gateway/pkg/auth/handlers"
	document "github.com/Feruz666/api-gateway/pkg/document/handlers"
	maps "github.com/Feruz666/api-gateway/pkg/maps/handlers"
	sensors "github.com/Feruz666/api-gateway/pkg/sensor/handlers"
	"github.com/Feruz666/api-gateway/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var MySigningKey = []byte(os.Getenv("SECRET_KEY"))

// Server serves HTTP requests and responses
type Server struct {
	config util.Config
	router *gin.Engine
}

func ServiceRoutes(config util.Config) (*Server, error) {
	server := &Server{
		config: config,
	}
	server.setupRouter()
	return server, nil
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf(("invalid Signing Method"))
				}
				aud := "billing.jwtgo.io"
				checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
				if !checkAudience {
					return nil, fmt.Errorf(("invalid aud"))
				}
				// verify iss claim
				iss := "jwtgo.io"
				checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
				if !checkIss {
					return nil, fmt.Errorf(("invalid iss"))
				}

				return MySigningKey, nil
			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {
			fmt.Fprintf(w, "No Authorization Token provided")
		}
	})
}

func (server *Server) setupRouter() {
	router := gin.Default()
	// Auth route
	router.POST("/auth", auth.CreateUser)
	router.POST("/auth/login", auth.LoginUser)

	// Sensors route
	router.POST("/sensors", sensors.CreateSensor)
	router.GET("/sensors", sensors.GetSensors)
	router.GET("/sensors/charts", sensors.GetSensorsCharts)

	// Example route
	router.POST("/doc", document.Example)
	router.GET("/getdoc", document.GetExample)
	router.DELETE("/deletedoc", document.GetExample)

	// Maps route
	router.GET("/maps/workspaces", maps.GetWorkspaces)
	router.GET("/maps/workspaces/workspace", maps.GetWorkspace)
	router.POST("/maps/workspaces/workspace", maps.CreateWorkspace)
	router.DELETE("/maps/workspaces/workspace", maps.DeleteWorkspace)
	router.GET("/maps/layers", maps.GetLayers)
	router.GET("/maps/layers/layer", maps.GetLayer)
	router.DELETE("/maps/layers/layer", maps.DeleteLayers)
	router.GET("/maps/datastores", maps.GetDatastores)
	router.GET("/maps/datastores/datastore", maps.GetDatastore)
	router.GET("/maps/datastores/datastore/exists", maps.DatastoreExists)
	router.POST("/maps/datastores/datastore", maps.CreateDatastore)
	router.DELETE("/maps/datastores/datastore", maps.DeleteDatastore)
	router.GET("/maps/layergroups", maps.GetLayerGroups)
	router.GET("/maps/layergroups/layergroup", maps.GetLayerGroup)
	router.GET("/maps/featuretypes", maps.GetFeatureTypes)
	router.GET("/maps/featuretypes/featuretype", maps.GetFeatureType)
	router.DELETE("/maps/featuretypes/featuretype", maps.DeleteFeatureType)
	router.GET("/maps/wms", maps.GetWMS)
	router.GET("/maps/styles", maps.GetStyles)

	server.router = router
}

// Start runs server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
