package mail

import "github.com/ksrkelvin/diinoTools/pkg/tools"

type Mailler struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	tools    tools.Tools
}

func Init(host string, port int, user string, pass string) (mailler *Mailler, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	newMailler := &Mailler{
		Host:     host,
		Port:     port,
		Username: user,
		Password: pass,
		tools:    tools.Tools{},
	}

	return newMailler, err

}
