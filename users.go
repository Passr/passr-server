package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/manyminds/api2go/jsonapi"
	"passr/authentication"
)

//-----------------------------------------------------------------
// User
//-----------------------------------------------------------------
type User struct {
	ID        int    `json:"-" gorm:"primary_key" sql:"AUTO_INCREMENT"`
	UserName  string `sql:"type:text"`
	UserID    int    `json:"-"`
	Password  string `sql:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

//-----------------------------------------------------------------
// POST /api/users
//-----------------------------------------------------------------

//-----------------------------------------------------------------
// POST /api/users
//-----------------------------------------------------------------
