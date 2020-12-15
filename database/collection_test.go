package database

import (
	. "challenge-SmartMEI/database/models"
	. "challenge-SmartMEI/helper_tests"
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
)

const testCollection = "test-collection-23"

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
		rand := rand.Intn(80000-1+1) + 1
		user, err := coll.CreateUser(User{
			Name:          "testCreate",
			Email:         fmt.Sprint("test0@email.com", rand),
			Collection:    make([]Book, 0),
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
		rand := rand.Intn(80000-1+1) + 1
		user, err := coll.CreateUser(User{
			Name:          "testAdd",
			Email:         fmt.Sprint("test1@email.com", rand), //generate new email
			Collection:    make([]Book, 0),
			LentBooks:     make([]BookLoan, 0),
			BorrowedBooks: make([]BookLoan, 0),
			CreatedAt:     time.Now(),
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(user).ShouldNot(BeNil())
		book, err := coll.AddBookToMyCollection(*user, Book{
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
		user1, err := coll.CreateUser(User{
			Name:          "testAdd",
			Email:         fmt.Sprint("test1@email.com", rand.Intn(80000-1)+1), //generate new email
			Collection:    make([]Book, 0),
			LentBooks:     make([]BookLoan, 0),
			BorrowedBooks: make([]BookLoan, 0),
			CreatedAt:     time.Now(),
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(user1).ShouldNot(BeNil())
		user2, err := coll.CreateUser(User{
			Name:          "testAdd",
			Email:         fmt.Sprint("test1@email.com", rand.Intn(80000-1)+1), //generate new email
			Collection:    make([]Book, 0),
			LentBooks:     make([]BookLoan, 0),
			BorrowedBooks: make([]BookLoan, 0),
			CreatedAt:     time.Now(),
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(user2).ShouldNot(BeNil())
		book, err := coll.AddBookToMyCollection(*user1, Book{
			Title:     "Title-test-4-lend",
			Pages:     "99",
			CreatedAt: time.Now(),
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(book).ShouldNot(BeNil())
		bookLoan, err := coll.LendBook(*user1, *user2, BookLoan{
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
		coll, err := db.Collection(testCollection)
		g.Expect(err).ShouldNot(HaveOccurred())
		user1, err := coll.CreateUser(User{
			Name:          "testAdd",
			Email:         fmt.Sprint("test1@email.com", rand.Intn(80000-1)+1), //generate new email
			Collection:    make([]Book, 0),
			LentBooks:     make([]BookLoan, 0),
			BorrowedBooks: make([]BookLoan, 0),
			CreatedAt:     time.Now(),
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(user1).ShouldNot(BeNil())
		user2, err := coll.CreateUser(User{
			Name:          "testAdd",
			Email:         fmt.Sprint("test1@email.com", rand.Intn(80000-1)+1), //generate new email
			Collection:    make([]Book, 0),
			LentBooks:     make([]BookLoan, 0),
			BorrowedBooks: make([]BookLoan, 0),
			CreatedAt:     time.Now(),
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(user2).ShouldNot(BeNil())
		book, err := coll.AddBookToMyCollection(*user1, Book{
			Title:     "Title-test-4-lend",
			Pages:     "99",
			CreatedAt: time.Now(),
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(book).ShouldNot(BeNil())
		bookLoan, err := coll.LendBook(*user1, *user2, BookLoan{
			Book:     *book,
			FromUser: user1.Id,
			ToUser:   user2.Id,
			LentAt:   time.Now(),
			Returned: false,
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(bookLoan).ShouldNot(BeNil())
		user1, err = coll.GetUserDetails(user1.Id)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(user1).ShouldNot(BeNil())
		user1.LentBooks[0].Returned = true
		user2, err = coll.GetUserDetails(user2.Id)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(user2).ShouldNot(BeNil())
		user2.BorrowedBooks[0].Returned = true
		err = coll.ReturnBook(*user2, *user1)
		g.Expect(err).ShouldNot(HaveOccurred())
	})

}
