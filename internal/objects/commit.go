package objects

type Commit struct {
	*object
	KVLM KWLM
}

func NewCommit() *Commit {
	return &Commit{
		object: &object{
			format: "commit",
		},
	}
}
