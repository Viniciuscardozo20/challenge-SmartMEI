package database

import (
	. "challenge-SmartMEI/database/models"
	. "challenge-SmartMEI/helper_tests"
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
)

const testCollection = "test-collection-30"

func TestCollection(t *testing.T) {
	g := NewGomegaWithT(t)
	client := MockClient(g)
	err := client.Database(DBNameTest).Drop(nil)
	g.Expect(err).ShouldNot(HaveOccurred())
	db, err := NewDatabase(FakeDbConfig())
	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(db).ShouldNot(BeNil())

	coll := MockCollection(g, testCollection)

	t.Run("validate the user creation", func(t *testing.T) {
		dbcoll, err := db.Collection(testCollection)
		g.Expect(err).ShouldNot(HaveOccurred())
		coll.DeleteMany(nil, bson.M{})
		user, err := dbcoll.CreateUser(User{
			Name:          "testCreate",
			Email:         "user1@email.com",
			Collection:    make([]Book, 0),
			LentBooks:     make([]BookLoan, 0),
			BorrowedBooks: make([]BookLoan, 0),
			CreatedAt:     time.Now(),
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(user).ShouldNot(BeNil())
		user2, err := dbcoll.CreateUser(User{
			Name:          "testCreate",
			Email:         "user2@email.com",
			Collection:    make([]Book, 0),
			LentBooks:     make([]BookLoan, 0),
			BorrowedBooks: make([]BookLoan, 0),
			CreatedAt:     time.Now(),
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(user2).ShouldNot(BeNil())
	})

	t.Run("validate add book in user collection", func(t *testing.T) {
		coll1, err := db.Collection(testCollection)
		g.Expect(err).ShouldNot(HaveOccurred())
		coll.DeleteMany(nil, bson.M{})
		user, err := coll1.CreateUser(User{
			Name:          "testAdd",
			Email:         "user1@email.com", //generate new email
			Collection:    make([]Book, 0),
			LentBooks:     make([]BookLoan, 0),
			BorrowedBooks: make([]BookLoan, 0),
			CreatedAt:     time.Now(),
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(user).ShouldNot(BeNil())
		book, err := coll1.AddBookToMyCollection(*user, Book{
			Title:     "Title-test-4",
			Pages:     "99",
			CreatedAt: time.Now(),
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(book).ShouldNot(BeNil())
	})

	t.Run("validate lend book", func(t *testing.T) {
		coll1, err := db.Collection(testCollection)
		g.Expect(err).ShouldNot(HaveOccurred())
		coll.DeleteMany(nil, bson.M{})
		user1, err := coll1.CreateUser(User{
			Name:          "user1",
			Email:         "user1@email.com",
			Collection:    make([]Book, 0),
			LentBooks:     make([]BookLoan, 0),
			BorrowedBooks: make([]BookLoan, 0),
			CreatedAt:     time.Now(),
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(user1).ShouldNot(BeNil())
		user2, err := coll1.CreateUser(User{
			Name:          "user2",
			Email:         "user2@email.com",
			Collection:    make([]Book, 0),
			LentBooks:     make([]BookLoan, 0),
			BorrowedBooks: make([]BookLoan, 0),
			CreatedAt:     time.Now(),
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(user2).ShouldNot(BeNil())
		book, err := coll1.AddBookToMyCollection(*user1, Book{
			Title:     "Title-test-4-lend",
			Pages:     "99",
			CreatedAt: time.Now(),
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(book).ShouldNot(BeNil())
		bookLoan, err := coll1.LendBook(*user1, *user2, BookLoan{
			Book:     *book,
			FromUser: user1.Id,
			ToUser:   user2.Id,
			LentAt:   time.Now(),
			Returned: false,
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(bookLoan).ShouldNot(BeNil())
	})

	t.Run("validate returned book", func(t *testing.T) {
		coll1, err := db.Collection(testCollection)
		g.Expect(err).ShouldNot(HaveOccurred())
		coll.DeleteMany(nil, bson.M{})
		user1, err := coll1.CreateUser(User{
			Name:          "user1",
			Email:         fmt.Sprint("user1@email.com"), //generate new email
			Collection:    make([]Book, 0),
			LentBooks:     make([]BookLoan, 0),
			BorrowedBooks: make([]BookLoan, 0),
			CreatedAt:     time.Now(),
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(user1).ShouldNot(BeNil())
		user2, err := coll1.CreateUser(User{
			Name:          "user2",
			Email:         fmt.Sprint("user2@email.com"), //generate new email
			Collection:    make([]Book, 0),
			LentBooks:     make([]BookLoan, 0),
			BorrowedBooks: make([]BookLoan, 0),
			CreatedAt:     time.Now(),
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(user2).ShouldNot(BeNil())
		book, err := coll1.AddBookToMyCollection(*user1, Book{
			Title:     "Title-test-4-lend",
			Pages:     "99",
			CreatedAt: time.Now(),
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(book).ShouldNot(BeNil())
		bookLoan, err := coll1.LendBook(*user1, *user2, BookLoan{
			Book:     *book,
			FromUser: user1.Id,
			ToUser:   user2.Id,
			LentAt:   time.Now(),
			Returned: false,
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(bookLoan).ShouldNot(BeNil())
		user1, err = coll1.GetUserDetails(user1.Id)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(user1).ShouldNot(BeNil())
		user1.LentBooks[0].Returned = true
		user2, err = coll1.GetUserDetails(user2.Id)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(user2).ShouldNot(BeNil())
		user2.BorrowedBooks[0].Returned = true
		err = coll1.ReturnBook(*user2, *user1)
		g.Expect(err).ShouldNot(HaveOccurred())
	})

}
