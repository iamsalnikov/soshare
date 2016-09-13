package googleplus

// Answer struct represent google plus API answer
type Answer struct {
	ID string `json:"id"`

	Result struct {
		Kind          string `json:"kind"`
		ID            string `json:"id"`
		IsSetByViewer bool   `json:"isSetByViewer"`
		Metadata      struct {
			Type         string `json:"type"`
			GlobalCounts struct {
				Count int64 `json:"count"`
			} `json:"globalCounts"`
		} `json:"metadata"`
	} `json:"result"`

	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}
