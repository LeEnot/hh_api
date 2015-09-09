package hhapi
import "net/http"

const hhVacanciesUrl = "https://api.hh.ru/vacancies"

func Search() *http.Request {
	req, _ := http.NewRequest("GET", hhVacanciesUrl, nil)
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "text/plain")
	values := req.URL.Query()
	values.Add("text", "android")
	values.Add("search_field", "name")
	values.Add("area", "1")
	values.Add("page", "0")
	req.URL.RawQuery = values.Encode()

	return req
}
