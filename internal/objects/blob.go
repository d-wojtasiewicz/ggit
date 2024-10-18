package objects

type Blob struct {
	object
}

func NewBlob(data string) *Blob {
	b := Blob{}
	b.format = "blob"
	b.Data = data
	return &b
}
