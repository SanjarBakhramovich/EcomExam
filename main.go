package main

import (
	"EcomExam/usecases"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Please provide order numbers as command line arguments.")
	}

	orderNumbers := strings.Split(args[0], ",")
	assemblyPage(orderNumbers)
}

func assemblyPage(orderNumbers []string) {
	usecase := usecases.NewAssemblyPageUsecase()
	output := usecase.AssemblePage(orderNumbers)
	log.Println(output)
}
