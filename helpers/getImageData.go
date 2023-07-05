package helpers

import (
	"errors"
	"io/ioutil"
	"mime/multipart"
	"regexp"
)

func GetImageData(file *multipart.FileHeader) ([]byte, error) {
	// check if the file extension is valid
	validExt := regexp.MustCompile(`\.(jpg|png)$`)
	if !validExt.MatchString(file.Filename) {

		return []byte{}, errors.New("only allow to upload jpg and png file")
	}

	// open file to take data
	data, err := file.Open()
	defer data.Close()
	if err != nil {
		return []byte{}, err
	}

	// get binary file from the data to save to mongoDB
	fileBytes, err := ioutil.ReadAll(data)
	if err != nil {
		return []byte{}, err
	}
	
	return fileBytes, nil
}
