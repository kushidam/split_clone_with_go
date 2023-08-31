package option

import (
	"bufio"
	"os"
	"testing"
)

func TestCreateOutputFile_l(t *testing.T) {
	// テストデータ
	prefix := "x"
	idx := []int{0, 0}
	expectFileName := "xaa"

	// 関数を呼び出し、出力されるファイル名を検証
	outputFile, err := createOutputFile_l(prefix, idx)
	if err != nil {
		t.Fatalf("createOutputFile ERROR: %v", err)
	}
	defer os.Remove(outputFile.Name())
	defer outputFile.Close()

	outputFileName := outputFile.Name()
	if outputFileName != expectFileName {
		t.Errorf("\nファイル名不一致\nExpect: %s, Result: %s",
			expectFileName, outputFileName)
	}
	
	idx = []int{25, 25}
	expectFileName = "xzz"

	// 関数を呼び出し、出力されるファイル名を検証
	outputFile, err = createOutputFile_l(prefix, idx)
	if err != nil {
		t.Fatalf("createOutputFile ERROR: %v", err)
	}
	defer os.Remove(outputFile.Name())
	defer outputFile.Close()

	outputFileName = outputFile.Name()
	if outputFileName != expectFileName {
		t.Errorf("\nファイル名不一致\nExpect: %s, Result: %s",
			expectFileName, outputFileName)
	}
}

func TestWriteOutputFile(t *testing.T) {
	// テストデータ
	tsetContentData := "Test Content Data"
	tmpFile, err := os.CreateTemp("", "test_output")
	if err != nil {
		t.Fatalf("CreateTemp ERROR: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// 関数を呼び出し、ファイルに書き込む
	err = writeOutputFile(tmpFile, tsetContentData)
	if err != nil {
		t.Fatalf("writeOutputFile ERROR: %v", err)
	}
	// Readポインタを先頭
	tmpFile.Seek(0, 0)
	// ファイルを読み出して内容を確認
	bufSize := make([]byte, len(tsetContentData))
	_, err = tmpFile.Read(bufSize)
	if err != nil {
		t.Fatalf("tmpFile.Read ERROR: %v", err)
	}

	if string(bufSize) != tsetContentData {
		t.Errorf("\n書き込まれた内容不一致\nExpect: %s, Result: %s",
			tsetContentData, string(bufSize))
	}
}

func TestWriteLinesFile(t *testing.T) {
	// テストデータ
	writeUnit := 2
	contentTestDate := "contentTestDate1\ncontentTestDate2\ncontentTestDate3\n"
	contentExpectDate := []string {"contentTestDate1\n", "contentTestDate2\n", "contentTestDate3\n"}
	idx := []int{0, 0}
	inputLine := 0
	outputString := ""

	// 関数を呼び出し、ファイルに書き込む
	err := writeLinesFile(writeUnit, contentTestDate, idx, &inputLine, &outputString)
	if err != nil {
		t.Fatalf("writeLinesFile ERROR: %v", err)
	}

	// ファイルを読み出して内容を確認
	tmpFile, err := os.CreateTemp("", "xaa")
	if err != nil {
		t.Fatalf("Open ERROR: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	scanner := bufio.NewScanner(tmpFile)
	for scanner.Scan() {
		line := scanner.Text()
		
		if line == contentExpectDate[inputLine] {
			t.Errorf("\n書き込まれた内容不一致\nExpect: %s, Result: %s",
				contentTestDate, line)
		}
		inputLine++
	}

	// ファイルを読み出して内容を確認
	tmpFile, err = os.CreateTemp("", "xab")
	if err != nil {
		t.Fatalf("Open ERROR: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	scanner = bufio.NewScanner(tmpFile)
	for scanner.Scan() {
		line := scanner.Text()
		
		if line == contentExpectDate[inputLine] {
			t.Errorf("\n書き込まれた内容不一致\nExpect: %s, Result: %s",
				contentTestDate, line)
		}
		inputLine++
	}
}