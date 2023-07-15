package config

// Environment
var Environment map[string]interface{} = map[string]interface{}{
	"app_name":    "ddd-demo-service",
	"app_port":    4321,
	"app_enpoint": "/v1/demo",
	"db_host":     "192.168.100.10",
	"db_port":     3310,
	"db_user":     "root",
	"db_password": "password",
	"db_name":     "db_demo",
}
