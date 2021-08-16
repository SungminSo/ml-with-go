package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// csv 오픈
	f, err := os.Open("../data/blah_example.csv")
	if err != nil {
		log.Fatal(err)
	}

	// csv 읽기
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 정수열에서 최댓값 구하기
	var intMax int
	for _, record := range records {
		// 정수값 해석
		intVal, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}

		// 적당한 경우 최댓값 변경
		if intVal > intMax {
			intMax = intVal
		}
	}

	// 최댓값 출력
	fmt.Println(intMax)

	// blah_example.csv
	// 1,blah1
	// 2,blah2
	// ,blah3

	// 출력값
	// strconv.Atoi: parsing "": invalid syntax
	// exit status 1
	// => 파이썬과는 달리 별다른 처리없이 데이터 무결성 유지 가능

	// 이와 같은 기능을 하는 파이썬 코드 예제
	// import pandas as pd
	//
	// cols = [
	// 	'integer_column',
	// 	'string_column'
	// ]
	//
	// data = pd.read_csv('../data/blah_example.csv', names=cols)
	// print(data['integer_column'].max())

	// 파이썬 예제 코드의 출력값
	// 2.0
	// => 누락된 데이터 여부 판단 x
	// ==> 편의성은 압도적으로 높지만 얼핏 잘못 사용하면 잘못된 결과가 나올 가능성이 높다
}

