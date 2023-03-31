package util

import "github.com/spf13/viper"

type Config struct{
    DBDriver string `mapstructure:"DB_DRIVER"`
    DBSource string `mapstructure:"DB_SOURCE"`
}


func LoadConfig(path string)(err error,config Config ){
    viper.AddConfigPath(path)
    viper.SetConfigType("env")
    viper.SetConfigFile(".env")

    viper.AutomaticEnv();
    err=viper.ReadInConfig()

    if err!=nil{
        return
    }

    err= viper.Unmarshal(&config)

    return err,config
}
