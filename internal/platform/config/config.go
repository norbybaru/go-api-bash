package config

func BootstrapConfig() {
	loadAppEnv()
	loadDatabaseEnv()
	loadSessionEnvConfig()
	loadJWTConfig()
}
