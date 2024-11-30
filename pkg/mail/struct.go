package mail

// Mailler - Main struct for Mailler
type Mailler struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Init - Initialize the mail connection
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
	}

	return newMailler, err

}
