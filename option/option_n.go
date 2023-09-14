package option

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func Option_n(chunks string, file *os.File) error {
	inOption := strings.Split(chunks, "/")
	argsLen := len(inOption)
	if 1 > argsLen || argsLen > 3 {
		return fmt.Errorf("[ERROR] 引数無効")
	}

	if errVarid := validateArgs(inOption[1:]); errVarid != nil {
		return errVarid
	}

	if inOption[0] == "r" {
		return fmt.Errorf("option 'r' 未実装")
	}
	if inOption[0] == "l" {
		return fmt.Errorf("option 'l' 未実装")
	}

	intOption, errAtoi := strconv.Atoi(inOption[0])
	if errAtoi != nil {
		return errAtoi
	}

	if argsLen == 1 {
		return option_n_default(intOption, file)
	} else {
		return option_n_k(intOption)
	}
}

func validateArgs(args []string) error {
	for i := 0; i < len(args); i++ {
		arg := args[i]
		_, errValidate := strconv.Atoi(arg)
		if errValidate != nil {
			return errValidate
		}
	}
	return nil
}

func option_n_default(splitNum int, file *os.File) error {
	fmt.Println("option_n_default")
	// xaa ~ xzz まで対応
	if splitNum < 1 || splitNum > 676 {
		return fmt.Errorf("[ERROR] 分割数無効")
	}

	var prefix string = "x"
	idx := []int{0, 0}
	asciiInt := 97 // `a`
	fileInfo, _ := file.Stat()
	splitSize := make([]byte, int64(math.Ceil(float64(fileInfo.Size())/float64(splitNum))))

	for {
		n, errIORead := file.Read(splitSize)
		if errIORead != nil && errIORead != io.EOF {
			return errIORead
		}
		if n == 0 {
			break
		}

		outputFileName := fmt.Sprintf("%s%s%s", prefix, string(rune(asciiInt+idx[1])), string(rune(asciiInt+idx[0])))
		outputFile, errIOCreate := os.Create(outputFileName)
		if errIOCreate != nil {
			return errIOCreate
		}
		defer outputFile.Close()

		_, errIOWrite := outputFile.Write(splitSize[:n])
		if errIOWrite != nil {
			return errIOWrite
		}

		idx[0]++
		if idx[0] == 26 {
			idx[0] = 0
			idx[1]++
		}
	}

	return nil
}

func option_n_k(k int) error {
	return nil
}
