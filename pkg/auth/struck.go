package auth

import "github.com/ksrkelvin/diinoTools/pkg/tools"

// Auth - Estrutura de autenticação
type Auth struct {
	Secret string
	tools  tools.Tools
}

// Init - Inicializa a autenticação
func Init(secret string) (auth *Auth, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	newAuth := &Auth{
		Secret: secret,
		tools:  tools.Tools{},
	}
	return newAuth, err
}
