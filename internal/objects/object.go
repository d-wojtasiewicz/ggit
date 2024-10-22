package objects

import (
	"crypto/sha1"
	"encoding/hex"
)

type GitObject interface {
	Serialize() string
	Deserialize(data string) error
	Format() string
	Hash() (string, error)
}

type object struct {
	format string
}

func (o *object) Format() string {
	return o.format
}

func (o *object) Serialize() string {
	panic("serialize not implemented")
}
func (o *object) Deserialize(data string) error {
	panic("deserialize not implemented")
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
