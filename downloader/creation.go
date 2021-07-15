package downloader

// a struct to tell if the creation of a downloaded was successful or contained an error
type Creation struct {
	Successful string
	ErrorMsg   string
}

func newCreation(successful, errorMsg string) *Creation {
	return &Creation{Successful: successful, ErrorMsg: errorMsg}
}
