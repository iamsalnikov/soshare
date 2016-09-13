package googleplus

//import
import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/iamsalnikov/soshare/helpers/validator"
)

// GooglePlus struct is a google plus share checker
type GooglePlus struct {
	BaseURL string
}

// New function return GooglePlus instance
func New() *GooglePlus {
	return &GooglePlus{
		BaseURL: "https://clients6.google.com/rpc",
	}
}

// GetShareCount return share count of url in Google+
func (gp GooglePlus) GetShareCount(url string) (int64, error) {
	if !validator.IsURL(url) {
		return -1, errors.New(url + " - is not valid url")
	}

	return gp.sendRequest(url)
}

func (gp GooglePlus) sendRequest(address string) (int64, error) {
	request, err := gp.prepareRequest(address)
	if err != nil {
		return -1, err
	}

	response, err := gp.getResponseBody(request)
	if err != nil {
		return -1, err
	}

	var answer Answer
	err = json.Unmarshal(response, &answer)
	if err != nil {
		return -1, nil
	}

	var nilErrorAnswer Answer
	if answer.Error != nilErrorAnswer.Error {
		return -1, errors.New(answer.Error.Message)
	}

	return answer.Result.Metadata.GlobalCounts.Count, nil
}

func (gp GooglePlus) prepareRequest(address string) (*http.Request, error) {
	data := gp.getRequestBody(address)
	request, err := http.NewRequest("POST", gp.BaseURL, bytes.NewBuffer(data))
	if err != nil {
		return &http.Request{}, err
	}

	request.Header.Set("Content-Type", "application/json")
	return request, nil
}

func (gp GooglePlus) getRequestBody(address string) []byte {
	return []byte(`{"method":"pos.plusones.get","id":"p","params":{"nolog":true,"id":"` + address + `","source":"widget","userId":"@viewer","groupId":"@self"},"jsonrpc":"2.0","key":"p","apiVersion":"v1"}`)
}

func (gp GooglePlus) getResponseBody(request *http.Request) ([]byte, error) {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("Google answer status " + string(response.StatusCode))
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
