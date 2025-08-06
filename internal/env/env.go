package env

import "os"

func DBUrl() string {
	return os.Getenv("DBUrl")
}

func AppName() string {
	return os.Getenv("AppName")
}

func AppPort() string {
	return os.Getenv("AppPort")
}

func AppToken() string {
	return os.Getenv("AppToken")
}
