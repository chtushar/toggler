package folder

import "github.com/labstack/echo/v4"

func FolderRoutes(g *echo.Group)  {
	folder := g.Group("/:orgUUID/folders")

	folder.GET("", handleGetAllFolders)
	folder.POST("/create", handleCreateFolder)
	folder.POST("/:folderUUID/update", handleUpdateFolder)
}
