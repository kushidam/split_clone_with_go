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

    outputFile, err := os.Create(outputFileName)
    if err != nil {
        return nil, err
    }
    return outputFile, nil
}


func writeOutputFile(file *os.File, content string) error {
	_, err := file.WriteString(content)
	return err
}

func writeLinesFile(writeUnit int, content string, idx []int, inputLine *int, outputString *string) error {
	if *inputLine >= writeUnit {
		outputFile, err := createOutputFile_l("x", idx)
		if err != nil {
			return err
		}
		defer outputFile.Close()

		if err := writeOutputFile(outputFile, *outputString); err != nil {
			return err
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
		if err := writeLinesFile(readUnit, outputString, idx, &inputLine, &outputString); err != nil {
			return err
		}
	}

	if inputLine > 0 {
		outputFile, err := createOutputFile_l("x", idx)
		if err != nil {
			return err
		}
		defer outputFile.Close()

		if err := writeOutputFile(outputFile, outputString); err != nil {
			return err
		}
	}

	return nil
}
