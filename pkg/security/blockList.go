package security

import (
	"log"

	"time"

	"github.com/ksrkelvin/diinoTools/pkg/models"
)

// IsProhibitedPath - Verifica se o caminho é proibido
func (p *Security) IsProhibitedPath(path string) (isProhibited bool) {

	prohibitedPaths, err := p.DB.Mongo.GetProhibitedPaths(path)
	if err != nil {
		log.Printf("Failed to get prohibited paths: %v", err)
		return true
	}
	if prohibitedPaths.Path != "" {
		return true
	}
	return false
}

// UpsertPath - Registra um novo caminho
func (p *Security) UpsertPath(path string) (err error) {
	err = p.DB.Mongo.UpsertPath(path)
	if err != nil {
		log.Printf("Failed to insert path: %v", err)
		return err
	}
	return err
}

// BlockIP - Bloqueia o IP
func (p *Security) BlockIP(ip string, path string) (err error) {

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
