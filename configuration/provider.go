package configuration

type Provider interface {
	GetString(key string, fallback string) string
	GetInt(key string, fallback int) int
	//getBool(key bool, fallback bool)
}
