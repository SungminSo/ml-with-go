package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	// BoltDB는 로컬에 키-값 저장소를 사용해야 할 때 많이 사용되는 라이브러리
	// https://github.com/boltdb/bolt

	// 현재 디렉토리에서 embedded.db 데이터 파일을 연다.
	// 파일이 존재하지 않는 경우에는 파일 생성
	db, err := bolt.Open("embedded.db", 0660, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 데이터를 저장하기 위해 boltdb 파일에 "bucket" 생성
	if err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("myBucket"))
		if err != nil {
			return fmt.Errorf("error occur when create bucket: %s", err)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// BoltDB 파일에 key-value 데이터 넣기
	if err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("myBucket"))
		err := b.Put([]byte("keyExample"), []byte("valueExample"))
		return err
	}); err != nil {
		log.Fatal(err)
	}

	// BoltDB 파일에 저장된 key-value 데이터를 읽어서 출력
	if err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("myBucket"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key: %s, value: %s", k, v)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
