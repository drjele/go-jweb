package manager

import (
    `go.mongodb.org/mongo-driver/mongo`
    `gorm.io/driver/mysql`
    `gorm.io/gorm`
    `gorm.io/gorm/logger`
    `gorm.io/gorm/schema`

    `gitlab.com/drjele-go/jweb/database/connection`
    jweberror `gitlab.com/drjele-go/jweb/error`
    `gitlab.com/drjele-go/jweb/kernel/environment`
)

func NewMysql(connection *connection.Connection, env string) Manager {
    m := Mysql{connection: connection}

    loggerLevel := logger.Info
    if env == environment.EnvProd {
        loggerLevel = logger.Warn
    }

    gormConfig := gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            SingularTable: true,
        },
        PrepareStmt: true,
        Logger:      logger.Default.LogMode(loggerLevel),
    }

    dsn := connection.GetUsername() + `:` + connection.GetPassword() +
        `@tcp(` + connection.GetHostname() + `:` + connection.GetPort() + `)/` + connection.GetDatabase() +
        `?charset=utf8mb4&parseTime=True&loc=Local` /** @todo have them as settings or config */

    db, err := gorm.Open(mysql.Open(dsn), &gormConfig)
    jweberror.Panic(err)

    db.Set(`gorm:table_options`, `ENGINE=InnoDB`)

    m.client = db

    return &m
}

type Mysql struct {
    connection *connection.Connection
    client     *gorm.DB
}

func (m *Mysql) GetConnection() *connection.Connection {
    return m.connection
}

func (m *Mysql) GetClient() interface{} {
    return m.client
}

func (m *Mysql) GetClientForMysql() *gorm.DB {
    return m.GetClient().(*gorm.DB)
}

func (m *Mysql) GetClientForMongo() *mongo.Client {
    jweberror.Panic(jweberror.New(`mongo client not available for mysql manager`))
    return &mongo.Client{} /** fake mandatory return */
}
