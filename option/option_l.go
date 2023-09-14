package option

import (
	"bufio"
	"fmt"
	"os"
)

func createOutputFile_l(prefix string, idx []int) (*os.File, error) {
	asciiOffset := 97 // 'a'
	asciiSuffix := []string{string(rune(asciiOffset + idx[0])), string(rune(asciiOffset + idx[1]))}
	outputFileName := fmt.Sprintf("%s%s%s", prefix, asciiSuffix[1], asciiSuffix[0])

    outputFile, errOsCreate := os.Create(outputFileName)
    if errOsCreate != nil {
        return nil, errOsCreate
    }
    return outputFile, nil
}


func writeOutputFile(file *os.File, content string) error {
	_, err := file.WriteString(content)
	return err
}

func writeLinesFile(writeUnit int, content string, idx []int, inputLine *int, outputString *string) error {
	if *inputLine >= writeUnit {
		outputFile, errIoCreate := createOutputFile_l("x", idx)
		if errIoCreate != nil {
			return errIoCreate
		}
		defer outputFile.Close()

		if errIoWrite := writeOutputFile(outputFile, *outputString); errIoWrite != nil {
			return errIoWrite
		}
		*inputLine = 0
		*outputString = ""
		idx[0]++
		if idx[0] == 26 {
			idx[0] = 0
			idx[1]++
		}
	}
	return nil
}

func Option_l(readUnit int, file *os.File) error {
	fmt.Println("Option_l", readUnit)

	scanner := bufio.NewScanner(file)
	idx := []int{0, 0}
	inputLine := 0
	var outputString string

	for scanner.Scan() {
		line := scanner.Text()
		outputString += line + "\n"
		inputLine++
		if errIoLine := writeLinesFile(readUnit, outputString, idx, &inputLine, &outputString); errIoLine != nil {
			return errIoLine
		}
	}

	if inputLine > 0 {
		outputFile, errIoCreate := createOutputFile_l("x", idx)
		if errIoCreate != nil {
			return errIoCreate
		}
		defer outputFile.Close()

		if errIoWrite := writeOutputFile(outputFile, outputString); errIoWrite != nil {
			return errIoWrite
		}
	}

	return nil
}
