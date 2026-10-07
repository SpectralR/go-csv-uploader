package jobs

import (
	"fmt"
	"log"
	"net/mail"
	"regexp"
	"strconv"
	"time"
)

func Validate(c chan [][]string) {
	fmt.Println("validating")
	channel := <-c
	var newChanData [][]string
	for line, channelData := range channel {
		_, err := validateLine(channelData)
		if err != nil {
			log.Printf("error line %d : %s \n", line, err.Error())
		} else {
			newChanData = append(newChanData, channelData)
		}
	}

	c <- newChanData
}

func validateLine(line []string) (bool, error) {
	for index, data := range line {
		if !validateMalicious(data) {
			return false, fmt.Errorf("file is not valid : %s \n", data)
		}

		switch index {
		case 7:
			if !validateEmail(data) {
				return false, fmt.Errorf("email is not valid : %s", data)
			}
		case 0:
			if !validateNumeric(data) {
				return false, fmt.Errorf("id is not valid : %s", data)
			}
		case 8:
			if !validateDate(data) {
				return false, fmt.Errorf("date is not valid : %s", data)
			}
		case 9:
			regex := `^https?:\/\/`
			if !validateRegex(data, regex) {
				return false, fmt.Errorf("website is not valid : %s", data)
			}
		}

	}

	return true, nil
}

func validateDate(date string) bool {
	_, err := time.Parse(time.DateOnly, date)
	return err == nil
}

func validateNumeric(data string) bool {
	_, err := strconv.Atoi(data)
	return err == nil
}

func validateEmail(data string) bool {
	_, err := mail.ParseAddress(data)
	return err == nil
}

func validateRegex(data string, pattern string) bool {
	compiled := regexp.MustCompile(pattern)
	return compiled.MatchString(data)
}

func validateMalicious(data string) bool {
	regex := regexp.MustCompile(`eval\(`)
	return !regex.MatchString(data)
}
