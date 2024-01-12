package internal

import "fmt"

type Publisher struct {
	Name     string
	Founders []string
}

func (p Publisher) TotalFounders() int {
	return len(p.Founders)
}

type Book struct {
	Title  string
	Author string
	Pages  int
	Publisher
}

func (b Book) ShowInfo() string {
	return fmt.Sprintf(
		"%s was written by %s and has %d pages- Total publisher founders: %d",
		b.Title, b.Author, b.Pages, b.TotalFounders(),
	)
}

type Movie struct {
	Title    string
	Director string
	IMB      float64
	Publisher
}

func (m Movie) ShowInfo() string {
	return fmt.Sprintf(
		"%s was directed by %s and has a rating of %.1f in IMB - Total publisher founders: %d",
		m.Title, m.Director, m.IMB, m.TotalFounders(),
	)
}
