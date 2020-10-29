package jwebregistry

import (
    jwebmanager `gitlab.com/drjele-go/jweb/database/manager`
)

func New() *Registry {
    return &Registry{
        // env:      env,
        managers: map[string]jwebmanager.Manager{},
    }
}

type Registry struct {
    // env      *libEnv.Env
    managers map[string]jwebmanager.Manager
}

//
// func (m *Registry) GetManager(id string) jwebmanager.Manager {
//     _, ok := m.managers[id]
//
//     if ok == false {
//         m.managers[id] = m.initManager(m.env.GetApp().GetDatabaseConnection(id))
//     }
//
//     manager, _ := m.managers[id]
//
//     return manager
// }
//
// func (m *Registry) initManager(connection *jwebconnection.Connection) jwebmanager.Manager {
//     var db jwebmanager.Manager
//
//     switch connection.GetDriver() {
//     case jwebconnection.DriverMysql:
//         db = jwebmanager.NewMysql(connection, m.env.GetEnv())
//         break
//     case jwebconnection.DriverMongo:
//         db = jwebmanager.NewMongo(connection)
//         break
//     default:
//         jweberror.Fatal(jweberror.New(`invalid connection driver "%v"`, connection.GetDriver()))
//     }
//
//     return db
// }
