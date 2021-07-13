package datatypes

// Datatype interface to store a Datatype
type Datatype interface {
	GetName() string
	Check(string) bool
	GetValue(string) (interface{}, error)
}