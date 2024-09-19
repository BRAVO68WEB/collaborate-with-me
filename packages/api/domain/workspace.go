package domain

import "github.com/BRAVO68WEB/collaborate-with-me/packages/api/db"

type WorkspaceDomain interface {
}

type workspaceDomain struct {
}

func NewWorkspaceDomain(
	db db.Connection,
) WorkspaceDomain {
	return &workspaceDomain{}
}
