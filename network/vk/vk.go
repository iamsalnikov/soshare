package vk

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"

	"github.com/iamsalnikov/soshare/helpers/validator"
)

var countRX, _ = regexp.Compile("^VK.Share.count\\(1, (\\d+)\\);$")

// VK is a struct for working with vk.com
type VK struct {
	BaseURL string
}

// New function return VK instance
func New() *VK {
	return &VK{
		BaseURL: "http://vkontakte.ru/share.php",
	}
}

// GetShareCount return share count of specify url
func (v VK) GetShareCount(url string) (int64, error) {
	if !validator.IsURL(url) {
		return -1, errors.New(url + "is not valid url")
	}

	return v.sendRequest(url)
}

func (v VK) sendRequest(url string) (int64, error) {
	response, err := http.Get(v.BaseURL + "?act=count&index=1&url=" + url)
	if err != nil {
		return -1, err
	}

	if response.StatusCode != 200 {
		return -1, errors.New("vk.com response " + string(response.StatusCode))
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return -1, err
	}

	return v.extractCount(string(body))
}

func (v VK) extractCount(str string) (int64, error) {
	matches := countRX.FindAllStringSubmatch(str, 1)
	if len(matches) != 1 {
		return -1, errors.New("Share count not found")
	}

	shareCount, err := strconv.ParseInt(matches[0][1], 10, 64)
	if err != nil {
		return -1, err
	}

	return shareCount, nil
}
