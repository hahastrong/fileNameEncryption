package fileNameEncryption

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//ValidateFileName filename format: xxxx20060102_XYZ.ext
func ValidateFileName(fileName string) (bool, error) {
	dotIndex := strings.LastIndex(fileName, ".")
	if dotIndex < 0 {
		return false, errors.New("failed to find ext start index")
	}

	if dotIndex < 7 {
		return false, errors.New("file name is too short to contain yyyymmdd_XYZ")
	}

	validCode := fileName[dotIndex-3 : dotIndex]
	fileDay := fileName[dotIndex-6 : dotIndex-4]
	fileMonth := fileName[dotIndex-8 : dotIndex-6]

	offset, err := generateOffset(fileMonth, fileDay)
	if err != nil {
		return false, err
	}

	sourceCode, _ := GenerateValidCode(offset)
	if sourceCode != validCode {
		return false, nil
	}

	return true, nil
}

//ValidateFileName filename format: xxxx20060102_XYZ.ext
func ValidateLastDirFileName(fileName string) (string, bool, error) {
	dotIndex := strings.LastIndex(fileName, "/")
	if dotIndex < 0 {
		return "", false, errors.New("failed to find ext start index")
	}

	if dotIndex < 7 {
		return "", false, errors.New("file name is too short to contain yyyymmdd_XYZ")
	}

	validCode := fileName[dotIndex-3 : dotIndex]
	fileDay := fileName[dotIndex-6 : dotIndex-4]
	fileMonth := fileName[dotIndex-8 : dotIndex-6]

	offset, err := generateOffset(fileMonth, fileDay)
	if err != nil {
		return "", false, err
	}

	sourceCode, lastSourceCode := GenerateValidCode(offset)
	if sourceCode != validCode && lastSourceCode != validCode {
		return "", false, nil
	}

	rawPath := fileName[:dotIndex-4] + fileName[dotIndex:]
	return rawPath, true, nil
}

func generateOffset(month, day string) (int, error) {
	m, err := strconv.Atoi(month)
	if err != nil {
		return 0, err
	}
	if m <=0 || m > 12 {
		return 0, errors.New("the month is invalid")
	}

	d, err := strconv.Atoi(day)
	if err != nil {
		return 0, err
	}
	if d <= 0 || d > 31 {
		return 0, errors.New("the day is invalid")
	}

	return m * d, nil
}

func GenerateValidCode(offset int) (string, string) {
	_, month, day := time.Now().Date()
	hour := time.Now().Hour()

	m := (int(month) + offset) % 27 + 'A'
	d := (day + offset + 1) % 27 + 'A'
	h := (hour + offset + 2) % 27 + 'A'

	offset = offset * ((h + 24 - 1) % 24)

	lastm := (int(month) + offset) % 27 + 'A'
	lastd := (day + offset* + 1) % 27 + 'A'
	lasth := (hour + offset + 2) % 27 + 'A'

	return fmt.Sprintf("%c%c%c", m, d, h), fmt.Sprintf("%c%c%c", lastm, lastd, lasth)
}

func GenerateFileName(fileName string) (string, error) {
	var resFileName string
	dotIndex := strings.LastIndex(fileName, ".")
	if dotIndex < 0 {
		return resFileName, errors.New("failed to find ext start index")
	}

	if dotIndex < 7 {
		return resFileName, errors.New("file name is too short to contain yyyymmdd_XYZ")
	}

	fileDay := fileName[dotIndex-2 : dotIndex]
	fileMonth := fileName[dotIndex-4 : dotIndex-2]

	offset, err := generateOffset(fileMonth, fileDay)
	if err != nil {
		return resFileName, err
	}

	sourceCode, _ := GenerateValidCode(offset)
	resFileName = fmt.Sprintf("%s_%s%s", fileName[:dotIndex], sourceCode, fileName[dotIndex:])

	return resFileName, nil
}

func GenerateLastDirFileName(fileName string) (string, error) {
	var resFileName string
	dotIndex := strings.LastIndex(fileName, "/")
	if dotIndex < 0 {
		return resFileName, errors.New("failed to find ext start index")
	}

	if dotIndex < 7 {
		return resFileName, errors.New("file name is too short to contain yyyymmdd_XYZ")
	}

	fileDay := fileName[dotIndex-2 : dotIndex]
	fileMonth := fileName[dotIndex-4 : dotIndex-2]

	offset, err := generateOffset(fileMonth, fileDay)
	if err != nil {
		return resFileName, err
	}

	sourceCode, _ := GenerateValidCode(offset)
	resFileName = fmt.Sprintf("%s_%s%s", fileName[:dotIndex], sourceCode, fileName[dotIndex:])

	return resFileName, nil
}

