package helper_tests

import (
	"fmt"

	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DBHostTest = "localhost"
const DBNameTest = "smartm-db"
const DBPassTest = "123qwe"
const DBUserTest = "root"
const DBPortTest = 8529

const ProductCollection = "product-collection"
const UserCollection = "user-collection"

func MockCollection(g *GomegaWithT, collName string) *mongo.Collection {
	db := MockClient(g).Database(DBNameTest)
	coll := db.Collection(collName)
	return coll
}

func MockClient(g *GomegaWithT) *mongo.Client {
	client, err := mongo.Connect(nil, options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%d", DBHostTest, DBPortTest)).
		SetAuth(options.Credential{
			Username: DBUserTest,
			Password: DBPassTest,
		}),
	)
	g.Expect(err).ToNot(HaveOccurred())
	return client
}
