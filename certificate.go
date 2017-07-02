package certificateS3

import (
	"io"
	"io/ioutil"
)

func getContentBytes(body io.ReadCloser) ([]byte, error) {
	content, err := ioutil.ReadAll(body)
	if err != nil {
		return make([]byte, 0), err
	}
	return content, nil
}
