package repository

import (
	"fmt"
	"ggit/internal/filesystem"
	"ggit/internal/objects"
)

func (r *Repository) CatObject(sha string) (string, error) {
	obj, err := r.ReadObject(sha)
	if err != nil {
		return "", err
	}

	return obj.(*objects.Blob).ReadData(), nil
}

type HashObject struct {
	File  string
	Type  string
	Write bool
}

func (r *Repository) HashObject(obj *HashObject) (string, error) {
	data, err := filesystem.ReadFileData(r.FS, obj.File)
	if err != nil {
		return "", err
	}

	var gitObject objects.GitObject
	switch obj.Type {
	case "blob":
		gitObject = objects.NewBlob(data)
	default:
		return "", fmt.Errorf("unknown type %s", obj.Type)
	}

	if obj.Write {
		return r.WriteObject(gitObject)
	}
	return gitObject.Hash()
}
