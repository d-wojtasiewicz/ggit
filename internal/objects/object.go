package objects

import (
	"crypto/sha1"
	"encoding/hex"
	"strconv"
)

type GitObject interface {
	Serialize() []byte
	Deserialize(data string)
	Format() []byte
	Hash() string
}

type object struct {
	format []byte
	data   []byte
}

func (b *object) Serialize() []byte {
	return b.data
}

func (b *object) Deserialize(data []byte) {
	b.data = data
}

func (b *object) Format() []byte {
	return b.format
}

func (o *object) Hash() (string, error) {
	data := o.Serialize()
	result := o.Format()
	result = append(result, []byte(" ")...)
	result = append(result, []byte(strconv.Itoa(len(data)))...)
	result = append(result, []byte([]byte("\x00"))...)
	result = append(result, o.Serialize()...)

	hasher := sha1.New()
	_, err := hasher.Write(result)
	if err != nil {
		return "", err
	}
	sha := hex.EncodeToString(hasher.Sum(nil))
	return sha, nil
}
