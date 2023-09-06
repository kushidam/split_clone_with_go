package fileutil

import (
	"os"
	"testing"
)



func TestCreateOutputFile(t *testing.T) {
	// テストデータ
	prefix := "x"
	idx := []int{0, 0}
	expectFileName := "xaa"

	// 関数を呼び出し、出力されるファイル名を検証
	outputFile, err := CreateOutputFile(prefix, idx)
	if err != nil {
		t.Fatalf("CreateOutputFile ERROR: %v", err)
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
	outputFile, err = CreateOutputFile(prefix, idx)
	if err != nil {
		t.Fatalf("CreateOutputFile ERROR: %v", err)
	}
	defer os.Remove(outputFile.Name())
	defer outputFile.Close()

	outputFileName = outputFile.Name()
	if outputFileName != expectFileName {
		t.Errorf("\nファイル名不一致\nExpect: %s, Result: %s",
			expectFileName, outputFileName)
	}
}

func TestWriteToByte(t *testing.T) {
	// テストデータ
	testContentData := []byte("Hello, World!")

	// テスト用の一時ファイルを作成
	file, err := os.CreateTemp("", "testfileutil")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())
	defer file.Close()

	// 関数を呼び出し
	err = WriteToByte(file, testContentData)

	// テストケース: エラーが発生しないことを検証
	if err != nil {
		t.Fatalf("WriteToByte failed: %v", err)
	}

	// ファイルから読み込んで内容を検証
	readContent := make([]byte, len(testContentData))
	_, err = file.ReadAt(readContent, 0)
	if err != nil {
		t.Fatalf("Failed to read from file: %v", err)
	}

	// テストケース: 書き込んだ内容が一致することを検証
	if string(readContent) != string(testContentData) {
		t.Errorf("Content mismatch: expected '%s', got '%s'", testContentData, readContent)
	}
}

func TestWriteToLine(t *testing.T) {
	// テストデータ
	testContentData := "Test Content Data"
	tmpFile, err := os.CreateTemp("", "test_output")
	if err != nil {
		t.Fatalf("CreateTemp ERROR: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// 関数を呼び出し、ファイルに書き込む
	err = WriteToLine(tmpFile, testContentData)
	if err != nil {
		t.Fatalf("writeOutputFile ERROR: %v", err)
	}
	// Readポインタを先頭
	tmpFile.Seek(0, 0)
	// ファイルを読み出して内容を確認
	bufSize := make([]byte, len(testContentData))
	_, err = tmpFile.Read(bufSize)
	if err != nil {
		t.Fatalf("tmpFile.Read ERROR: %v", err)
	}

	if string(bufSize) != testContentData {
		t.Errorf("\n書き込まれた内容不一致\nExpect: %s, Result: %s", testContentData, string(bufSize))
	}
}
