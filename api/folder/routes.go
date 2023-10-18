package folder

import "github.com/labstack/echo/v4"

func FolderRoutes(g *echo.Group)  {
	folder := g.Group("/folder")

	folder.POST("/create", handleCreateFolder)
	folder.DELETE("/:folderUUID", handleDeleteFolder)
	folder.GET("/", handleGetAllFolders)
	folder.GET("/:folderUUID/flagsgroups", handleGetAllFlagsGroups)
}
