package facade

import "fmt"

type Dbc interface {
	Connect() string
}

func NewDbc() Dbc {
	return &dbc{
		a: NewMysqlDbc(),
		b: NewMongoDbc(),
	}
}

type dbc struct {
	a MysqlDbc
	b MongoDbc
}

func (d *dbc) Connect() string {
	retA := d.a.ConnectMysql()
	retB := d.b.ConnectMongo()
	return fmt.Sprintf("%s\n%s", retA, retB)
}

type MysqlDbc interface {
	ConnectMysql() string
}

type mysqlDbcImpl struct {
}

func NewMysqlDbc() MysqlDbc {
	return &mysqlDbcImpl{}
}

func (m *mysqlDbcImpl) ConnectMysql() string {
	return "connection of mysql."
}

var _ MysqlDbc = &mysqlDbcImpl{}

//====================

type MongoDbc interface {
	ConnectMongo() string
}

type mongoDbcImpl struct {
}

func NewMongoDbc() MongoDbc {
	return &mongoDbcImpl{}
}

func (m mongoDbcImpl) ConnectMongo() string {
	return "connection of mongo."
}

var _ MongoDbc = &mongoDbcImpl{}
