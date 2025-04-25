package domain

// get hotel by id
type Hotel struct {
	Id         int
	Name       string
	Location   Location
	Roomscount int
	Rating     float32
}
type Location struct {
	Country string
	City    string
	Street  string
	Number  int
	Zipcode int
}
