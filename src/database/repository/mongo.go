package jwebrepository

import (
    `context`

    `go.mongodb.org/mongo-driver/bson`
    `go.mongodb.org/mongo-driver/bson/primitive`
    `go.mongodb.org/mongo-driver/mongo`

    jwebentity `gitlab.com/drjele-go/jweb/src/database/entity`
    jwebmanager `gitlab.com/drjele-go/jweb/src/database/manager`
    jweberror `gitlab.com/drjele-go/jweb/src/error`
)

/** @todo format this to return basic types */
type MongoAgregateResult []bson.M

type Mongo struct {
    Manager *jwebmanager.Mongo
}

func (m *Mongo) Agregate(result interface{}, collectionName string, filter bson.D, groupBy bson.D) {
    ctx := context.TODO()

    db := m.Manager.GetClientForMongo().Database(m.Manager.GetConnection().GetDatabase())

    collection := db.Collection(collectionName)

    cursor, err := collection.Aggregate(
        ctx,
        mongo.Pipeline{
            bson.D{{`$match`, filter}},
            bson.D{{`$group`, groupBy}},
        },
    )
    jweberror.Panic(err)

    defer m.closeCursor(cursor, ctx)

    err = cursor.All(ctx, result)
    jweberror.Panic(err)
}

func (m *Mongo) Find(result interface{}, collectionName string, filter interface{}) {
    if filter == nil {
        filter = bson.M{}
    }

    ctx := context.TODO()

    db := m.Manager.GetClientForMongo().Database(m.Manager.GetConnection().GetDatabase())

    collection := db.Collection(collectionName)

    cursor, err := collection.Find(ctx, filter)
    jweberror.Panic(err)

    defer m.closeCursor(cursor, ctx)

    err = cursor.All(ctx, result)
    jweberror.Panic(err)
}

func (m *Mongo) FindOne(result interface{}, collectionName string, filter interface{}) {
    if filter == nil {
        filter = bson.M{}
    }

    ctx := context.TODO()

    db := m.Manager.GetClientForMongo().Database(m.Manager.GetConnection().GetDatabase())

    collection := db.Collection(collectionName)

    err := collection.FindOne(ctx, filter).Decode(result)

    if err == mongo.ErrNoDocuments {
        return
    }

    jweberror.Panic(err)
}

func (m *Mongo) Insert(document jwebentity.Mongo) string {
    ctx := context.TODO()

    db := m.Manager.GetClientForMongo().Database(m.Manager.GetConnection().GetDatabase())

    result, err := db.Collection(document.GetCollectionName()).InsertOne(ctx, document)
    jweberror.Panic(err)

    return result.InsertedID.(primitive.ObjectID).String()
}

func (m *Mongo) closeCursor(cursor *mongo.Cursor, ctx context.Context) {
    err := cursor.Close(ctx)
    jweberror.Panic(err)
}
