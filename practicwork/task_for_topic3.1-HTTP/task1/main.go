package main

// Изучить документацию для API сервиса coincap.io (https://pro.coincap.io/api-docs)
// Создайте аккуаунт, создайте API ключ и на ЯП Go сделайте GET запрос: /assets/. Вывод включает много полей. Выведите на экран ограниченным списком:
// [{
//     name: Bitcoin,
//     priceUsd: 57435.2862859343831301,
// },
// {
//     name: Ethereum,
//     priceUsd: 3686.1846963999934439,

// }
// ...
// ]

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Resp struct {
	Name     string `json:"name"`
	PriceUsd string `json:"priceUsd"`
}

var Assets struct {
	Data []Resp `json:"data"`
}

func main() {
	resp, err := http.Get("https://rest.coincap.io/v3/assets?apiKey={my_tocken}")
	if err != nil {
		log.Println(err)
	}
	if err = json.NewDecoder(resp.Body).Decode(&Assets); err != nil {
		log.Println(err)
	}
	var out []Resp
	for _, v := range Assets.Data {
		out = append(out, Resp{Name: v.Name, PriceUsd: v.PriceUsd})
	}
	dataByte, err := json.MarshalIndent(out, "", "    ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(dataByte))
}
