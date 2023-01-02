package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"

	userskillcourse "github.com/naufalandika/designpattern/trial/user-skill-course"
	usersubmission "github.com/naufalandika/designpattern/trial/user-submission"
)

var (
	uscSvc = userskillcourse.NewService()
	usSvc  = usersubmission.NewService()

	rgxGetQuestions, _   = regexp.Compile("getquestions/*")
	rgxSubmitPostTest, _ = regexp.Compile("submit/posttest")
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	go func() {
		for {
			r, _ := reader.ReadString('\n')
			r = strings.TrimSuffix(r, "\n")
			resp, err := forwardReq(r)
			if err != nil {
				fmt.Println("error: ", err.Error())
			} else {
				fmt.Println("success: ", resp)
			}

			fmt.Println()
		}
	}()

	sigchnl := make(chan os.Signal, 1)
	signal.Notify(sigchnl, os.Interrupt)
	<-sigchnl
}

func forwardReq(r string) (interface{}, error) {
	switch {
	case rgxGetQuestions.MatchString(r):
		return uscSvc.GetQuestions(context.Background(), &userskillcourse.GetQuestionsRequest{})
	case rgxSubmitPostTest.MatchString(r):
		return usSvc.SubmitPostTest(context.Background(), &usersubmission.SubmitPostTestRequest{})
	default:
		return nil, errors.New("unimplemented")
	}
}
