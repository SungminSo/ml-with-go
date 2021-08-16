package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// iris 데이터셋 오픈
	f, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 열린 파일 읽어오는 새 CSV reader 생성
	reader := csv.NewReader(f)

	// 각 라인의 필드수를 모른다고 가정
	// FieldsPerRecord를 음수로 설정해 각 행의 필드의 수 확인 가능
	reader.FieldsPerRecord = -1

	// 모든 CSV record 읽기
	rawCSVData, err := reader.ReadAll()  // 여기서 rawCSVData 타입은 [][]string
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(rawCSVData))

	////////////////////////////////////////////////////////////
	// 여기는 ReadAll이 아닌 무한루프로 하나씩 레코드를 읽어오는 경우
	var _rawCSVData [][]string
	for {
		// record를 읽을 때, 파일 종료 지점에 도달했는지 확인
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// 데이터 집합에 레코드 추가
		_rawCSVData = append(_rawCSVData, record)
	}
}
