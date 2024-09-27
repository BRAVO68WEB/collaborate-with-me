package domain

import (
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/db"
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/models"
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkspaceDomain interface {
	CreateWorkspace(name string, userID string, isPublic bool) (models.Workspace, error)
	UpdateWorkspace(userID string, workspaceID string, w_name string) (models.Workspace, error)
	DeleteWorkspace(id string) error
	AddUserToWorkspace(workspaceID string, userID string) (models.Workspace, error)
	RemoveUserFromWorkspace(workspaceID string, userID string) (models.Workspace, error)
	AddExcalidrawObject(workspaceID string, object models.ExcaliObjects) (models.Workspace, error)
	RemoveExcalidrawObject(workspaceID string, objectID string) (models.Workspace, error)
	GetWorkspace(workspaceID string) (models.Workspace, error)
	GetWorkspaces(userID string) ([]models.Workspace, error)
	GetWorkspaceUsers(
		workspaceID string,
		u *userDomain,
	) ([]models.User, error)
	GetWorkspaceExcalidrawObjects(workspaceID string) ([]models.ExcaliObjects, error)
	GetWorkspaceExcalidrawObject(workspaceID string, objectID string) (models.ExcaliObjects, error)
}

type workspaceDomain struct {
	workspace_repo repository.WorkSpaceRepository
}

func NewWorkspaceDomain(
	db db.Connection,
) WorkspaceDomain {
	return &workspaceDomain{
		workspace_repo: repository.NewWorkspaceRepository(
			db.GetCollection("workspace"),
		),
	}
}

func (d *workspaceDomain) CreateWorkspace(name string, userID string, isPublic bool) (models.Workspace, error) {
	return d.workspace_repo.CreateWorkSpace(userID, name)
}

func (d *workspaceDomain) UpdateWorkspace(userID string, workspaceID string, w_name string) (models.Workspace, error) {
	return d.workspace_repo.UpdateWorkSpace(userID, workspaceID, w_name)
}

func (d *workspaceDomain) DeleteWorkspace(id string) error {
	_, err := d.workspace_repo.DeleteWorkSpace(id)
	if err != nil {
		return err
	}
	return nil
}

func (d *workspaceDomain) AddUserToWorkspace(workspaceID string, userID string) (models.Workspace, error) {
	return d.workspace_repo.AddCollaborator(workspaceID, userID)
}

func (d *workspaceDomain) RemoveUserFromWorkspace(workspaceID string, userID string) (models.Workspace, error) {
	return d.workspace_repo.RemoveCollaborator(workspaceID, userID)
}

func (d *workspaceDomain) AddExcalidrawObject(workspaceID string, object models.ExcaliObjects) (models.Workspace, error) {
	return d.workspace_repo.AddExcalidrawObject(workspaceID, object)
}

func (d *workspaceDomain) RemoveExcalidrawObject(workspaceID string, objectID string) (models.Workspace, error) {
	return d.workspace_repo.DeleteExcaldrawObject(workspaceID, objectID)
}

func (d *workspaceDomain) GetWorkspace(workspaceID string) (models.Workspace, error) {
	return d.workspace_repo.GetWorkSpaceByID(workspaceID)
}

func (d *workspaceDomain) GetWorkspaces(userID string) ([]models.Workspace, error) {
	return d.workspace_repo.GetAllWorkSpaces(0)
}

func (d *workspaceDomain) GetWorkspaceUsers(workspaceID string, u *userDomain) ([]models.User, error) {
	w, err := d.workspace_repo.GetWorkSpaceByID(workspaceID)

	if err != nil {
		return nil, err
	}

	var users []models.User

	for _, user := range w.Collaborators {
		user, err := u.GetUserByID(user.Hex())
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (d *workspaceDomain) GetWorkspaceExcalidrawObjects(workspaceID string) ([]models.ExcaliObjects, error) {
	w, err := d.workspace_repo.GetWorkSpaceByID(workspaceID)

	if err != nil {
		return nil, err
	}

	return w.ExcalidrawObjects, nil
}

func (d *workspaceDomain) GetWorkspaceExcalidrawObject(workspaceID string, objectID string) (models.ExcaliObjects, error) {
	w, err := d.workspace_repo.GetWorkSpaceByID(workspaceID)

	if err != nil {
		return models.ExcaliObjects{}, err
	}

	primitiveObjectID, err := primitive.ObjectIDFromHex(objectID)
	if err != nil {
		return models.ExcaliObjects{}, err
	}

	for _, object := range w.ExcalidrawObjects {
		if object.ID == primitiveObjectID {
			return object, nil
		}
	}

	return models.ExcaliObjects{}, nil
}
