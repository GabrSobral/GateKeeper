package signin

import (
	application_utils "github.com/gate-keeper/internal/application/utils"
	"github.com/gate-keeper/internal/domain/entities"
)

func assignTokenParams(userProfile entities.UserProfile, user entities.ApplicationUser) (string, error) {

	return application_utils.CreateToken(
		application_utils.JWTClaims{
			UserID:        user.ID,
			Email:         user.Email,
			FirstName:     userProfile.FirstName,
			LastName:      userProfile.LastName,
			DisplayName:   userProfile.DisplayName,
			ApplicationID: user.ApplicationID,
		},
	)
}
