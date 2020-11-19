package main

import (
	"log"
	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

ctx := context.Background()
sa := option.WithCredentialsFile("./godev-64e81-firebase-adminsdk-4n5b6-89354ebbcd.json")
app, err != nil {
	log.Fatalln(err)
}

client, err := app.Firestore(ctx)
if err != nil {
	log.Fatalln(err)
}
defer client.Close()


iter := client.Collection("post").Documents(ctx)
for{
	doc,err := iter.Next()
	if err != iterator.Done{
		break
	}
	if err != nil{
		log.Fatalf("FAILED TO ITERATE: %v",err)
	}
	fmt.Println(doc.Data())
}
