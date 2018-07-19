package main

import (
	"context"
	"testing"

	"github.com/golang/glog"
)

func BenchmarkSnapshot(t *testing.B) {
	ctx := context.Background()
	app := initFirebase()
	client, err := app.Firestore(ctx)
	if err != nil {
		glog.Errorln(err)
	}

	collection := client.Collection("users-dev")

	// documentを指定して取得
	doc := collection.Doc("test")
	snapshot, err := doc.Get(ctx)
	if err != nil {
		glog.Errorln(err)
	}

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		_ = snapshot.Data()
	}
	t.StopTimer()
}

func BenchmarkMapping(t *testing.B) {
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

	collection := client.Collection("users-dev")

	// documentを指定して取得
	doc := collection.Doc("test")
	snapshot, err := doc.Get(ctx)
	ent := writeData{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		_ = snapshot.DataTo(&ent)
	}
	t.StopTimer()
}
