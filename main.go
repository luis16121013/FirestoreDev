package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	//	"google.golang.org/api/option"
	"google.golang.org/api/iterator"
)

/*
type Post struct{
	ID int64
	Title string
	Text string
}*/

func main() {
	ctx := context.Background()
	//sa := option.WithCredentialsFile("./godev-64e81-firebase-adminsdk-4n5b6-89354ebbcd.json")
	//app,err := firebase.NewApp(ctx,nil,sa)
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	/*
			_, _, err = client.Collection("post").Add(ctx, map[string]interface{}{
		        "ID": 5,
		        "Text":  "alvaro",
		        "Title":  "pfcanales",
			})
			if err != nil {
		        log.Fatalf("Failed adding COPYPAGE: %v", err)
			}
	*/
	defer client.Close()

	/*data := Post{
		Password: "carla12",
		Username: "calita",
	}*/
	/*
			_, err1 := client.Collection("post").Doc("new-post-id").Set(ctx, data)
			if err1 != nil {
			  // Handle any errors in an appropriate way, such as returning them.
		  	log.Printf("An error has occurred: %s", err)
			}*/

	/*
		iter := client.Collection("post").Documents(ctx)
		for{
			doc,err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil{
				log.Fatalf("FAILED TO ITERATE: %v",err)
				break
			}
			fmt.Println(doc.Data()["Title"].(string))
		}*/

	i := client.Collection("users").Documents(ctx)
	for {
		doc, err := i.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("error => %v", err)
			break
		}

		if doc.Data()["Username"].(string) == "luis" && doc.Data()["Password"].(string) == "1234" {
			fmt.Println("user encontrado")
		}
	}

}
