package main

import (
	"flag"
	"fmt"
	"os"
	"split_clone_with_go/option"
)

func main() {
	var (
		splitSize = flag.Int("b", 0, "--bytes=SIZE ファイルを分割するバイト単位のサイズ")
		lines     = flag.Int("l", 1000, "--lines=NUMBER 出力ファイルごとに行数/レコード数を指定")
		chunks    = flag.String("n", "", "--number=CHUNKS CHUNKS出力ファイルを生成")
	)

	// コマンドライン引数を解析する
	flag.Parse()

	// コマンドライン引数からファイル名を取得する
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Error: ファイル名が指定されていません。")
		fmt.Println("split [options] FILENAME")
		os.Exit(1)
	}
	filename := args[len(args)-1]

	// ファイル情報を取得
	_, err := os.Stat(filename)
	// エラーが nil でない場合はファイルが存在しないと判定(正しくない入力)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File", filename, ":が存在しない")
		} else {
			fmt.Println("Error:", err)
		}
		return
	}

	// ファイルを開く
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("[Error]file open:", err)
		return
	}
	defer file.Close()

	var resultErr error
	if *splitSize > 0 {
		resultErr = option.Option_b(make([]byte, *splitSize), file)
	} else if *lines > 0 {
		resultErr = option.Option_l(*lines, file)
	} else if len(*chunks) > 0 {
		resultErr = option.Option_n(*chunks, file)
	} else {
		fmt.Println("Error: 処理対象のオプションが指定されていません。")
	}

	if resultErr != nil {
		fmt.Println("Error:", resultErr)
	} else {
		fmt.Println("Complete")
	}
}
