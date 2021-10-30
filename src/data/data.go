package data

import (
	"context"
	"log"
	"net/http"
	"treatMap/trick-or-treat/src/structs"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func get_Client() *firestore.Client {
	sa := option.WithCredentialsFile("./credFile.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func Get_Candy(g *gin.Context) {
	client := get_Client()
	result := client.Collection("Candies").Documents(context.Background())
	var res []structs.Candy
	var temp_Candy structs.Candy

	temp, err := result.GetAll()

	if err != nil {
		log.Fatalln(err)
	}
	for _, a := range temp {
		if err := a.DataTo(&temp_Candy); err != nil {
			log.Fatalln(err)
		}
		res = append(res, temp_Candy)
	}

	g.IndentedJSON(http.StatusOK, res)
	client.Close()
}

func Get_User(g *gin.Context) {
	client := get_Client()
	userid := g.Param("id")
	result, err := client.Collection("Users").Doc(userid).Get(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	var res structs.User
	if err := result.DataTo(&res); err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
	g.IndentedJSON(http.StatusOK, res)
	client.Close()

}

func Get_Report(g *gin.Context) {
	client := get_Client()
	result := client.Collection("Reports").Documents(context.Background())
	var res []structs.Report
	var temp_Report structs.Report

	temp, err := result.GetAll()

	if err != nil {
		log.Fatalln(err)
	}
	for _, a := range temp {
		if err := a.DataTo(&temp_Report); err != nil {
			log.Fatalln(err)
		}
		res = append(res, temp_Report)
	}

	log.Println(res)
	g.IndentedJSON(http.StatusOK, res)
	client.Close()
}

func Set_User(g *gin.Context) {
	client := get_Client()

	var user structs.User
	if err := g.BindJSON(&user); err != nil {
		log.Fatalln(err)
	}
	result, err := client.Collection("Users").Doc(user.UUID).Set(context.Background(), user)
	if err != nil {
		log.Fatalln(err)
	}
	g.IndentedJSON(http.StatusCreated, user)
	log.Println(result)
	client.Close()
}

func Set_Report(g *gin.Context) {
	client := get_Client()

	var report structs.Report
	if err := g.BindJSON(&report); err != nil {
		log.Fatalln(err)
	}
	result, err := client.Collection("Reports").Doc(report.ID).Set(context.Background(), report)
	if err != nil {
		log.Fatalln(err)
	}
	g.IndentedJSON(http.StatusCreated, report)
	log.Println(result)
	client.Close()
}
