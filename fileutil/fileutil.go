package fileutil

import (
	"fmt"
	"os"
)

type IfileUtil interface {
	CreateOutputFile(prefix string, idx []int) (*os.File, error)
	WriteToByte(f *os.File, content []byte) error
	WriteToLine(f *os.File, content []byte) error
}

func CreateOutputFile(prefix string, idx []int) (*os.File, error) {
	asciiOffset := 97 // 'a'
	asciiSuffix := []string{string(rune(asciiOffset + idx[0])), string(rune(asciiOffset + idx[1]))}
	outputFileName := fmt.Sprintf("%s%s%s", prefix, asciiSuffix[1], asciiSuffix[0])

    outputFile, err := os.Create(outputFileName)
    if err != nil {
        return nil, err
    }
    return outputFile, nil
}

func WriteToByte(file *os.File, content []byte) error {
    if _, err := file.Write(content); err != nil {
        return err
    }
    return nil
}

func WriteToLine(file *os.File, content string) error {
	if _, err := file.WriteString(content); err != nil {
        return err
    }
    return nil
}
