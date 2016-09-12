package facebook

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/iamsalnikov/soshare/helpers/validator"
)

// Facebook struct is a facebook share checker
type Facebook struct {
	BaseURL string
}

// New function return instance of Facebook checker
func New() *Facebook {
	return &Facebook{
		BaseURL: "http://graph.facebook.com",
	}
}

// GetShareCount return share count
func (f Facebook) GetShareCount(url string) (int64, error) {
	if !validator.IsURL(url) {
		return -1, errors.New(url + " - is not valid url")
	}

	return f.sendRequest(url)
}

func (f Facebook) sendRequest(address string) (int64, error) {
	response, err := http.Get(f.BaseURL + "?id=" + address)
	if err != nil {
		return -1, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return -1, err
	}

	var answer Answer
	err = json.Unmarshal(body, &answer)
	if err != nil {
		return -1, err
	}

	var nilError AnswerError
	if answer.Error != nilError {
		return -1, errors.New(answer.Error.Message)
	}

	return answer.Share.ShareCount, nil
}
