/*
Package xmldecoding provides functions for decoding XML files into Presume models.
*/
package xmldecoding

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/b-vennes/presume/pkg/models"
)

type newDecoderError struct{}

func (e newDecoderError) Error() string {
	return fmt.Sprintln("Failed to create XML file decoder.")
}

func makeNewDecoderError() newDecoderError {
	return newDecoderError{}
}

// Decodes resume content from the given XML file path.
func Decode(filePath string) (*models.ResumeContentData, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := xml.NewDecoder(file)

	if decoder == nil {
		err = makeNewDecoderError()
		return nil, err
	}

	var result models.ResumeContentData
	err = decoder.Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
