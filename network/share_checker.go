package network

// ShareChecker interface must be implemented by networks
type ShareChecker interface {
	GetShareCount(url string) (int64, error)
}
