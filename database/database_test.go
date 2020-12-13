package database

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestNewDatabase(t *testing.T) {
	g := NewGomegaWithT(t)

}

func FakeDbConfig() Config {
	return Config{
		Host:     DBHostTest,
		Port:     DBPortTest,
		User:     DBUserTest,
		Password: DBPassTest,
		Database: DBNameTest,
	}
}
