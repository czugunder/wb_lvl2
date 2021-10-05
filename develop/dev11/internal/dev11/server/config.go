package server

type config struct {
	address string
}

func DefaultConfig() *config {
	return &config{
		address: "localhost:8888",
	}
}

func NewConfig(address string) *config {
	return &config{
		address: address,
	}
}
