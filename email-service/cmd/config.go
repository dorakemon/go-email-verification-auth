package main

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type config struct {
	FromMail  string `env:"FROM_MAIL,required"`
	GmailPass string `env:"GMAIL_PASS,required"`
}

func loadEnv() (*config, error) {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
		return nil, err
	}
	return &cfg, nil
}

// func main() {
// 	cfg := config{}
// 	if err := env.Parse(&cfg); err != nil {
// 		fmt.Printf("%+v\n", err)
// 	}

// 	fmt.Printf("%+v\n", cfg)
// }
