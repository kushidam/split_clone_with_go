package option

import (
	"fmt"
	"io"
	"os"
)
func createOutputFile_b(prefix string, idx []int) (*os.File, error) {
	asciiOffset := 97 // 'a'
	asciiSuffix := []string{string(rune(asciiOffset + idx[0])), string(rune(asciiOffset + idx[1]))}
	outputFileName := fmt.Sprintf("%s%s%s", prefix, asciiSuffix[1], asciiSuffix[0])

    outputFile, err := os.Create(outputFileName)
    if err != nil {
        return nil, err
    }
    return outputFile, nil
}

func writeToFile(file *os.File, content []byte) error {
    _, err := file.Write(content)
    return err
}

func writeFinalFile(fileName string, content []byte) error {
    finalFile, err := os.Create(fileName)
    if err != nil {
        return err
    }
    defer finalFile.Close()

    if _, err := finalFile.Write(content); err != nil {
        return err
    }
    return nil
}

func Option_b(bufSize []byte, file *os.File) error {
    var prefix string = "x"
    idx := []int{0, 0}

    fileInfo, err := file.Stat()
    if err != nil {
        return err
    }
    remainSize := fileInfo.Size()

    for {
        n, err := file.Read(bufSize)
        if err != nil && err != io.EOF {
            return err
        }
        if n == 0 {
            break
        }

        outputFile, err := createOutputFile_b(prefix, idx)
        if err != nil {
            return err
        }
        defer outputFile.Close()

        if err := writeToFile(outputFile, bufSize[:n]); err != nil {
            return err
        }

        remainSize -= int64(n)

        idx[0]++
        if idx[0] == 26 {
            idx[0] = 0
            idx[1]++
        }

        if idx[0] == 25 && idx[1] == 25 && remainSize > 0 {
			finalFileName := "xzz" // 最終ファイル名

			lastSize := make([]byte, remainSize)
			n, err := file.Read(lastSize)
			if err != nil && err != io.EOF {
				return err
			}
			if n == 0 {
				break
			}
			if err := writeFinalFile(finalFileName, lastSize[:n]); err != nil {
				return err
			}
		}
	}
	return nil
}