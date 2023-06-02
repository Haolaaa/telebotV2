package config

type Config struct {
	Zap   Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Canal Canal `mapstructure:"canal" json:"canal" yaml:"canal"`
}
