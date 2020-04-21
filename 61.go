package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
)

type Artist struct {
	Name string `json:"name"`
	Tags []struct {
		Count int    `json:"count"`
		Value string `json:"value"`
	} `json:"tags"`
	SortName string `json:"sort_name"`
	Gid      string `json:"gid"`
	id       int    `json:"id"`
	Area     string `json:"area"`
	Aliases  []struct {
		Name     string `json:"name"`
		SortName string `json:"sort_name"`
	} `json:"aliases"`
	Begin struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Date  int `json:"date"`
	} `json:"begin"`
	End struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Date  int `json:"date"`
	} `json:"end"`
	Rating struct {
		Count int `json:"count"`
		Value int `json:"value"`
	} `json:"rating"`
}

func main() {
	filename := "artist.json"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	db, err := leveldb.OpenFile("chapter7_kvs.db", nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	reader := bufio.NewReader(file)
	for i := 0; i < 10; i++ {
		b, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		artist := Artist{}
		json.Unmarshal([]byte(b), &artist)
		data, err := db.Get([]byte(artist.Name), nil)
		if err != nil {
			panic(err)
		}
		fmt.Print(artist.Name, ":", string(data), "\n")
	}

}
