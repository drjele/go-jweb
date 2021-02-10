package manager

import (
    `go.mongodb.org/mongo-driver/mongo`
    `gorm.io/gorm`

    `gitlab.com/drjele-go/jweb/database/connection`
)

type Map map[string]Manager

type Manager interface {
    GetConnection() *connection.Connection

    GetClient() interface{}

    /** @todo maybe remove these shortcuts */
    /* shortcut methods used for valid return types */
    GetClientForMysql() *gorm.DB

    GetClientForMongo() *mongo.Client
}
