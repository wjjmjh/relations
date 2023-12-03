package config

type Neo4jConfig struct {
	Uri      string `json:"uri"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SvcConfig struct {
	SvcHost     string      `json:"svc_host"`
	SvcPort     int         `json:"svc_port"`
	Neo4jConfig Neo4jConfig `json:"neo4j_config"`
}

func init() {}
