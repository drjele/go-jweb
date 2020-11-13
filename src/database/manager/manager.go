package jwebmanager

import (
    `go.mongodb.org/mongo-driver/mongo`
    `gorm.io/gorm`

    jwebconnection `gitlab.com/drjele-go/jweb/src/database/connection`
)

type Map map[string]Manager

type Manager interface {
    GetConnection() *jwebconnection.Connection

    GetClient() interface{}

    /** @todo maybe remove these shortcuts */
    /* shortcut methods used for valid return types */
    GetClientForMysql() *gorm.DB

    GetClientForMongo() *mongo.Client
}
