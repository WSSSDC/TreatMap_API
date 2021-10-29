package structs

type Candy struct {
	ID    int8   `json:"id"`
	Name  string `json:"name"`
	Value int8   `json:"value"`
	Image string `json:"imageUrl"`
}

type User struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Points   int32  `json:"points"`
	Reports  int32  `json:"reports"`
	Ranking  int32  `json:"ranking"`
}

type Report struct {
	ID     int16   `json:"id"`
	Lng    float32 `json:"lng"`
	Lat    float32 `json:"lat"`
	Report []int8  `json:"report"`
}
