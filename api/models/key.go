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
	"github.com/UTx10101/scrapehero/auth"
	"github.com/UTx10101/scrapehero/security"
	"github.com/UTx10101/scrapehero/db"
	"github.com/UTx10101/scrapehero/utils"
	
	// vendored
	"github.com/apex/log"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
	
)

type APIKey struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	Key         string        `json:"key" bson:"key"`
	Status      string        `json:"status" bson:"status"`
	CreateTs time.Time     `json:"create_ts" bson:"create_ts"`
}

func (ak *APIKey) Create() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := db.DBClient.Connect(ctx); err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return err
	}
	
	defer db.DBClient.Disconnect(ctx)
	
	col := db.DBClient.Database("scphero").Collection("apikeys")
	ak.ID = bson.NewObjectId()
	
	if ak.Key, err := auth.CreateToken('ANY', ak.ID.(bson.ObjectID).Hex(), 'API'); err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return err
	}
	
	ak.Key = security.Hash(ak.Key)
	
	ak.Key = html.EscapeString(strings.TrimSpace(ak.Key))
	ak.Status = html.EscapeString(strings.TrimSpace(ak.Status))
	ak.CreateTs = time.Now()
	
	if res, err := col.InsertOne(ctx, ak); err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return err
	}
	
	fmt.Println("Created APIKey: ", res.InsertedID)

	return nil
}

func GetKeys(filter interface{}) ([]APIKey, error) {
	var keys []APIKey
	
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := db.DBClient.Connect(ctx); err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return keys, err
	}
	
	defer db.DBClient.Disconnect(ctx)
	
	col := db.DBClient.Database("scphero").Collection("apikeys")
	opts := options.Find()
	opts.SetSort(bson.D{{"create_ts", -1}})
	
	if err := col.Find(ctx, filter, opts).All(ctx, &keys); err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return keys, err
	}

	return keys, nil
}

func GetKeysCount(filter interface{}) (int, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := db.DBClient.Connect(ctx); err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return 0, err
	}
	
	defer db.DBClient.Disconnect(ctx)
	
	col := db.DBClient.Database("scphero").Collection("apikeys")

	if count, err := col.Find(ctx, filter).Count(); err != nil {
		return 0, err
	}

	return count, nil
}

func CheckAPIKey(token string) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := db.DBClient.Connect(ctx); err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return err
	}
	
	defer db.DBClient.Disconnect(ctx)
	
	col := db.DBClient.Database("scphero").Collection("apikeys")

	if count, err := col.Find(ctx, filter).Count(); err != nil {
		return errors.New("unauthorized")
	}

	return nil
}

func UpdateKeyStatus(kid bson.ObjectId, status string) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := db.DBClient.Connect(ctx); err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return err
	}
	
	defer db.DBClient.Disconnect(ctx)
	
	col := db.DBClient.Database("scphero").Collection("apikeys")
	status = html.EscapeString(strings.TrimSpace(status))
	update = bson.D{{"$set", bson.D{{"status", status}}}}
	
	if res, err := col.UpdateOne(ctx, bson.M{"_id": kid}, update); err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return err
	}
	
	fmt.Printf("Updated %v APIKey!\n", res.ModifiedCount)
	
	return nil
}

func RemoveKey(kid bson.ObjectId) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := db.DBClient.Connect(ctx); err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return err
	}
	
	defer db.DBClient.Disconnect(ctx)
	
	col := db.DBClient.Database("scphero").Collection("apikeys")

	res, err := col.DeleteOne(ctx, bson.M{"_id": kid})
	if err != nil {
		debug.PrintStack()
		log.Fatal(err)
		return err
	}
	
	return nil
}