package hhapi
import "fmt"

type VacancySearchResult struct {
	VacanciesFound int `json:"found,int64"`
	SearchUrl string `json:"alternate_url"`
	Pages int `json:"pages,int64"`
	CurrentPage int `json:"page,int64"`
	Clusters []Cluster `json:"clusters"`
	Vacancies []VacancyShort `json:"items"`
}

func (vsr VacancySearchResult) String() string {
	return fmt.Sprintf("Vacancies found: %v\nSearch Url: %v\nPages: %v\nCurrent page: %v\nClusters: %v\nVacancies: %v\n",
		vsr.VacanciesFound,
		vsr.SearchUrl,
		vsr.Pages,
		vsr.CurrentPage,
		vsr.Clusters,
		vsr.Vacancies,
	)
}

type VacancyShort struct {
	Id int `json:"id,string"`
	Premium bool `json:"premium"`
	Address Address `json:"address"`
	AlternateUrl string `json:"alternate_url"`
	Salary Salary `json:"salary"`
	Name string `json:"name"`
	Area Area `json:"area"`
	Url string `json:"url"`
	PublishedAt string `json:"published_at"`
	Employer EmployerShort `json:"employer"`
	ResponseRequired bool `json:"response_letter_required"`
	Type Type `json:"type"`
}

func (vs VacancyShort) String() string {
	return fmt.Sprintf("\n--------------------------\nID: %v\nPremium: %v\nAddress: %v\nAlt. url: %v\nSalary: %v\nName: %v\nArea: %v\nUrl: %v\nPublished: %v\nEmployer: %v\nResp.required: %v\nType: %v",
		vs.Id,
		vs.Premium,
		vs.Address,
		vs.AlternateUrl,
		vs.Salary,
		vs.Name,
		vs.Area,
		vs.Url,
		vs.PublishedAt,
		vs.Employer,
		vs.ResponseRequired,
		vs.Type,
	)
}

type Vacancy struct {
	Name string `json:"name"`
	Description string `json:"description"`

}

type Cluster struct {
	// todo
}

type Address struct {
	City string `json:"city,omitempty"`
	Street string `json:"street,omitempty"`
	Building string `json:"building,omitempty"`
	Description string `json:"description,omitempty"`
	Lat float64 `json:"lat,omitempty"`
	Lng float64 `json:"street,omitempty"`
	Raw string `json:"raw,omitempty"`
	MetroStations []MetroStation `json:"metro_stations"`
}

func (a Address)String()string{
	if(a.Raw != ""){
		return fmt.Sprintf("%v", a.Raw)
	} else {
		return fmt.Sprintf("-")
	}
}

type MetroStation struct {
	// todo
}

type Salary struct {
	To int64 `json:"to,int64"`
	From int64 `json:"from,int64"`
	Currency string `json:"currency"`
}

func (c Salary) String() string {
	if c.From != 0 && c.To != 0 {
		if c.From == c.To {
			return fmt.Sprintf("%v %v", c.From, c.Currency)
		}
		return fmt.Sprintf("from %v to %v %v", c.From, c.To, c.Currency)
	}
	if c.From != 0 && c.To == 0 {
		return fmt.Sprintf("from %v %v", c.From, c.Currency)
	}

	if c.From == 0 && c.To != 0 {
		return fmt.Sprintf("up to %v %v", c.To, c.Currency)
	}


	return fmt.Sprintf("not specified")
}

type Area struct {
	Url string `json:"url"`
	Id string `json:"id"` //,int64
	Name string `json:"name"`
}

func (a Area) String() string {
	return fmt.Sprintf("%v", a.Name)
}

type EmployerShort struct {
	Name string `json:"name"`
	Type string `json:"type"`
	//todo
}

func (es EmployerShort)String() string {
	return fmt.Sprintf("%v", es.Name)
}

type Type struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

func (t Type)String()string{
	return fmt.Sprintf("%v", t.Name)
}

type SearchRequest struct {
	Text string
	SearchField VacancySearchField
}