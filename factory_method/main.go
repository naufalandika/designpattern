package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/naufalandika/designpattern/factory_method/service"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	go func() {
		for {
			t, _ := reader.ReadString('\n')
			t = strings.TrimSuffix(t, "\n")
			svc := service.New(t)

			svc.GetQuestions()
			fmt.Println()
		}
	}()

	sigchnl := make(chan os.Signal, 1)
	signal.Notify(sigchnl, os.Interrupt)
	<-sigchnl
}
