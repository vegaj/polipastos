package auth

import (
	"github.com/gobuffalo/uuid"
)

var ()

type authProvider struct {
	resources map[string]uuid.UUID
}

//Provider is the instance that will authenticate and expend auth tokens for
//the application registered services
type Provider *authProvider

//RegisterInternalResource will associate a ID to a permission
func (p authProvider) RegisterInternalResource(permName string) uuid.UUID {

	if id, ok := p.resources[permName]; ok {
		//resource already registered
		return id
	}

	if id, err := uuid.NewV4(); err == nil {
		p.resources[permName] = id
		return id
	}

	//Unable to create
	return uuid.Nil
}
