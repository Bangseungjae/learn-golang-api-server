package main

import (
	"context"
	"database/sql"
	"errors"
	"go-web-application-dev/ch_02/code2.8/auth"
)

type MyDB struct {
	db *sql.DB
}

type Result any // 예시를 위해

// ExecContext는 *sql.DB와 같은 명칭의 메서드를 context 객체를 통해
// 인증 기능으로 래핑(wrapping)한 메서드
func (mydb *MyDB) ExecContext(ctx context.Context, query string, args ...any) (Result, error) {
	if !auth.Writeable(ctx) {
		return errors.New("작성 권한이 없는 사용자가 실행했다."), nil
	}
	return mydb.db.ExecContext(ctx, query, args...)
}

func main() {

}
