package singleton

type Singleton struct {
}

var instance *Singleton = &Singleton{}

func GetInstance() *Singleton {
	return instance
}
