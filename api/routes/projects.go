package model

import (
	// builtin
	"context"
	"runtime/debug"
	"time"
	
	// vendored
	"github.com/UTx10101/scrapehero/constants"
	"github.com/UTx10101/scrapehero/db"
	"github.com/apex/log"
	"github.com/globalsign/mgo/bson"
	
)

type Project struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	Name        string        `json:"name" bson:"name"` 
	CreateTs time.Time     `json:"create_ts" bson:"create_ts"`
	UpdateTs time.Time     `json:"update_ts" bson:"update_ts"`
}

func (p *Project) Update() error {	
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := db.DBClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	
	defer db.DBClient.Disconnect(ctx)
	
	col := db.DBClient.Database("scphero").Collection("projects")
	
	p.UpdateTs = time.Now()
	
	filter := 
	
	if res, err := col.UpdateOne(ctx, bson.D{{"_id", p.ID}}, p); err != nil {
		log.Fatal(err)
		debug.PrintStack()
		return err
	}

	fmt.Printf("Updated %v projects.\n", res.ModifiedCount)

	return nil
}

func (p *Project) Create() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := db.DBClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	
	defer db.DBClient.Disconnect(ctx)
	
	col := db.DBClient.Database("scphero").Collection("projects")
	
	p.ID = bson.NewObjectId()
	p.UpdateTs = time.Now()
	p.CreateTs = time.Now()
	
	if res, err := col.InsertOne(ctx, p); err != nil {
		log.Fatal(err)
		debug.PrintStack()
		return err
	}
	
	fmt.Println("Project Created: ", res.InsertedID)

	return nil
}

func GetProject(id bson.ObjectId) (Project, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := db.DBClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	
	defer db.DBClient.Disconnect(ctx)
	
	col := db.DBClient.Database("scphero").Collection("projects")
	
	var p Project
	
	if err := col.FindOne(ctx, bson.D{{"_id", id}}).Decode(&&p); err != nil {
		log.Fatal(err)
		debug.PrintStack()
		return p, err
	}
	return p, nil
}

func GetProjects(filter interface{}, sortKey string) ([]Project, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := db.DBClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	
	defer db.DBClient.Disconnect(ctx)
	
	col := db.DBClient.Database("scphero").Collection("projects")

	var projects []Project
	
	if err := col.Find(filter).Sort(sortKey).All(&projects); err != nil {
		debug.PrintStack()
		return projects, err
	}

	return projects, nil
}

func GetProjectsCount(filter interface{}) (int, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := db.DBClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	
	defer db.DBClient.Disconnect(ctx)
	
	col := db.DBClient.Database("scphero").Collection("projects")

	count, err := c.Find(filter).Count()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func UpdateProject(id bson.ObjectId, item Project) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := db.DBClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	
	defer db.DBClient.Disconnect(ctx)
	
	col := db.DBClient.Database("scphero").Collection("projects")

	var res Project
	if err := col.FindOne(ctx, bson.M{"_id": id}).Decode(&res); err != nil {
		debug.PrintStack()
		return err
	}

	if err := item.Update(); err != nil {
		return err
	}
	return nil
}

func RemoveProject(id bson.ObjectId) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := db.DBClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	
	defer db.DBClient.Disconnect(ctx)
	
	col := db.DBClient.Database("scphero").Collection("projects")

	var result User
	if err := c.FindId(id).One(&result); err != nil {
		return err
	}

	if err := c.RemoveId(id); err != nil {
		return err
	}

	return nil
}