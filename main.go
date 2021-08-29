package main

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
)

func main () {
	println("Initializing COLIH app ...")
	ctx := context.Background()
	conf := &firebase.Config{
		ProjectID: projectID
	}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	defer client.Close()


	iter := client.Collection("PERGUNTAR A GUGA").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		fmt.Println(doc.Data())

	}
}