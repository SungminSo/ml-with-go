package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
)

func main() {
	// csv 오픈
	irisFile, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// csv 파일로부터 데이터프레임 생섬
	// 열의 유형은 추론됨
	irisDF := dataframe.ReadCSV(irisFile)
	// stdout으로 출력 시 gota 패키지가 적절한 형태로 출력될 수 있도록 데이터프레임의 형식을 지정
	fmt.Println(irisDF)

	// 데이터 프레임의 필터 생성
	filter := dataframe.F{
		Colname: "Species",
		Comparator: "==",
		Comparando: "Iris-versicolor",
	}

	// 붓꽃(iris) 품종이 "Iris-versicolor"인 행만 볼 수 있도록
	// 데이터프레임 필터링
	versicolorDF := irisDF.Filter(filter)
	if versicolorDF.Err != nil {
		log.Fatal(versicolorDF.Err)
	}
	fmt.Println(versicolorDF)

	// 데이터프레임 다시 필터링
	// 이번에는 SepalWidthCm, Species 열만 선택
	versicolorDF = irisDF.Filter(filter).Select([]string{"SepalWidthCm", "Species"})
	fmt.Println(versicolorDF)

	// 데이터프레임 다시 필터링
	// 이번에는 처음 3개의 row만 선택
	versicolorDF = irisDF.Filter(filter).Select([]string{"SepalWidthCm", "Species"}).Subset([]int{0, 1, 2})
	fmt.Println(versicolorDF)
}