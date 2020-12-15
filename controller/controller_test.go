package controller

import (
	. "challenge-SmartMEI/controller/dto"
	. "challenge-SmartMEI/database"
	. "challenge-SmartMEI/database/models"
	. "challenge-SmartMEI/helper_tests"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

const testCollection = "test-controller-1"

func TestController(t *testing.T) {
	g := NewGomegaWithT(t)
	db, err := NewDatabase(FakeDbConfig())
	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(db).ShouldNot(BeNil())
	coll, err := db.Collection(testCollection)
	g.Expect(err).ShouldNot(HaveOccurred())
	t.Run("validate the user creation", func(t *testing.T) {
		user, err := coll.CreateUser(User{
			Name:          "test6",
			Email:         "test1@email",
			LentBooks:     make([]BookLoan, 0),
			BorrowedBooks: make([]BookLoan, 0),
			CreatedAt:     time.Now(),
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(user).ShouldNot(BeNil())
	})
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

func fakeUser() CreateUserInput {
	return CreateUserInput{
		Name:  "Ravi",
		Email: "Ravi@gmail.com",
	}
}

func fakeBook() AddBookInput {
	return AddBookInput{
		Title: "Lord of Rings",
		Pages: "789",
	}
}
