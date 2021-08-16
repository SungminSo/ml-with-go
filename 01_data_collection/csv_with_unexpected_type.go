package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// CSVRecord 필드 타입 확인을 위한 구조체 정의
type CSVRecord struct {
	Id            int64
	SepalLengthCm float64
	SepalWidthCm  float64
	PetalLengthCm float64
	PetalWidthCm  float64
	Species       string
	ParseError    error
}

func main() {
	// Go는 정적으로 타입을 지정하기 때문에 csv 필드에 대해 엄격하게 검사를 할 수 있다.
	f, err := os.Open("../data/iris_with_unexpected_type.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	// rawCSVData는 성공적으로 파싱된 데이터 저장
	var csvData []CSVRecord

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// 열을 저장하기 위한 CSVRecord 생성
		var csvRecord CSVRecord

		// 기대하는 타입을 기반으로 레코드의 각 값 읽기
		for idx, value := range record {
			// value 타입이 float이여야 할 경우
			var floatValue float64
			if idx != 0 && idx != 5 {
				if floatValue, err = strconv.ParseFloat(value, 64); err != nil {
					log.Printf("Unexpected type in column %d\n", idx)
					csvRecord.ParseError = fmt.Errorf("could not parse float")
					break
				}
			}

			switch idx {
			case 0:
				var intValue int64
				if intValue, err = strconv.ParseInt(value, 10, 64); err != nil {
					log.Println("Unexpected type in column 0")
					csvRecord.ParseError = fmt.Errorf("value of id is not int")
				}
				csvRecord.Id = intValue
			case 1:
				csvRecord.SepalLengthCm = floatValue
			case 2:
				csvRecord.SepalWidthCm = floatValue
			case 3:
				csvRecord.PetalLengthCm = floatValue
			case 4:
				csvRecord.PetalWidthCm = floatValue
			case 5:
				// 빈 문자열이 아닌지 체크
				if value == "" {
					log.Println("Unexpected type in column 5")
					csvRecord.ParseError = fmt.Errorf("empty string value")
				}
				csvRecord.Species = value
			}
		}

		if csvRecord.ParseError == nil {
			csvData = append(csvData, csvRecord)
		}
	}
}
