package security

import (
	"log"

	"time"

	"github.com/ksrkelvin/diinoTools/pkg/models"
)

// IsProhibitedPath - Verifica se o caminho é proibido
func (p *Security) IsProhibitedPath(path string) (isProhibited bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	prohibitedPaths, err := p.DB.GetProhibitedPaths(path)
	if err != nil {
		log.Printf("Failed to get prohibited paths: %v", err)
		return true, err
	}
	if prohibitedPaths.Path != "" && prohibitedPaths.IsProhibited {
		return true, err
	}
	return false, err
}

// IsBlockedIP - Verifica se o IP é blockeado
func (p *Security) IsBlockedIP(ip string) (isProhibited bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	blockedIP, err := p.DB.CheckIP(ip)
	if err != nil {
		log.Printf("Failed to get prohibited paths: %v", err)
		return true, err
	}
	if blockedIP.IP != "" {
		return true, err
	}
	return false, err
}

// UpsertPath - Registra um novo caminho
func (p *Security) UpsertPath(path string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	err = p.DB.UpsertPath(path)
	if err != nil {
		log.Printf("Failed to insert path: %v", err)
		return err
	}
	return err
}

// BlockIP - Bloqueia o IP
func (p *Security) BlockIP(ip string, path string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	ipToBlock := models.BlockedIPsStruct{
		IP:        ip,
		Path:      path,
		Timestamp: time.Now(),
	}
	err = p.DB.BlockIP(ipToBlock)
	if err != nil {
		log.Printf("Failed to block IP: %v", err)
	}
	return
}
