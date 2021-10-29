package data

import (
	"context"
	"log"
	"treatMap/trick-or-treat/src/structs"

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
	client.Close()
}

func Get_Candy() *structs.Candy {
	//TODO: create a function to get a Candy value from Firestore
	return nil
}

func Get_User() *structs.User {
	//TODO: create a function to get a User value from Firestore
	return nil
}

func Get_Report() *structs.Report {
	//TODO: create a function to get a Report value from Firestore
	return nil
}

func Set_User() {

}

func Set_Report() {

}
