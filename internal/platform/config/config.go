package config

func BootstrapConfig() {
	loadAppEnv()
	loadDatabaseEnv()
}
