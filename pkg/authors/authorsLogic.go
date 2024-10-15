package authors

import (
	"errors"
	"strconv"
	"tz2/pkg"
)

// Функция сравнивает id и максимальное id автора в таблице
func SelectAuthorMaxId(idCheck string) error {
	authors, err := pkg.SelectAllAuthors("SELECT * FROM author")
	if err != nil {
		return err
	}

	var maxAuthorsId int
	for _, v := range authors {
		if v.Id > maxAuthorsId {
			maxAuthorsId = v.Id
		}
	}
	idInt, err := strconv.Atoi(idCheck)
	if err != nil {
		return err
	}

	if idInt > maxAuthorsId {
		return errors.New("max authors id out of range")
	}
	return nil
}
