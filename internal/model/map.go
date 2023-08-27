package model

type Planet struct {
	PlanetName string `json:"planetName"`
	PlanetSize int    `json:"planetSize"`
	Rooms      []struct {
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
	} `json:"rooms"`
}
