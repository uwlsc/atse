package models

import (
	"magazine_api/lib"

	"github.com/google/uuid"
)

// Normal user profile
type UserProfileBase struct {
	UserId uuid.UUID `json:"user_id"`

	Name          *string        `json:"name"`
	Email         *string        `json:"email"`
	ContactNumber *string        `json:"contact_number"`
	Picture       *lib.SignedURL `json:"picture"`
}

type UserProfile struct {
	Base
	BaseDate
	UserProfileBase
}
