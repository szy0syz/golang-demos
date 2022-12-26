package api

type QuerySomething struct {
	Name string
}

type ResponseSomething struct {
	Namanya string `josn:"namanya" mapstructure:"name"`
}
