package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/andreis3/api-golang/internal/entity"
)

func Test_CREATE_USER(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	assert.Nil(t, err)
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("John Dae", "j@j.com", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, user.ID).Error
	assert.Nil(t, err)

}

func Test_FIND_USER_BY_EMAIL(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	assert.Nil(t, err)
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("John Dae", "j@j.com", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail("j@j.com")
	assert.Nil(t, err)
	assert.NotNil(t, userFound)

	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)

}
