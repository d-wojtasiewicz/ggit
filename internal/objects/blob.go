package objects

type Blob struct {
	*object
}

func NewBlob(data string) *Blob {
	b := &Blob{
		object: &object{
			format: "blob",
			data:   data,
		},
	}
	return b
}
