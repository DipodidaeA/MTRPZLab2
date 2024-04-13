package lab2

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestCompute(t *testing.T) {
	ch := ComputeHandler{YourPath: t.TempDir()}
	pr := filepath.Join(ch.YourPath, "Text.md")
	os.WriteFile(pr, []byte("**Test**"), 0666) // створення файлу для читання у тимчасовій директорії
	for _, tc := range []struct {
		name string
		pkey string
		wkey string
	}{
		{name: "1", pkey: "Text.md", wkey: "Result.html"},
		{name: "2", pkey: "Text.md", wkey: ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := ch.Compute(tc.pkey, tc.wkey)
			if err != "" {
				t.Errorf("Помилка: %s", err)
			} else {
				fmt.Println("Виконання успішне")
			}
		})
	}
}
