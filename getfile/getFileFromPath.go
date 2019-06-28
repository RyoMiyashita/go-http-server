package getfile
import (
	"errors"
	"io/ioutil"
)

func GetFileFromPath(root string, path string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(root + path)
    if err != nil {
        return nil, errors.New("can't read file : " + root + path)
	}
	return bytes, nil
}
