package objects

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type Blob struct {
	*object
	data string
}

func NewBlob(data string) *Blob {
	b := &Blob{
		object: &object{
			format: "blob",
		},
		data: data,
	}
	return b
}

func (b *Blob) SetData(data string) {
	b.data = data
}

func (b *Blob) ReadData() string {
	return b.data
}

func (o *Blob) Serialize() string {
	return fmt.Sprintf("%v %v%v%v", o.format, len(o.data), "\x00", o.data)
}

func (o *Blob) Deserialize(data string) error {
	x := strings.Index(data, " ")
	format := data[0:x]

	y := strings.Index(data, "\x00")
	size, err := strconv.Atoi(data[x+1 : y])
	if err != nil {
		return fmt.Errorf("unable to read object size")
	}

	if size != len(data)-y-1 {
		return fmt.Errorf("malformed object %s: bad length", data)
	}
	o.format = format
	o.data = data[y+1:]
	return nil
}

func (b *Blob) Hash() (string, error) {
	data := b.Serialize()
	hasher := sha1.New()
	_, err := hasher.Write([]byte(data))
	if err != nil {
		return "", err
	}
	sha := hex.EncodeToString(hasher.Sum(nil))
	return sha, nil
}
