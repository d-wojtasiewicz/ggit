package objects

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type GitObject interface {
	Serialize() string
	Deserialize(data string) error
	Format() string
	Hash() (string, error)
	ReadData() string
	SetData(string)
}

type object struct {
	format string
	data   string
}

func (o *object) Serialize() string {
	return fmt.Sprintf("%v %v%v%v", o.format, len(o.data), "\x00", o.data)
}

func (o *object) Deserialize(data string) error {
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

func (b *object) Format() string {
	return b.format
}

func (o *object) SetData(data string) {
	o.data = data
}

func (o *object) ReadData() string {
	return o.data
}

func (o *object) Hash() (string, error) {
	data := o.Serialize()
	hasher := sha1.New()
	_, err := hasher.Write([]byte(data))
	if err != nil {
		return "", err
	}
	sha := hex.EncodeToString(hasher.Sum(nil))
	return sha, nil
}
