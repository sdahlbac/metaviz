package generator

type Entity struct {
	Name        string
	Description string
	Fields      []Field
	Relations   []Relation
}

type Field struct {
	Id          string
	Name        string
	Type        string
	Mandatory   bool
	Description string
}

type Relation struct {
	Type   string
	Id     string
	Name   string
	Target Target
}

type Target struct {
	Table  string
	Fields map[string]interface{}
}
