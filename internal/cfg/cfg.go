package cfg

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/go-pg/pg"
)

type Cfg struct { 
	Port   string
	DbName string
	DbUser string
	DbPass string
	DbHost string
	DbPort string
}

func LoadAndStoreConfig() Cfg {
	v := viper.New() 
	v.SetEnvPrefix("SERV") 
	v.SetDefault("PORT", "8080")
	v.SetDefault("DBUSER", "kirill")
	v.SetDefault("DBPASS", "postgres")
	v.SetDefault("DBHOST", "localhost")
	v.SetDefault("DBPORT", "5433")
	v.SetDefault("DBNAME", "test")
	v.AutomaticEnv()

	var cfg Cfg

	err := v.Unmarshal(&cfg)
	if err != nil {
		fmt.Println(err)
	}
	return cfg
}

func (cfg *Cfg) GetDBString() pg.Options { 
	return pg.Options{
		User:     cfg.DbUser,
		Password: cfg.DbPass,
		Database: cfg.DbName,
		Addr: "localhost:5433",
	}
	// return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
}