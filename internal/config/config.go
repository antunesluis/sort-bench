package config

// Config armazena as configurações globais do sistema.
type Config struct {
	InputFile  string
	OutputFile string
	Algorithm  string
	Mode       string
	Analyze    bool
}
