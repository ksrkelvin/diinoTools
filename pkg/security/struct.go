package security

import (
	"github.com/ksrkelvin/diinoTools/pkg/database/mongo"
	securityDB "github.com/ksrkelvin/diinoTools/pkg/database/security"
)

// Security - Main struct for Security
type Security struct {
	DB *securityDB.DB
}

// InitSecurity - Initialize the security connection
func InitSecurity(db *mongo.DB) (security *Security, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	dbSecurity, err := securityDB.Init(db.Client)
	if err != nil {
		return security, err
	}

	newSecurity := &Security{
		DB: dbSecurity,
	}

	return newSecurity, err
}
