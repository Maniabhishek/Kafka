package messageconfig

type BrokerConfig struct {
	Url string
	Username string
	Password string
	EnableAuth bool
	InsecureSkipVerify bool
}