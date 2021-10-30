package data

import (
	"context"
	"log"
	"treatMap/trick-or-treat/src/structs"

	"cloud.google.com/go/firestore"
)

func Get_Candy(candyid string, client *firestore.Client) structs.Candy {
	result, err := client.Collection("Candies").Doc("kitkat").Get(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	var res structs.Candy

	if err := result.DataTo(&res); err != nil {
		log.Fatalln(err)
	}
	return res
}

func Get_User(uuid string, client *firestore.Client) structs.User {
	result, err := client.Collection("Users").Doc(uuid).Get(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	var res structs.User
	if err := result.DataTo(&res); err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
	return res
}

func Get_Report(id string, client *firestore.Client) structs.Report {
	result, err := client.Collection("Users").Doc(id).Get(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	var res structs.Report
	if err := result.DataTo(&res); err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
	return res
}

func Set_User(user structs.User, client *firestore.Client) {
	result, err := client.Collection("Users").Doc(user.UUID).Set(context.Background(), user)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result)
}

func Set_Report(report *structs.Report, client *firestore.Client) {
	result, err := client.Collection("Reports").Doc(report.ID).Set(context.Background(), report)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result)
}
