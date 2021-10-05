package server

// Config тип, хранящий конфигурацию сервера
type Config struct {
	address string
}

// DefaultConfig возвращает конфигурацию по умолчанию
func DefaultConfig() *Config {
	return &Config{
		address: "localhost:8888",
	}
}

// NewConfig создает экземпляр Config
func NewConfig(address string) *Config {
	return &Config{
		address: address,
	}
}
