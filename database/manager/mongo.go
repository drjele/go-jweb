package manager

import (
    `context`

    `go.mongodb.org/mongo-driver/mongo`
    `go.mongodb.org/mongo-driver/mongo/options`
    `gorm.io/gorm`

    `gitlab.com/drjele-go/jweb/database/connection`
    jweberror `gitlab.com/drjele-go/jweb/error`
)

func NewMongo(connection *connection.Connection) Manager {
    m := Mongo{connection: connection}

    client, err := mongo.NewClient(
        options.Client().ApplyURI(
            `mongodb://` + connection.GetUsername() + `:` + connection.GetPassword() +
                `@` + connection.GetHostname() + `:` + connection.GetPort(),
        ),
    )
    jweberror.Panic(err)

    err = client.Connect(context.TODO())
    jweberror.Panic(err)

    /** @todo maybe find a place to disconnect */

    m.client = client

    return &m
}

type Mongo struct {
    connection *connection.Connection
    client     *mongo.Client
}

func (m *Mongo) GetConnection() *connection.Connection {
    return m.connection
}

func (m *Mongo) GetClient() interface{} {
    return m.client
}

func (m *Mongo) GetClientForMysql() *gorm.DB {
    jweberror.Panic(jweberror.New(`mysql client not available for mongo manager`))
    return &gorm.DB{} /** fake mandatory return */
}

func (m *Mongo) GetClientForMongo() *mongo.Client {
    return m.GetClient().(*mongo.Client)
}
