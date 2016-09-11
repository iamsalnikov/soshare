package facebook

type Facebook struct {
	BaseURL string
}

func New() *Facebook {
	return &Facebook{
		BaseURL: "http://graph.facebook.com",
	}
}
