package request

import (
	"github.com/AdairHdz/OTW-Rest-API/entity"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Names        string `validate:"required,max=100,alpha"`
	LastName     string `validate:"required,max=100,alpha"`
	EmailAddress string `validate:"required,email,max=254"`
	Password     string `validate:"required,min=8,securepass,max=150"`
	UserType     int    `validate:"oneof=1 2"`
	StateID      string `validate:"required,uuid4"`
	BusinessName string `validate:"required_if=UserType 1"`
	BusinessPicture string
}

func (u *User) ToEntity() (sr *entity.ServiceRequester, sp *entity.ServiceProvider, err error) {
	var hashedPassword []byte
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	if u.UserType == entity.SERVICE_PROVIDER {
		sp = &entity.ServiceProvider{
			EntityUUID: entity.EntityUUID{
				ID: uuid.NewV4().String(),
			},
			BusinessName: u.BusinessName,
			BusinessPicture: u.BusinessPicture,
			User: entity.User{
				EntityUUID: entity.EntityUUID{
					ID: uuid.NewV4().String(),
				},
				Names:    u.Names,
				Lastname: u.LastName,
				StateID:  u.StateID,
				Account: entity.Account{
					EntityUUID: entity.EntityUUID{
						ID: uuid.NewV4().String(),
					},
					EmailAddress: u.EmailAddress,
					Password:     string(hashedPassword),
					UserType:     u.UserType,
					Verified:     false,
				},
				Score: entity.Score{
					EntityUUID: entity.EntityUUID{
						ID: uuid.NewV4().String(),
					},					
					AverageScore: 0,
					MaxTotalPossible: 0,
					ObtainedPoints: 0,
				},
			},
		}
		return
	}

	sr = &entity.ServiceRequester{
		EntityUUID: entity.EntityUUID{
			ID: uuid.NewV4().String(),
		},
		User: entity.User{
			EntityUUID: entity.EntityUUID{
				ID: uuid.NewV4().String(),
			},
			Names:    u.Names,
			Lastname: u.LastName,
			StateID:  u.StateID,
			Account: entity.Account{
				EntityUUID: entity.EntityUUID{
					ID: uuid.NewV4().String(),
				},
				EmailAddress: u.EmailAddress,
				Password:     string(hashedPassword),
				UserType:     u.UserType,
				Verified:     false,
			},		
		},
	}
	
	return
}
