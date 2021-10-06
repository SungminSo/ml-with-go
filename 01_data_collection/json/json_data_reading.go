package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// citiBikeURL : CitiBike 자전거 공유 정류장의 상황을 알려준다.
const citiBikeURL = "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"

// statusData : citiBikeURL로부터 반환된 JSON 문서의 구문을 분석하는데 사용
type statusData struct {
	LastUpdated int `json:"last_updated"`
	TTL         int `json:"ttl"`
	Data        struct {
		Stations []station `json:"stations"`
	} `json:"data"`
}

// station : stationData 안의 각 station 문서의 구문을 분석하는데 사용
type station struct {
	ID                string `json:"station_id"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumBikesDisabled  int    `json:"num_bikes_disabled"`
	NumDocksAvailable int    `json:"num_docks_available"`
	NumDocksDisabled  int    `json:"num_docks_disabled"`
	IsInstalled       int    `json:"is_installed"`
	IsRenting         int    `json:"is_renting"`
	IsReturning       int    `json:"is_returning"`
	LastReported      int    `json:"last_reported"`
	HasAvailableKeys  bool   `json:"has_available_keys"`
}

func main() {
	// URL로부터 JSON 응답을 얻는다.
	response, err := http.Get(citiBikeURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// 응답의 Body를 []byte로 읽는다.
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// stationData 타입의 변수 선언
	var sd statusData
	if err := json.Unmarshal(body, &sd); err != nil {
		log.Fatal(err)
	}

	// 첫번째 정류장 정보를 출력
	fmt.Printf("%+v\n\n", sd.Data.Stations[0])

	// JSON 출력
	// 데이터를 다시 marshal
	outputData, err := json.Marshal(sd)
	if err != nil {
		log.Fatal(err)
	}

	// JSON 형식으로 생성된 데이터를 파일에 저장
	if err := ioutil.WriteFile("citibike.json", outputData, 0644); err != nil {
		log.Fatal(err)
	}
}
