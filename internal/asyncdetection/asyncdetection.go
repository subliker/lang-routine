package asyncdetection

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/subliker/LangRoutine/internal/detection"
)

func asyncDetectFile(wg *sync.WaitGroup, detectChan chan detection.DetectResult, fpath string) {
	// создание структры для результата анализа
	dr := detection.DetectResult{
		Fpath: fpath,
	}

	// получение языка из текста файла
	drf := detection.DetectFile(fpath)
	if drf.Err != nil {
		// Добавление ошибки в структуру результата и отправка результата в канал
		dr.Err = drf.Err
		detectChan <- dr
		wg.Done()
		return
	}

	// Добавление языка в структуру результата и отправка результата в канал
	dr.Lang = drf.Lang
	detectChan <- dr
	wg.Done()
}

func AsyncDetect(path string) {
	// создание канала для передачи данных анализа
	detectChan := make(chan detection.DetectResult)

	var wg sync.WaitGroup

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

		// создание горутины для анализа файла
		go asyncDetectFile(&wg, detectChan, filepath.Join(path, f.Name()))
		// добавление единицы в счетчик
		wg.Add(1)
	}

	// функция для закрытия канала
	go func() {
		wg.Wait()
		close(detectChan)
	}()

	for result := range detectChan {
		// обработка ошибок анализа
		if result.Err != nil {
			fmt.Printf("Detecting file error: %s\n", result.Err)

		} else {
			// вывод результата
			fmt.Printf("%s detected language: %s\n", result.Fpath, result.Lang)
		}
	}
}
