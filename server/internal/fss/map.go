package fss

type Game struct {
	Users   []*User
	Planets []*Planet
}

func NewGame() *Game {
	return &Game{}
}

type User struct{}

func NewUser() *User {
	return &User{}
}

type Planet struct {
	PlanetName string `json:"planetName"`
	PlanetSize int    `json:"planetSize"`
	Rooms      []*Room
}

func NewPlanet() *Planet {
	return &Planet{}
}

type Room struct {
	Rid      string `json:"rid"`
	Title    string `json:"title"`
	Describe string `json:"describe"`
	Ways     []struct {
		S string `json:"s"`
	} `json:"ways"`
	Mobs []struct {
		Mid  string `json:"mid"`
		Name string `json:"name"`
	} `json:"mobs"`
}

func NewRoom() *Room {
	return &Room{}
}
