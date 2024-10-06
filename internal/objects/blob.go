package objects

type Blob struct {
	Object
}

func NewBlob(data string) *Blob {
	if data == "" {
		data = "blob"
	}
	b := Blob{}
	b.Deserialize(data)
	return &b
}

func (b *Blob) Serialize() []byte {
	return b.data
}

func (b *Blob) Deserialize(data string) {
	b.data = []byte(data)
}
