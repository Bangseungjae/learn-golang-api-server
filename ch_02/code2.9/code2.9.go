package code2_9

import (
	"context"
	"fmt"
)

// context 미적용 함수. 오류 처리는 생략
func GetCompanyUsecase(userID UserID) (*Company, error) {
	// context 적용 완료 함수
	u, _ := GetUser(context.TODO(), userID)
	fmt.Println(u)
	// context 미적용 함수
	c, _ := GetCompanyByUserID(u)
	return c, nil
}

func GetUser(ctx context.Context, id UserID) (*User, error) {
	// 어떤 처리
	return &User{}, nil
}

func GetCompanyByUserID(u *User) (*Company, error) {
	// 어떤 처리
	return &Company{}, nil
}

type UserID struct{}
type User struct{}

type Company struct{}
