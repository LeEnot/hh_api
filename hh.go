package main
import (
	"fmt"
	"net/http"
	//"bytes"
	"io/ioutil"
	"./hhapi"
	"encoding/json"
	"log"
	"strconv"
)

func main(){

	url := "https://api.hh.ru/vacancies" //?text=android&area=1
	fmt.Println("URL:>", url)

	//var query = []byte(`?text=android&area=1`) //bytes.NewBuffer(query)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "text/plain")
	values := req.URL.Query()
	values.Add("text", "android")
	values.Add("search_field", "name")
	values.Add("area", "1")
	values.Add("page", "0")
	req.URL.RawQuery = values.Encode()
	var pagesMax int = 1
	var currentPage int = 0
	var vacancies hhapi.VacancySearchResult
	var allVacancies []hhapi.VacancyShort
	client := &http.Client{}

	//_ = &hhapi.SearchRequest{Text:"ololo", SearchField:hhapi.Description,}

	for currentPage < pagesMax {

		values := req.URL.Query()
		values.Set("page", strconv.Itoa(currentPage))
		req.URL.RawQuery = values.Encode()


		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		//fmt.Println("response Status:", resp.Status)
		//fmt.Println("response Headers:", resp.Header)
		//fmt.Println()
		body, _ := ioutil.ReadAll(resp.Body)


		if err := json.Unmarshal(body, &vacancies); err != nil {
			log.Fatal(err)
		}
		pagesMax = vacancies.Pages
		currentPage++
		fmt.Println("Page ", currentPage - 1, " len: " ,len(vacancies.Vacancies))
		allVacancies = append(allVacancies, vacancies.Vacancies...)
	}
	fmt.Println(allVacancies)
	//fmt.Println("response Body:", string(body))

}
