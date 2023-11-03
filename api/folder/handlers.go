package folder

import (
	"net/http"

	"github.com/chtushar/toggler/api/app"
	"github.com/chtushar/toggler/api/responses"
	"github.com/chtushar/toggler/db"
	"github.com/chtushar/toggler/db/queries"
	"github.com/chtushar/toggler/utils"
	"github.com/labstack/echo/v4"
)

func handleCreateFolder(c echo.Context) error {
	var (
		app = c.Get("app").(*app.App)
		orgUUID = c.Param("orgUUID")
		req = &struct{
			Name string `json:"name" validate:"required,min=3"`
		}{}
	)

	ok, err := utils.IsValidUUID(orgUUID)

	if !ok {
		app.Log.Println("Failed to parse org uuid", err)
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	folder, err := db.WithDBTransaction[queries.Folder](app, c.Request().Context(), func(q *queries.Queries) (*queries.Folder, error) {
		org, err := q.GetOrganizationByUUID(c.Request().Context(), orgUUID)

		if err != nil {
			app.Log.Println("Couldn't get the org", err)
			return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		}

		folder, err := q.CreateFolder(c.Request().Context(), queries.CreateFolderParams{
			Name: req.Name,
			OrgID: org.ID,
		})

		if err != nil {
			app.Log.Println("Couldn't create folder", err)
			return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		}

		return &folder, nil
	})

	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, responses.ResponseType{
		Success: true,
		Data: folder,
		Error: nil,
	})
	return nil
}

func handleGetAllFolders(c echo.Context) error {
	var (
		app = c.Get("app").(*app.App)
		orgUUID = c.Param("orgUUID")
	)

	ok, err := utils.IsValidUUID(orgUUID)

	if !ok {
		app.Log.Println("Failed to parse org uuid", err)
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}

	org, err := app.Q.GetOrganizationByUUID(c.Request().Context(), orgUUID)

	if err != nil {
		app.Log.Println("Couldn't get the org", err)
		return echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
	}

	folders, err := app.Q.GetOrgFolders(c.Request().Context(), org.ID)

	if err != nil {
		app.Log.Println(" Couldn't get the folders", err)
		return echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
	}

	c.JSON(http.StatusOK, responses.ResponseType{
		Success: true,
		Data: folders,
		Error: nil,
	})
	return nil
}

func handleUpdateFolder(c echo.Context) error {
	var (
		app = c.Get("app").(*app.App)
		folderUUID = c.Param("folderUUID")
		req = &struct{
			Name string `json:"name" validate:"omitempty,min=3"`
		}{}
	)

	ok, err := utils.IsValidUUID(folderUUID)

	if !ok {
		app.Log.Println("Failed to parse folder uuid", err)
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}

	if err := c.Bind(req); err != nil {
		app.Log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}

	if err := c.Validate(req); err != nil {
		app.Log.Println(err)
		return err
	}

	_, err = db.WithDBTransaction[bool](app, c.Request().Context(), func(q *queries.Queries) (*bool, error) {
		if req.Name != "" {
			err := q.UpdateFolderName(c.Request().Context(), queries.UpdateFolderNameParams{
				Name: req.Name,
				Uuid: folderUUID,
			})

			if err != nil {
				app.Log.Println("Couldn't update the folder name")
				return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
			}
		}
		
		ok := true
		return &ok, nil
	})

	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, responses.ResponseType{
		Success: true,
		Data: "Folder updated",
		Error: nil,
	})
	return nil
}

func handleGetAllFlagsGroups(c echo.Context) error {
	return nil
}