package linkedin

// Answer struct represent linkedin answer
type Answer struct {
	Count       int64  `json:"count"`
	FCnt        string `json:"fCnt"`
	FCntPlusOne string `json:"fCntPlusOne"`
	URL         string `json:"url"`
}
