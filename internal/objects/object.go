package objects

type GitObject interface {
	Serialize() []byte
	Deserialize(data string)
}

type Object struct {
	data []byte
}

func (o *Object) Serialize() []byte {
	panic("serialize not implemented")
}

func (o *Object) Deserialize(data string) {
	panic("deserialize not implemented")
}
