package objects

type Blob struct {
	*object
}

func NewBlob(data string) *Blob {
	b := &Blob{
		object: &object{
			format: "blob",
			Data:   data,
		},
	}
	return b
}
