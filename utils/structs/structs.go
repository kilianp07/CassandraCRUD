package structs

type Coordinates struct {
	Type        string     `json:"type"`
	Coordinates [2]float64 `json:"coordinates"`
}

type Address struct {
	Id           string      `json:"id"`
	Building     string      `json:"building"`
	Coord        Coordinates `json:"coord"`
	Street       string      `json:"street"`
	Zipcode      string      `json:"zipcode"`
	RestaurantId string      `json:"restaurant_id"`
}

type Grade struct {
	Date struct {
		Date int64 `json:"$date"`
	} `json:"date"`
	Grade string `json:"grade"`
	Score int    `json:"score"`
}

type Restaurant struct {
	Address Address `json:"address"`
	Borough string  `json:"borough"`
	Cuisine string  `json:"cuisine"`
	Grades  []Grade `json:"grades"`
	Name    string  `json:"name"`
	Id      string  `json:"id"`
}

type RestaurantRequest struct {
	Borough string `json:"borough"`
	Cuisine string `json:"cuisine"`
	Name    string `json:"name"`
}

type AddressRequest struct {
	Building     string `json:"building"`
	Street       string `json:"street"`
	Zipcode      string `json:"zipcode"`
	RestaurantId string `json:"restaurant_id"`
}
