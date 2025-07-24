package internal

type Config struct{
	URL string
	Requests int
	Concurrency int
}

type Result struct {
	StatusCode int
	Error error
}
