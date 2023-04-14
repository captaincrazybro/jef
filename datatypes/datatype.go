package datatypes

type dataType struct {
	name string
}

func (dt dataType) GetName() string {
	return dt.name
}
