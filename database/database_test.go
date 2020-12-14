package database

import (
	"testing"

	. "challenge-SmartMEI/helper_tests"

	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
)

func TestNewDatabase(t *testing.T) {
	g := NewGomegaWithT(t)
	nwDb, err := NewDatabase(FakeDbConfig())
	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(nwDb).ShouldNot(BeNil())
	coll := MockCollection(g, "teste-coll")
	_, err = coll.InsertOne(nil, bson.M{"Test": "init"})
	g.Expect(err).ShouldNot(HaveOccurred())
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
