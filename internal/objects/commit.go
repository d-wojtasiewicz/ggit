package objects

import (
	"crypto/sha1"
	"encoding/hex"
)

type Commit struct {
	*object
	KVLM kvlm
}

func NewCommit() *Commit {
	return &Commit{
		object: &object{
			format: "commit",
		},
	}
}

func (c *Commit) Serialize() string {
	return c.KVLM.Serialize()
}
func (c *Commit) Deserialize(data string) error {
	var kvlm kvlm
	kvlm.Deserialize(data)
	c.KVLM = kvlm
	return nil
}
func (c *Commit) Hash() (string, error) {
	data := c.Serialize()
	hasher := sha1.New()
	_, err := hasher.Write([]byte(data))
	if err != nil {
		return "", err
	}
	sha := hex.EncodeToString(hasher.Sum(nil))
	return sha, nil
}
