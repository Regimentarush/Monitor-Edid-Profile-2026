package main

import (
	"fmt"
	"sort"
)

type Book struct {
	Title  string
	Author string
	Pages  int
	Rating float64
}

type Library struct {
	Books []Book
}

func (l *Library) AddBook(title, author string, pages int, rating float64) {
	l.Books = append(l.Books, Book{
		Title:  title,
		Author: author,
		Pages:  pages,
		Rating: rating,
	})
}

func (l *Library) SortByRating() {
	sort.Slice(l.Books, func(i, j int) bool {
		return l.Books[i].Rating > l.Books[j].Rating
	})
}

func (l *Library) TotalPages() int {
	total := 0

	for _, book := range l.Books {
		total += book.Pages
	}

	return total
}

func (l *Library) AverageRating() float64 {
	if len(l.Books) == 0 {
		return 0
	}

	total := 0.0

	for _, book := range l.Books {
		total += book.Rating
	}

	return total / float64(len(l.Books))
}

func (l *Library) PrintReport() {
	fmt.Println("Library Report")
	fmt.Println("==============")

	for _, book := range l.Books {
		fmt.Printf(
			"%s | %s | %d pages | Rating: %.1f\n",
			book.Title,
			book.Author,
			book.Pages,
			book.Rating,
		)
	}

	fmt.Println("==============")
	fmt.Printf("Books: %d\n", len(l.Books))
	fmt.Printf("Total Pages: %d\n", l.TotalPages())
	fmt.Printf("Average Rating: %.2f\n", l.AverageRating())
}

func main() {
	library := Library{}

	library.AddBook("The Silent Ocean", "James Carter", 320, 8.7)
	library.AddBook("Digital Dreams", "Emily Stone", 280, 9.1)
	library.AddBook("Beyond Tomorrow", "Daniel Brooks", 410, 8.5)
	library.AddBook("The Last Journey", "Michael Reed", 365, 9.3)

	library.SortByRating()
	library.PrintReport()
}