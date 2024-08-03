package main

import "sort"

type bookCollection map[Book]struct{}

type bookReccomendations map[Book]bookCollection

func newCollection() bookCollection {
	return make(bookCollection)
}

func recommendOtherBooks(bookworms []Bookworm) []Bookworm{
	sb := make(bookReccomendations)

	for _, bookworm :=range bookworms{
		for i, book := range bookworm.Books {
			otherBooksOnShelves := listOtherBooksOnShelves(i, bookworm.Books)
			registerBookRecommendations(sb, book, otherBooksOnShelves)
		}
	}

	recommendations := make([]Bookworm, len(bookworms))
	for i, bookworm := range bookworms {
		recommendations[i]= Bookworm{
			Name: bookworm.Name,
			Books: bookworm.Books,
		}
	}

	return recommendations
}

func listOtherBooksOnShelves(bookIndexToRemove int, myBooks []Book) []Book {
	otherBooksOnShelves := make([]Book, bookIndexToRemove, len(myBooks)-1)
	copy(otherBooksOnShelves, myBooks[:bookIndexToRemove])
	otherBooksOnShelves = append(otherBooksOnShelves, myBooks[bookIndexToRemove +1:]...)

	return otherBooksOnShelves
}

func registerBookRecommendations(recommendations bookReccomendations, reference Book, otherBooksOnShelves []Book) {
	for _, book := range otherBooksOnShelves {
		collection, ok := recommendations[reference]
		if !ok {
			collection = newCollection()
			recommendations[reference] = collection
		}

		collection[book] = struct{}{}
	}
}

func bookCollectionToListOfBooks(bc bookCollection) []Book {
	bookList := make([]Book, 0, len(bc))
	for book := range bc {
		bookList = append(bookList, book)
	}

	sort.Slice(bookList, func(i, j int) bool {
		if bookList[i].Author != bookList[j].Author {
			return bookList[i].Author < bookList[j].Author
		}

		return bookList[i].Title < bookList[j].Title
	})

	return bookList
}

func recommendBooks(recommendations bookReccomendations, myBooks []Book) []Book {
	bc := make(bookCollection)

	myShelf := make(map[Book]bool)
	for _, myBook := range myBooks {
		myShelf[myBook] = true
	}

	for _, myBook := range myBooks {
		for recommendation := range recommendations[myBook] {
			if myShelf[recommendation] {
				continue
			}

			bc[recommendation] = struct{}{}
		}
	}

	recommendationsForBook := bookCollectionToListOfBooks(bc)

	return recommendationsForBook
}
