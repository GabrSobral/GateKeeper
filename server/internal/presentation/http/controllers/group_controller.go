package http_controllers

import (
	"net/http"

	creategroup "github.com/gate-keeper/internal/application/services/group/create-group"
	removegroup "github.com/gate-keeper/internal/application/services/group/remove-group"
	"github.com/gate-keeper/internal/infra/database/repositories"
	http_router "github.com/gate-keeper/internal/presentation/http"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GroupController struct {
	DbPool *pgxpool.Pool
}

// Create a group to an application
func (c *GroupController) CreateGroup(writter http.ResponseWriter, request *http.Request) {
	var createGroupRequest creategroup.Request

	if err := http_router.ParseBodyToSchema(&createGroupRequest, request); err != nil {
		panic(err)
	}

	params := repositories.Params[creategroup.Request, creategroup.CreateGroupService]{
		DbPool:  c.DbPool,
		New:     creategroup.New,
		Request: createGroupRequest,
	}

	if err := repositories.WithTransaction(request.Context(), params); err != nil {
		panic(err)
	}

	http_router.SendJson(writter, nil, http.StatusCreated)
}

// Remove a group from application
func (c *GroupController) RemoveGroup(writter http.ResponseWriter, request *http.Request) {
	var removeGroupRequest removegroup.Request

	if err := http_router.ParseBodyToSchema(&removeGroupRequest, request); err != nil {
		panic(err)
	}

	params := repositories.Params[removegroup.Request, removegroup.RemoveGroupService]{
		DbPool:  c.DbPool,
		New:     removegroup.New,
		Request: removeGroupRequest,
	}

	if err := repositories.WithTransaction(request.Context(), params); err != nil {
		panic(err)
	}

	http_router.SendJson(writter, nil, http.StatusNoContent)
}
