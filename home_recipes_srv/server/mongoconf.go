package server

// MongoConf DTO
type MongoConf struct {
	host string
	port string
}

// GetHost - returns the host
func (mc *MongoConf) GetHost() string {
	return mc.host
}

// GetPort - returns the port
func (mc *MongoConf) GetPort() string {
	return mc.port
}
