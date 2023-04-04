package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

/*
	Perintah Challenge:

	Buatlah sebuah service untuk melakukan POST data dalam format JSON dan secara acak setiap 15 detik dengan angka random antara 1-100 untuk value water dan wind. Gunakan url post berikut untuk menjalankan simulasi post request :
	"https://jsonplaceholder.typicode.com/posts"

	ketentuan :
	jika water dibawah 5 maka status aman
	jika water antara 6 - 8 maka status siaga
	jika water diatas 8 maka status bahaya

	jika wind dibawah 6 maka status aman
	jika wind antara 7 - 15 maka status siaga
	jika wind diatas 15 maka status bahaya

	value water dalam satuan meter
	value wind dalam satuan meter per detik
*/

func main() {
	for true {
		post()
		time.Sleep(time.Second * 15) // setiap 15 detik
	}
}

func post() {
	data := map[string]interface{}{
		"water": rand.Intn(100) + 1, // acak 1-100
		"wind":  rand.Intn(100) + 1, // acak 1-100
	}

	requestJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body)) // ketambahan id dari "backend"

	var responseMap map[string]interface{}
	json.Unmarshal(body, &responseMap) // unmarshal response biar bisa diolah/dibaca go
	// delete(responseMap, "id")
	// newResponse, _ := json.Marshal(responseMap)
	// fmt.Println(string(newResponse))
	// kalo mau ngilangin prop "id", uncomment comment diatas

	statusWater(responseMap["water"].(float64))
	statusWind(responseMap["wind"].(float64))
}

func statusWater(water float64) {
	// float64 soalnya dari json.Unmarshal balikannya float64
	// kalau dibuat int, bakal `panic: interface conversion: interface {} is float64, not int`
	switch {
	case water < 5:
		fmt.Println("status water: aman")
	case water >= 6 && water <= 8:
		fmt.Println("status water: siaga")
	case water > 8:
		fmt.Println("status water: bahaya")
	}
}

func statusWind(wind float64) {
	switch {
	case wind < 6:
		fmt.Println("status wind: aman")
	case wind >= 7 && wind <= 15:
		fmt.Println("status wind: siaga")
	case wind > 15:
		fmt.Println("status wind: bahaya")
	}
}
