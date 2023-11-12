package define

var (
	// AppName
	AppName = "redis-client"
	// AppVersion
	AppVersion = "1.0.0"
	// AppDescription
	AppDescription = "Redis client"
	// ConfigName
	ConfigName = "redis-client.conf"
)

type Connection struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Config struct {
	Connections []*Connection `json:"connections"`
}
