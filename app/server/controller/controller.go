package controller

import db "github.com/NeptuneG/go-back/db/sqlc"

type Controller struct {
	store *db.Store
}

func NewController(store *db.Store) *Controller {
	return &Controller{store: store}
}