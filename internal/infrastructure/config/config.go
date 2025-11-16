package config

import (
    "github.com/spf13/viper"
)


type ServerCfg struct {
    Host string `mapstructure:"host"`
    Port int    `mapstructure:"port"`
}

type DBCfg struct {
    Driver string `mapstructure:"driver"`
    DSN    string `mapstructure:"dsn"`
}

type JWTCfg struct {
    Secret     string `mapstructure:"secret"`
    TTLMinutes int    `mapstructure:"ttl_minutes"`
}

type Config struct {
    Server ServerCfg `mapstructure:"server"`
    DB     DBCfg     `mapstructure:"db"`
    JWT    JWTCfg    `mapstructure:"jwt"`
}

func Load() (*Config, error) {
    v := viper.New()
    v.SetConfigName("config")
    v.AddConfigPath("configs")
    v.SetConfigType("yaml")
    if err := v.ReadInConfig(); err != nil {
        return nil, err
    }
    var cfg Config
    if err := v.Unmarshal(&cfg); err != nil {
        return nil, err
    }
    return &cfg, nil
}
