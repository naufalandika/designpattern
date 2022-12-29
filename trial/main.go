package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	go func() {
		for {
			t, _ := reader.ReadString('\n')
			t = strings.TrimSuffix(t, "\n")
			fmt.Println(t)

			fmt.Println()
		}
	}()

	sigchnl := make(chan os.Signal, 1)
	signal.Notify(sigchnl, os.Interrupt)
	<-sigchnl
}
