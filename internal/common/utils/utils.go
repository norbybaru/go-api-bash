package utils

import "github.com/segmentio/ksuid"

func GenUUID() ksuid.KSUID {
	return ksuid.New()
}

func GenUUIDString() string {
	id := ksuid.New()

	return id.String()
}
