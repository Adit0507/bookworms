package main

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
)

func TestLoadBookworms_Success(t *testing.T) {
	tests := map[string]struct {
		bookwormsFile string
		want          []Bookworm
		wantErr       bool
	}{
		"file exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
				},
				wantErr: false,				
		},
		"file doesn't exist": {...},
		"invalid JSON": {...},
	}

	for name, testCase := range test {
		t.Run(name, func(t *testing.T){
			got, err := loadBookworms(testCase.bookwormsFile)
			if err != nil && !testCase.wantErr { 
				t.Fatalf("expected an error %s, got none", err.Error())
				}
				if err == nil && testCase.wantErr { 
				t.Fatalf("expected no error, got one %s", err.Error())
				}
				if !equalBookworms(got, testCase.want) {
				t.Fatalf("different result: got %v, expected %v", got, testCase.want)
				}	
		})
	}


}

func equalBookworms(bookworms, target []Bookworm) bool {
	if len(bookworms)!= len(target) {
		return false
	}

	for i := range bookworms {
		if bookworms[i].Name != target[i].Name{
			return false
		}

		if !equalBooks(bookworms[i].Books, target[i].Books){
			return false
		}
	}

	return true
}

func equalBooks(books, target []Book) bool{
	if len(books)!= len(target){
		return false
	}

	for i := range books {
		if books[i] != target[i] { 
			return false
		}
	}

	return true
}

func equalBooksCount(got, want map[Book] uint) bool {
	if len(got) != len(want){
		return false
	}

	for book, targetCount := range want{
		count, ok := got[book]
		if !ok || targetCount != count {
			return false
		}
	}

	return true
}

func TestBooksCount(t *testing.T) {
	tt := map[string]struct {
	input []Bookworm
	want map[Book]uint
	}{
	"nominal use case": {
	input: []Bookworm{
	{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
	{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
	},
	want: map[Book]uint{handmaidsTale: 2, theBellJar: 1, 
	},
	"no bookworms": {
		input: []Bookworm{},
		want: map[Book]uint{},
	},
	"bookworm without books": {...},
		"bookworm with twice the same book": {...},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
		got := booksCount(tc.input)
		if !equalBooksCount(t, tc.want, got) { 
		t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
		}
	})
	}
}

func TestFindCommonBooks(t *testing.T) {
	tt := map[string]struct {
		input []Bookworm
		want []Book
	}{
		"no common book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre}},
			},
			want: nil,
			},
			"one common book": {...},
			"three bookworms have the same books on their shelves": {...},
		}
			for name, tc := range tt {
				t.Run(name, func(t *testing.T) {
				got := findCommonBooks(tc.input)
			if !equalBooks(t, tc.want, got) { 
			t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
			})
		}
}