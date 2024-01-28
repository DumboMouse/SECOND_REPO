package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Ваша структура данных, соответствующая формату файла JSON
type Data []struct {
	GlobalID int `json:"global_id"`
}

func main() {
	url := "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	bodybytes, err := io.ReadAll(response.Body)

	var yourDataStruct Data

	err = json.Unmarshal(bodybytes, &yourDataStruct)
	if err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
		return
	}

	var sum int

	for _, element := range yourDataStruct {
		el := element.GlobalID
		sum += el
	}

	fmt.Println(sum)
}
