package linkedin

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/iamsalnikov/soshare/helpers/validator"
)

// Linkedin struct is a linkedin share checker
type Linkedin struct {
	BaseURL string
}

// New function return instance of Facebook checker
func New() *Linkedin {
	return &Linkedin{
		BaseURL: "https://www.linkedin.com/countserv/count/share",
	}
}

// GetShareCount return share count
func (l Linkedin) GetShareCount(url string) (int64, error) {
	if !validator.IsURL(url) {
		return -1, errors.New(url + " - is not valid url")
	}

	return l.sendRequest(url)
}

func (l Linkedin) sendRequest(url string) (int64, error) {
	response, err := http.Get(l.BaseURL + "?format=json&url=" + url)
	if err != nil {
		return -1, err
	}

	if response.StatusCode != 200 {
		return -1, errors.New("Linkedin answer status is " + string(response.StatusCode))
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

	return answer.Count, nil
}
