package configuration

type Provider interface {
	GetString(key string) string
	GetStringOr(key string, fallback string) string
	GetInt(key string) int
	GetIntOr(key string, fallback int) int
}
