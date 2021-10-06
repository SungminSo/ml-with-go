package main

import (
	"database/sql"
	"fmt"

	// pq는 databases/sql를 활용해 postgres에 연결하는 기능을 제공하는 라이브러리
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	// PGURL : postgres url. format: "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	pgURL := os.Getenv("PGURL")
	if pgURL == "" {
		log.Fatal("PGURL empty")
	}

	// DB open
	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // 여기 주의!
	// 여기서 db를 만든것이 DB 연결에 성공한 것은 아니다.

	// DB 연결 확인
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	////////////////////////////////////////////////////////////////////////////////
	// 데이터 쿼리
	// 붓꽃 데이터 DB가 있다고 가정
	rows, err := db.Query(`
		SELECT
			sepal_length as sLength,
		    sepal_width as sWidth,
		    petal_length as pLength,
		    petal_width as pWidth
		FROM iris
		WHERE species = $1`, "Iris-setosa")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() // 여기도 주의!

	// loop로 행을 읽어서 데이터 출력
	for rows.Next() {
		var (
			sLength float64
			sWidth  float64
			pLength float64
			pWidth  float64
		)

		if err = rows.Scan(&sLength, &sWidth, &pLength, &pWidth); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%.2f, %.2f, %.2f, %.2f\n", sLength, sWidth, pLength, pWidth)
	}
	// loop 작업이 완료되면 오류 여부 확인
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	////////////////////////////////////////////////////////////////////////////////
	// 일부 값 update
	res, err := db.Exec("UPDATE iris SET species = 'setosa' WHERE species = 'Iris-setosa'")
	if err != nil {
		log.Fatal(err)
	}

	// update 여부 확인 - 해당 UPDATE문에 의해 영향받은 행의 수 확인
	rowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("affected = %d\n", rowCount)
}
