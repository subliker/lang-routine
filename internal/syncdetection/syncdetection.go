package syncdetection

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/subliker/LangRoutine/internal/detection"
)

func SyncDetect(path string) {
	// обход файлов в папке
	files, err := os.ReadDir(path)
	// ошибка: директория не найдено
	if os.IsNotExist(err) {
		fmt.Println("Folder", path, "not found")
		return
	} else if err != nil {
		// ошибка: обход
		fmt.Println("Folder", path, "can't be read")
		return
	}

	for _, f := range files {
		// если найдена папка
		if f.IsDir() {
			continue
		}

		// получение языка из текста файла
		dr := detection.DetectFile(filepath.Join(path, f.Name()))
		if dr.Err != nil {
			fmt.Printf("Error detecting file: %s\n", dr.Err)
			return
		}

		// вывод найденного языка
		fmt.Printf("%s detected language: %s\n", f.Name(), dr.Lang)
	}
}
