package objects

type Blob struct {
	object
}

func NewBlob(data []byte) *Blob {
	b := Blob{}
	b.format = []byte("blob")
	b.Deserialize(data)
	return &b
}
