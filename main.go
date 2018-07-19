package main

import (
	"context"

	"firebase.google.com/go"
	"github.com/golang/glog"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	app := initFirebase()
	client, err := app.Firestore(ctx)
	if err != nil {
		glog.Errorln(err)
	}

	type state struct {
		Height int    `firestore:"height"`
		Sex    string `firestore:"sex"`
		Age    int    `firestore:"age"`
	}

	type writeData struct {
		UserID int64  `firestore:"user_id"`
		Email  string `firestore:"email"`
		State  state  `firestore:"state"`
	}

	/*
		collection := client.Collection("users-dev")
		doc := collection.Doc("test")
		result, err := doc.Set(ctx, writeData{
			UserID: 40000,
			Email:  "takochuu@hogehoge.jp",
			State: state{
				Sex:    "male",
				Height: 160,
				Age:    25,
			},
		})
		if err != nil {
			glog.Errorln(err)
		}
	*/

	collection := client.Collection("users-dev")

	// documentを指定して取得
	doc := collection.Doc("test")
	snapshot, err := doc.Get(ctx)
	if err != nil {
		glog.Errorln(err)
	}
	_ = snapshot.Data()
	if err != nil {
		glog.Errorln(err)
	}

	// 構造体にマッピング
	ent := writeData{}
	if err = snapshot.DataTo(&ent); err != nil {
		glog.Errorln(err)
	}

	// Queryで取得する方法
	_ = client.Collection("users-dev").Where("user_id", "=", 40000)
}

func initFirebase() *firebase.App {
	opt := option.WithCredentialsFile("/path/to/development.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		glog.Errorln("Error")
	}
	return app
}
