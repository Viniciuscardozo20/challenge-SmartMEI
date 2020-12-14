package database

import (
	"testing"
	"time"

	. "challenge-SmartMEI/helper_tests"

	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
)

const testCollection = "test-collection-10"

func TestCollection(t *testing.T) {
	g := NewGomegaWithT(t)
	db, err := NewDatabase(FakeDbConfig())
	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(db).ShouldNot(BeNil())

	t.Run("validate the collection creation", func(t *testing.T) {
		coll1, err := db.Collection(testCollection)

		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(coll1).ShouldNot(BeNil())
		coll := MockCollection(g, testCollection)
		names, err := coll.Database().ListCollectionNames(nil, bson.M{})
		g.Expect(err).ShouldNot(HaveOccurred())
		exist := false
		for _, collName := range names {
			if collName == testCollection {
				exist = true
				break
			}
		}
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(exist).Should(BeTrue())
	})

	t.Run("validate the user creation", func(t *testing.T) {
		coll, err := db.Collection(testCollection)
		g.Expect(err).ShouldNot(HaveOccurred())
		user, err := coll.CreateUser(User{
			Id:            01,
			Name:          "test6",
			Email:         "test1@email",
			LentBooks:     make([]BookLoan, 0),
			BorrowedBooks: make([]BookLoan, 0),
			CreatedAt:     time.Now(),
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(user).ShouldNot(BeNil())
	})

	t.Run("validate add book in user collection", func(t *testing.T) {
		coll, err := db.Collection(testCollection)
		g.Expect(err).ShouldNot(HaveOccurred())
		book, err := coll.AddBookToMyCollection(01, Book{
			Id:        04,
			Title:     "Title-test-4",
			Pages:     "99",
			CreatedAt: time.Now(),
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(book).ShouldNot(BeNil())
	})

	t.Run("validate lend book", func(t *testing.T) {
		coll, err := db.Collection(testCollection)
		g.Expect(err).ShouldNot(HaveOccurred())
		_, err = coll.CreateUser(User{
			Id:            40,
			Name:          "test89",
			Email:         "test1@email",
			LentBooks:     make([]BookLoan, 0),
			BorrowedBooks: make([]BookLoan, 0),
			CreatedAt:     time.Now(),
		})
		_, err = coll.CreateUser(User{
			Id:            41,
			Name:          "test90",
			Email:         "test1@email",
			LentBooks:     make([]BookLoan, 0),
			BorrowedBooks: make([]BookLoan, 0),
			CreatedAt:     time.Now(),
		})
		book, err := coll.AddBookToMyCollection(40, Book{
			Id:        1,
			Title:     "Title-test",
			Pages:     "589",
			CreatedAt: time.Now(),
		})
		bookLoan, err := coll.LendBook(40, BookLoan{
			Book:     *book,
			FromUser: 40,
			ToUser:   41,
			LentAt:   time.Now(),
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(bookLoan).ShouldNot(BeNil())
	})

	t.Run("validate returned book", func(t *testing.T) {
		coll, err := db.Collection(testCollection)
		g.Expect(err).ShouldNot(HaveOccurred())
		_, err = coll.CreateUser(User{
			Id:            20,
			Name:          "test89",
			Email:         "test1@email",
			LentBooks:     make([]BookLoan, 0),
			BorrowedBooks: make([]BookLoan, 0),
			CreatedAt:     time.Now(),
		})
		_, err = coll.CreateUser(User{
			Id:            21,
			Name:          "test90",
			Email:         "test1@email",
			LentBooks:     make([]BookLoan, 0),
			BorrowedBooks: make([]BookLoan, 0),
			CreatedAt:     time.Now(),
		})
		book, err := coll.AddBookToMyCollection(20, Book{
			Id:        1,
			Title:     "Title-test",
			Pages:     "589",
			CreatedAt: time.Now(),
		})
		bookLoan, err := coll.LendBook(20, BookLoan{
			Book:     *book,
			FromUser: 20,
			ToUser:   21,
			LentAt:   time.Now(),
		})
		bookLoan, err = coll.ReturnBook(21, *&book.Id)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(bookLoan).ShouldNot(BeNil())
	})

}
