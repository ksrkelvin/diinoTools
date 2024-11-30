package security

import "github.com/ksrkelvin/diinoTools/pkg/database"

// Security - Main struct for Security
type Security struct {
	DB *database.DB
}

// InitSecurity - Initialize the security connection
func InitSecurity(db *database.DB) (conn *Security) {
	return &Security{
		DB: db,
	}
}
