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
		return false
	}
	return true
}

// BlockIP - Bloqueia o IP
func (p *Security) BlockIP(ip, path string) (err error) {

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
