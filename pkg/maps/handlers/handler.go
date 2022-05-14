package maps

import (
	"github.com/Feruz666/api-gateway/util"
	"github.com/gin-gonic/gin"
)

// Workspace handlers
func GetWorkspaces(ctx *gin.Context) {
	util.GetGateWayUrl("http://localhost:3001/0.0/workspaces/", ctx)
}

func CreateWorkspace(ctx *gin.Context) {
	util.PostGateWayUrl("http://localhost:3001/0.0/workspaces/workspace", ctx)
}

func GetWorkspace(ctx *gin.Context) {
	util.GetGateWayUrl("http://localhost:3001/0.0/workspaces/workspace", ctx)
}

func DeleteWorkspace(ctx *gin.Context) {
	util.DeleteGateWayUrl("http://localhost:3001/0.0/workspaces/workspace", ctx)
}

// Layers handlers
func GetLayers(ctx *gin.Context) {
	util.GetGateWayUrl("http://localhost:3001/0.0/layers", ctx)
}

func GetLayer(ctx *gin.Context) {
	util.GetGateWayUrl("http://localhost:3001/0.0/layers/layer", ctx)
}

func DeleteLayers(ctx *gin.Context) {
	util.DeleteGateWayUrl("http://localhost:3001/0.0/layers/layer", ctx)
}

// Datastore handlers
func GetDatastores(ctx *gin.Context) {
	util.GetGateWayUrl("http://localhost:3001/0.0/datastores", ctx)
}

func GetDatastore(ctx *gin.Context) {
	util.GetGateWayUrl("http://localhost:3001/0.0/datastores/datastore", ctx)
}

func DatastoreExists(ctx *gin.Context) {
	util.GetGateWayUrl("http://localhost:3001/0.0/datastores/datastore/exists", ctx)
}

func CreateDatastore(ctx *gin.Context) {
	util.PostGateWayUrl("http://localhost:3001/0.0/datastores/datastore", ctx)
}

func DeleteDatastore(ctx *gin.Context) {
	util.DeleteGateWayUrl("http://localhost:3001/0.0/datastores/datastore", ctx)
}

// LayerGroups handlers
func GetLayerGroups(ctx *gin.Context) {
	util.GetGateWayUrl("http://localhost:3001/0.0/layergroups/", ctx)
}

func GetLayerGroup(ctx *gin.Context) {
	util.GetGateWayUrl("http://localhost:3001/0.0/layergroups/layergroup", ctx)
}

// FeatureType handlers
func GetFeatureTypes(ctx *gin.Context) {
	util.GetGateWayUrl("http://localhost:3001/0.0/featuretypes", ctx)
}

func GetFeatureType(ctx *gin.Context) {
	util.GetGateWayUrl("http://localhost:3001/0.0/featuretypes/featuretype", ctx)
}

func DeleteFeatureType(ctx *gin.Context) {
	util.DeleteGateWayUrl("http://localhost:3001/0.0/featuretypes/featuretype", ctx)
}
