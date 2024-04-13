package lab2

import (
	"fmt"
	"os"
	"path/filepath"
)

type ComputeHandler struct {
	YourPath string
}

func (ch *ComputeHandler) Compute(pkey, wkey string) string {
	err := ""
	res := ""

	if pkey != "" {
		pr := filepath.Join(ch.YourPath, pkey)
		data, merr1 := os.ReadFile(pr)
		if merr1 != nil {
			err = "Помилка читання файлу"
		} else {
			res, err = ChengeText(string(data))
			if err == "" {
				if wkey != "" {
					pw := filepath.Join(ch.YourPath, wkey)
					merr2 := os.WriteFile(pw, []byte(res), 0666)
					if merr2 != nil {
						err = "Помилка запису у файл"
					}
				} else {
					fmt.Println("Відповідь: ", res)
				}
			}
		}
	} else {
		err = "Невказано файлу"
	}
	return err
}
