package pkg

type Book struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	AuthorId      int    `json:"author_id"`
	PublishedDate string `json:"published_date"`
}

type Author struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	BirthYear string `json:"birth_year"`
}
