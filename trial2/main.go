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

	"github.com/naufalandika/designpattern/trial2/pb/accomplishment"
	"github.com/naufalandika/designpattern/trial2/pb/certificate"
	"github.com/naufalandika/designpattern/trial2/pb/exercise"
	userskillcourse "github.com/naufalandika/designpattern/trial2/user-skill-course"
	usersubmission "github.com/naufalandika/designpattern/trial2/user-submission"
)

var (
	uscSvc = userskillcourse.NewService(accomplishment.New())
	usSvc  = usersubmission.NewService(accomplishment.New(), certificate.New(), exercise.New())

	rgxGetQuestions, _   = regexp.Compile("getquestions/*/*")
	rgxSubmitPostTest, _ = regexp.Compile("submitposttest/*/*")
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
	ut, ct := parseReq(r)

	switch {
	case rgxGetQuestions.MatchString(r):
		return uscSvc.GetQuestions(context.Background(), &userskillcourse.GetQuestionsRequest{
			UserType:   ut,
			CourseType: ct,
		})
	case rgxSubmitPostTest.MatchString(r):
		return usSvc.SubmitPostTest(context.Background(), &usersubmission.SubmitPostTestRequest{
			UserType:   ut,
			CourseType: ct,
		})
	default:
		return nil, errors.New("unimplemented")
	}
}

func parseReq(r string) (string, string) {
	arr := strings.Split(r, "/")

	if len(arr) == 0 {
		return "", ""
	}

	return arr[1], arr[2]
}
