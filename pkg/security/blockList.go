package security

import (
	"log"

	"time"

	"github.com/ksrkelvin/diinoTools/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// IsProhibitedPath - Verifica se o caminho é proibido
func (p *Security) IsProhibitedPath(path string) (isProhibited bool) {

	prohibitedPaths, err := p.DB.Mongo.GetProhibitedPaths(path)
	if err != nil {
		log.Printf("Failed to get prohibited paths: %v", err)
		return true
	}
	if prohibitedPaths.Path != "" {
		return false
	}
	return true
}

// BlockIP - Bloqueia o IP
func (p *Security) BlockIP(ip, path string, client *mongo.Client) (err error) {

	ipToBlock := models.BlockedIPsStruct{
		IP:        ip,
		Path:      path,
		Timestamp: time.Now(),
	}
	err = p.DB.Mongo.BlockIP(ipToBlock)
	if err != nil {
		log.Printf("Failed to block IP: %v", err)
	}
	return
}

// ValidateConnection - Valida a conexão
func (p *Security) ValidateConnection(ip, path string) (isBlocked bool) {
	if p.IsProhibitedPath(path) {
		p.BlockIP(ip, path, p.DB.Mongo.Client)
		return true
	}
	return false
}
