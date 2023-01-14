package gorm_bookrepository_test

import (
	"context"
	"testing"
	"time"

	bookrepository "github.com/kyzykyky/softwarearch/bookservice/internal/data/bookrepository/gorm"
	"github.com/kyzykyky/softwarearch/bookservice/internal/data/config/gorm/sqlite"

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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	testIsbn := "12345-test-6789"
	book, err := bookrepo.AddBook(ctx, domain.Book{
		Isbn:   testIsbn,
		Title:  "First book",
		Author: "John Doe",
		Price:  105,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Book ID: ", book.Id)

	book, err = bookrepo.GetBook(ctx, book.Id)
	if err != nil {
		t.Fatal(err)
	} else if book.Isbn != testIsbn {
		t.Fatal("ISBN does not match, expected 12345-test-6789, got ", book.Isbn)
	}

	book, err = bookrepo.UpdateBook(ctx, domain.Book{
		Id:     book.Id,
		Isbn:   testIsbn,
		Title:  "First book",
		Author: "John Doe",
		Price:  110,
	})
	if err != nil {
		t.Fatal(err)
	} else if book.Price != 110 {
		t.Fatalf("Price does not match, expected 110, got %f", book.Price)
	}

	err = bookrepo.DeleteBook(ctx, book.Id)
	if err != nil {
		t.Fatal(err)
	}
}
