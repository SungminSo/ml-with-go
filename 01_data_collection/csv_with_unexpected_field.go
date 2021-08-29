package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {
	// 데이터가 항상 명확하고 깔끔하게 되어있다는 보장이 없다.
	// csv 레코드에서 예상치못한 필드나 필드의 수를 발견할 수 있다.
	// iris_with_unexpected_field.csv 파일을 활용해 확인
	f, err := os.Open("data/iris_with_unexpected_field.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	// 파악된 데이터가 정상적인 경우라면 각 라인에 6개의 필드가 있어야 함
	// FieldsPerRecord를 6으로 설정하면 CSV의 각 행에 정확한 개수의 필드가 존재하는지 확인 가능
	reader.FieldsPerRecord = 6

	// rawCSVData는 성공적으로 파싱된 데이터 저장
	var rawCSVData [][]string

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			// record 읽는 중에 오류 발견 시 기록만 하고 계속 진행
			log.Println(err)
			continue
		}

		rawCSVData = append(rawCSVData, record)
	}

	// reader.ReadAll 방식으로 읽을 경우 err 발생으로 인해 rawCSVData에는 데이터가 없음
	// rawCSVData, err := reader.ReadAll()
	// if err != nil {
	// 	 log.Println(err)
	// }
	// fmt.Println(len(rawCSVData))  # 0
}
