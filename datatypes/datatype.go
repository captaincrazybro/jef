package datatypes

type datatype struct {
	name string
}

func (dt *datatype) GetName() string {
	return dt.name
}