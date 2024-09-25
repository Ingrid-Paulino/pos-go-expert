package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	user, err := NewUser("Ingrid Paulino", "i@p.com", "123456")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Ingrid Paulino", user.Name)
	assert.Equal(t, "i@p.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Ingrid Paulino", "i@p.com", "123456")

	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.False(t, user.ValidatePassword("12345"))
	assert.NotEqual(t, "123456", user.Password)
}
