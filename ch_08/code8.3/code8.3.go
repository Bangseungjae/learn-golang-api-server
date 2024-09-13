package code8_3

import (
	"errors"
	"fmt"
)

func GetAuthor(id AuthorID) (*Author, error) {
	if !id.Valid() {
		return nil, errors.New("GetAuthor: AuthorID is invalid")
		// 오류문에 id값을 사용할 때는 Errorf 함수를 사용한다.
		//return nil, fmt.Errorf("GetAuthor: AuthorID(%d) is invalid", AuthorID)
	}
	return nil, nil
}

type Book struct {
	AuthorID AuthorID
}

func GetAuthorName(b *Book) (string, error) {
	a, err := GetAuthor(b.AuthorID)

	if err != nil {
		return "", fmt.Errorf("GetAuthorName: %v", err)
	}
	return a.Name(), nil
}

type AuthorID int
type Author struct {
	name string
}

func (a Author) Name() string {
	return a.name
}

func (id AuthorID) Valid() bool {
	return true
}
