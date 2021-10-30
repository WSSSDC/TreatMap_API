package structs

type Candy struct {
	ID    string `firestore:"id"`
	Name  string `firestore:"name"`
	Value int64  `firestore:"value"`
	Image string `firestore:"imageUrl"`
}

type User struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Points   int32  `json:"points"`
	Reports  int32  `json:"reports"`
	Ranking  int32  `json:"ranking"`
}

type Report struct {
	ID     string  `json:"id"`
	Lng    float32 `json:"lng"`
	Lat    float32 `json:"lat"`
	Report []int8  `json:"report"`
}
