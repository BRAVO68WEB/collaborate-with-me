package repository

import (
	"context"
	"errors"
	"time"

	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkSpaceRepository interface {
	CreateWorkSpace(
		userID string,
		workspaceName string,
	) (models.Workspace, error)
	UpdateWorkSpace(
		userID string,
		workspaceID string,
		workspaceName string,
	) (models.Workspace, error)
	GetWorkSpaceByID(
		ID string,
	) (models.Workspace, error)
	GetWorkSpacesByUserID(
		page int,
		userID string,
	) ([]models.Workspace, error)
	GetAllWorkSpaces(
		page int,
	) ([]models.Workspace, error)
	AddCollaborator(
		workspaceID string,
		collaboratorID string,
	) (models.Workspace, error)
	RemoveCollaborator(
		workspaceID string,
		collaboratorID string,
	) (models.Workspace, error)
	DeleteWorkSpace(
		workspaceID string,
	) (models.Workspace, error)
	SetWorkSpaceVisibility(
		workspaceID string,
		isPublic bool,
	) (models.Workspace, error)
	AddExcalidrawObject(
		workspaceID string,
		excalidrawObject interface{},
	) (models.Workspace, error)
	UpdateExcalidrawObject(
		workspaceID string,
		excalidrawObject interface{},
	) (models.Workspace, error)
	DeleteExcaldrawObject(
		workspaceID string,
		excalidrawObject interface{},
	) (models.Workspace, error)
	DeleteAllExcalidrawObjects(
		workspaceID string,
	) (models.Workspace, error)
}

type workspaceRepository struct {
	workspace *mongo.Collection
}

func NewWorkspaceRepository(
	col *mongo.Collection,
) WorkSpaceRepository {
	return &workspaceRepository{
		workspace: col,
	}
}

func (r *workspaceRepository) CreateWorkSpace(
	userID string,
	workspaceName string,
) (models.Workspace, error) {
	id := primitive.NewObjectID()
	_userID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return models.Workspace{}, err
	}

	workspace := models.Workspace{
		ID:                id,
		Name:              workspaceName,
		IsActive:          true,
		IsPublic:          false,
		Owner:             _userID,
		Collaborators:     []primitive.ObjectID{},
		ExcalidrawObjects: []models.ExcaliObjects{},
		CreatedAt:         primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:         primitive.NewDateTimeFromTime(time.Now()),
	}

	_, err = r.workspace.InsertOne(
		context.Background(),
		workspace,
	)

	if err != nil {
		return models.Workspace{}, err
	}

	return workspace, nil
}
func (r *workspaceRepository) GetWorkSpaceByID(
	ID string,
) (models.Workspace, error) {
	_ID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return models.Workspace{}, err
	}

	var workspace models.Workspace

	err = r.workspace.FindOne(
		context.Background(),
		bson.M{"_id": _ID},
	).Decode(&workspace)

	if err != nil {
		return models.Workspace{}, err
	}

	return workspace, nil
}

func (r *workspaceRepository) GetWorkSpacesByUserID(
	page int,
	userID string,
) ([]models.Workspace, error) {
	var workspaces []models.Workspace
	_userID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return []models.Workspace{}, err
	}

	cursor, err := r.workspace.Find(
		context.Background(),
		bson.M{"owner._id": _userID},
	)

	if err != nil {
		return []models.Workspace{}, err
	}

	err = cursor.All(context.Background(), &workspaces)

	if err != nil {
		return []models.Workspace{}, err
	}

	return workspaces, nil
}

func (r *workspaceRepository) GetAllWorkSpaces(
	page int,
) ([]models.Workspace, error) {
	var workspaces []models.Workspace

	cursor, err := r.workspace.Find(
		context.Background(),
		bson.M{},
	)

	if err != nil {
		return []models.Workspace{}, err
	}

	err = cursor.All(context.Background(), &workspaces)

	if err != nil {
		return []models.Workspace{}, err
	}

	return workspaces, nil
}

func (r *workspaceRepository) AddCollaborator(
	workspaceID string,
	collaboratorID string,
) (models.Workspace, error) {
	_workspaceID, err := primitive.ObjectIDFromHex(workspaceID)
	_collaboratorID, err := primitive.ObjectIDFromHex(collaboratorID)
	if err != nil {
		return models.Workspace{}, err
	}

	workspace, err := r.GetWorkSpaceByID(workspaceID)
	if err != nil {
		return models.Workspace{}, err
	}

	workspace.Collaborators = append(workspace.Collaborators, _collaboratorID)
	_, err = r.workspace.UpdateOne(
		context.Background(),
		bson.M{"_id": _workspaceID},
		bson.M{"$set": workspace},
	)
	if err != nil {
		return models.Workspace{}, err
	}

	return workspace, nil
}

func (r *workspaceRepository) RemoveCollaborator(
	workspaceID string,
	collaboratorID string,
) (models.Workspace, error) {
	_workspaceID, err := primitive.ObjectIDFromHex(workspaceID)
	_collaboratorID, err := primitive.ObjectIDFromHex(collaboratorID)
	if err != nil {
		return models.Workspace{}, err
	}
	workspace, err := r.GetWorkSpaceByID(workspaceID)
	if err != nil {
		return models.Workspace{}, err
	}
	workspace.Collaborators = []primitive.ObjectID{}
	for _, collaborator := range workspace.Collaborators {
		if collaborator != _collaboratorID {
			workspace.Collaborators = append(workspace.Collaborators, collaborator)
		}
	}
	_, err = r.workspace.UpdateOne(
		context.Background(),
		bson.M{"_id": _workspaceID},
		bson.M{"$set": workspace},
	)
	if err != nil {
		return models.Workspace{}, err
	}
	return workspace, nil
}

