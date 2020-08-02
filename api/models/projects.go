package models

import (
	// builtin
	"context"
	"html"
	"options"
	"runtime/debug"
	"time"
	"strings"
	
	// self
	"github.com/UTx10101/scrapehero/db"
	"github.com/UTx10101/scrapehero/utils"
	
	// vendored
	"github.com/apex/log"
	"github.com/globalsign/mgo/bson"
	
)

type Project struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	Name        string        `json:"name" bson:"name"` 
	CreateTs time.Time     `json:"create_ts" bson:"create_ts"`
	UpdateTs time.Time     `json:"update_ts" bson:"update_ts"`
}

func (p *Project) Create() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := db.DBClient.Connect(ctx); err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return err
	}
	
	defer db.DBClient.Disconnect(ctx)
	
	col := db.DBClient.Database("scphero").Collection("projects")
	p.ID = bson.NewObjectId()
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.UpdateTs = time.Now()
	p.CreateTs = time.Now()
	
	if res, err := col.InsertOne(ctx, p); err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return err
	}
	
	fmt.Println("Created Project: ", res.InsertedID)
	
	if err := utils.InitProject(p.ID.(bson.ObjectID).Hex()) {
		debug.PrintStack()
		log.Fatal(err)
		return err
	}

	return nil
}

func GetProject(pid bson.ObjectId) (Project, error) {
	var p Project
	
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := db.DBClient.Connect(ctx); err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return p, err
	}
	
	defer db.DBClient.Disconnect(ctx)
	
	col := db.DBClient.Database("scphero").Collection("projects")
	
	if err := col.FindOne(ctx, bson.D{{"_id", pid}}).Decode(&p); err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return p, err
	}
	return p, nil
}

func GetProjects(filter interface{}) ([]Project, error) {
	var projects []Project
	
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := db.DBClient.Connect(ctx); err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return projects, err
	}
	
	defer db.DBClient.Disconnect(ctx)
	
	col := db.DBClient.Database("scphero").Collection("projects")
	opts := options.Find()
	opts.SetSort(bson.D{{"create_ts", -1}})
	
	if err := col.Find(ctx, filter, opts).All(ctx, &projects); err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return projects, err
	}

	return projects, nil
}

func GetProjectsCount(filter interface{}) (int, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := db.DBClient.Connect(ctx); err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return 0, err
	}
	
	defer db.DBClient.Disconnect(ctx)
	
	col := db.DBClient.Database("scphero").Collection("projects")

	if count, err := col.Find(ctx, filter).Count(); err != nil {
		return 0, err
	}

	return count, nil
}

func UpdateProject(pid bson.ObjectId, item Project) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := db.DBClient.Connect(ctx); err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return err
	}
	
	defer db.DBClient.Disconnect(ctx)
	
	col := db.DBClient.Database("scphero").Collection("projects")
	item.Name = html.EscapeString(strings.TrimSpace(item.Name))
	item.UpdateTs = time.Now()
	update = bson.D{{"$set", item}}
	
	if res, err := col.UpdateOne(ctx, bson.M{"_id": pid}, update); err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return err
	}
	
	fmt.Printf("Updated %v Project!\n", res.ModifiedCount)
	
	return nil
}

func RemoveProject(pid bson.ObjectId) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := db.DBClient.Connect(ctx); err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return err
	}
	
	defer db.DBClient.Disconnect(ctx)
	
	col := db.DBClient.Database("scphero").Collection("projects")

	if res, err := col.DeleteOne(ctx, bson.M{"_id": pid}); err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return err
	}
	
	if err := utils.DeleteProject(pid.(bson.ObjectID).Hex()) {
		debug.PrintStack()
		log.Fatal(err)
		return err
	}
	
	return nil
}