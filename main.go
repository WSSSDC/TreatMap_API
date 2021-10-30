package main

import (
	"context"
	"log"
	data_api "treatMap/trick-or-treat/src/data"
	"treatMap/trick-or-treat/src/structs"
	s "treatMap/trick-or-treat/src/structs"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	sa := option.WithCredentialsFile("./credFile.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	test_user := s.User{
		UUID:     "test",
		Username: "Bramogus",
		Points:   100,
		Reports:  100,
		Ranking:  100,
	}
	var user structs.User
	data_api.Set_User(test_user, client)
	user = data_api.Get_User(test_user.UUID, client)
	log.Print(user)
	client.Close()
}