func (r *workspaceRepository) DeleteWorkSpace(
	workspaceID string,
) (models.Workspace, error) {
	_workspaceID, err := primitive.ObjectIDFromHex(workspaceID)
	if err != nil {
		return models.Workspace{}, err
	}
	workspace, err := r.GetWorkSpaceByID(workspaceID)
	if err != nil {
		return models.Workspace{}, err
	}
	_, err = r.workspace.DeleteOne(
		context.Background(),
		bson.M{"_id": _workspaceID},
	)

	if err != nil {
		return models.Workspace{}, err
	}

	return workspace, nil
}

func (r *workspaceRepository) AddExcalidrawObject(
	workspaceID string,
	excalidrawObject interface{},
) (models.Workspace, error) {
	_workspaceID, err := primitive.ObjectIDFromHex(workspaceID)
	if err != nil {
		return models.Workspace{}, err
	}
	workspace, err := r.GetWorkSpaceByID(workspaceID)
	if err != nil {
		return models.Workspace{}, err
	}
	if obj, ok := excalidrawObject.(models.ExcaliObjects); ok {
		workspace.ExcalidrawObjects = append(workspace.ExcalidrawObjects, obj)
	} else {
		return models.Workspace{}, errors.New("invalid type for excalidrawObject")
	}
	_, err = r.workspace.UpdateOne(
		context.Background(),
		bson.M{"_id": _workspaceID},
		bson.M{"$set": workspace},
	)
	if err != nil {
		return models.Workspace{}, err
	}
	return workspace, nil
}

func (r *workspaceRepository) UpdateExcalidrawObject(
	workspaceID string,
	excalidrawObject interface{},
) (models.Workspace, error) {
	_workspaceID, err := primitive.ObjectIDFromHex(workspaceID)
	if err != nil {
		return models.Workspace{}, err
	}
	workspace, err := r.GetWorkSpaceByID(workspaceID)
	if err != nil {
		return models.Workspace{}, err
	}
	for i, excalidrawObject := range workspace.ExcalidrawObjects {
		if excalidrawObject == excalidrawObject {
			workspace.ExcalidrawObjects[i] = excalidrawObject
		}
	}
	_, err = r.workspace.UpdateOne(
		context.Background(),
		bson.M{"_id": _workspaceID},
		bson.M{"$set": workspace},
	)
	if err != nil {
		return models.Workspace{}, err
	}
	return workspace, nil
}

func (r *workspaceRepository) DeleteExcaldrawObject(
	workspaceID string,
	excalidrawObject interface{},
) (models.Workspace, error) {
	_workspaceID, err := primitive.ObjectIDFromHex(workspaceID)
	if err != nil {
		return models.Workspace{}, err
	}
	workspace, err := r.GetWorkSpaceByID(workspaceID)
	if err != nil {
		return models.Workspace{}, err
	}
	for i, excalidrawObject := range workspace.ExcalidrawObjects {
		if excalidrawObject == excalidrawObject {
			workspace.ExcalidrawObjects = append(workspace.ExcalidrawObjects[:i], workspace.ExcalidrawObjects[i+1:]...)
		}
	}
	_, err = r.workspace.UpdateOne(
		context.Background(),
		bson.M{"_id": _workspaceID},
		bson.M{"$set": workspace},
	)
	if err != nil {
		return models.Workspace{}, err
	}
	return workspace, nil
}

func (r *workspaceRepository) DeleteAllExcalidrawObjects(
	workspaceID string,
) (models.Workspace, error) {
	_workspaceID, err := primitive.ObjectIDFromHex(workspaceID)
	if err != nil {
		return models.Workspace{}, err
	}
	workspace, err := r.GetWorkSpaceByID(workspaceID)
	if err != nil {
		return models.Workspace{}, err
	}
	workspace.ExcalidrawObjects = []models.ExcaliObjects{}
	_, err = r.workspace.UpdateOne(
		context.Background(),
		bson.M{"_id": _workspaceID},
		bson.M{"$set": workspace},
	)
	if err != nil {
		return models.Workspace{}, err
	}
	return workspace, nil
}

func (r *workspaceRepository) SetWorkSpaceVisibility(
	workspaceID string,
	isPublic bool,
) (models.Workspace, error) {
	_workspaceID, err := primitive.ObjectIDFromHex(workspaceID)
	if err != nil {
		return models.Workspace{}, err
	}
	workspace, err := r.GetWorkSpaceByID(workspaceID)
	if err != nil {
		return models.Workspace{}, err
	}
	workspace.IsPublic = isPublic
	_, err = r.workspace.UpdateOne(
		context.Background(),
		bson.M{"_id": _workspaceID},
		bson.M{"$set": workspace},
	)
	if err != nil {
		return models.Workspace{}, err
	}
	return workspace, nil
}

func (r *workspaceRepository) UpdateWorkSpace(
	userID string,
	workspaceID string,
	workspaceName string,
) (models.Workspace, error) {
	_workspaceID, err := primitive.ObjectIDFromHex(workspaceID)
	if err != nil {
		return models.Workspace{}, err
	}
	workspace, err := r.GetWorkSpaceByID(workspaceID)
	if err != nil {
		return models.Workspace{}, err
	}
	workspace.Name = workspaceName
	_, err = r.workspace.UpdateOne(
		context.Background(),
		bson.M{"_id": _workspaceID},
		bson.M{"$set": workspace},
	)
	if err != nil {
		return models.Workspace{}, err
	}
	return workspace, nil
}
