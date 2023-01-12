package sqlite_bookrepository_test

import (
	"testing"

	bookrepository "github.com/kyzykyky/softwarearch/bookservice/internal/data/bookrepository/sqlite"
	"github.com/kyzykyky/softwarearch/bookservice/internal/data/config/sqlite"

	"github.com/kyzykyky/softwarearch/bookservice/pkg/domain"
)

func TestSqlite(t *testing.T) {
	conf := sqlite.Config{Path: "test.sqlite"}
	repo, err := conf.Connect()
	if err != nil {
		t.Fatal(err)
	}

	bookrepo, err := bookrepository.Config{
		DbConnection: repo.DbConnection,
	}.Init()
	if err != nil {
		t.Fatal(err)
	}

	testIsbn := "12345-test-6789"
	_, err = bookrepo.AddBook(domain.Book{
		Isbn:   testIsbn,
		Title:  "First book",
		Author: "John Doe",
		Price:  105,
	})
	if err != nil {
		t.Fatal(err)
	}

	book, err := bookrepo.GetBook(testIsbn)
	if err != nil {
		t.Fatal(err)
	} else if book.Isbn != testIsbn {
		t.Fatal("ISBN does not match")
	}

	book, err = bookrepo.UpdateBook(testIsbn, domain.Book{
		Isbn:   testIsbn,
		Title:  "First book",
		Author: "John Doe",
		Price:  110,
	})
	if err != nil {
		t.Fatal(err)
	} else if book.Price != 110 {
		t.Fatal("Price does not match")
	}

	err = bookrepo.DeleteBook(testIsbn)
	if err != nil {
		t.Fatal(err)
	}
}
