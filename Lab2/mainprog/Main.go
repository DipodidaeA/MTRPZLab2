package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	lab2 "github.com/DipodidaeA/MTRPZLab2"
)

var (
	inputFile  = flag.Bool("f", false, "Вказування файлу з текстом")
	outputFile = flag.Bool("o", false, "Запис тексту в файл")
)

func main() {
	flag.Parse()

	var pkey, wkey string = "", ""

	if *inputFile {
		pkey = flag.Arg(0)
	}
	if *outputFile {
		wkey = flag.Arg(1)
	}

	tempDirPath, err3 := ioutil.TempDir("", "Run")
	if err3 != nil {
		fmt.Println("Помилка створення тимчасової директорії")
	} else {
		database := lab2.ComputeHandler{YourPath: tempDirPath}
		psw := filepath.Join(database.YourPath, "Text.md")
		os.WriteFile(psw, []byte("**Text** _Text_ `Text` ```T*e_x`t```"), 0666)

		err := database.Compute(pkey, wkey)
		if err != "" {
			fmt.Println("Помилка: " + err)
		} else {
			fmt.Println("Перетворення тексту успішне")
		}

		if wkey != "" {
			prw := filepath.Join(database.YourPath, wkey)
			data1, merr1 := os.ReadFile(prw)
			if merr1 != nil {
				fmt.Println("Помилка читання файлу")
			} else {
				fmt.Println("Текст у новому файлі: ", string(data1))
			}
		}
	}
	defer os.RemoveAll(tempDirPath)
}
