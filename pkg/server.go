package pkg

import (
	auth "github.com/Feruz666/api-gateway/pkg/auth/handlers"
	document "github.com/Feruz666/api-gateway/pkg/document/handlers"
	maps "github.com/Feruz666/api-gateway/pkg/maps/handlers"
	"github.com/Feruz666/api-gateway/util"
	"github.com/gin-gonic/gin"
)

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

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/auth", auth.CreateUser)
	router.POST("/auth/login", auth.LoginUser)

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

	server.router = router
}

// Start runs server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
