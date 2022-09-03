package singleton

var dbInstance *database

type database struct {
	name string
}

func Instance() *database {
	if dbInstance == nil {
		dbInstance = &database{}
	}

	return dbInstance
}

func (db database) Name() string {
	return db.name
}

func (db *database) SetName(name string) {
	db.name = name
}
