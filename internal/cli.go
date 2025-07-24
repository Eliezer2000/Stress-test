package internal

import (
	"flag"
	"fmt"
	"os"
)

func ParseFlags() Config {
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 1, "Número total de requests")
	concurrency := flag.Int("concurrency", 1, "Numero de chamadas simultâneas")
	flag.Parse()

	if *url == "" {
		fmt.Println("A URL é obrigatória. Use --url para especificar.")
		os.Exit(1)
	}

	if *requests < 1 || *concurrency < 1 {
		fmt.Println("Requests e concurrency devem ser maiores que zero.")
		os.Exit(1)
	}
	
	return Config{
		URL: *url,
		Requests: *requests,
		Concurrency: *concurrency,
	}
}