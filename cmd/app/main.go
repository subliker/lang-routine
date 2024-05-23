package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/subliker/LangRoutine/internal/asyncdetection"
	"github.com/subliker/LangRoutine/internal/detection"
	"github.com/subliker/LangRoutine/internal/syncdetection"
)

func main() {
	// флаг синхронного режима работы -sync
	sync := flag.Bool("sync", false, "choose sync mode")
	// флаг пути к папке с файлами для чтения и анализа -path
	path := flag.String("path", "assets", "folder path")
	// флаг пути к файлу для чтения и анализа -fpath
	fpath := flag.String("fpath", "", "file path")

	// чтение флагов
	flag.Parse()

	if *fpath != "" {
		// получение языка из текста файла
		dr := detection.DetectFile(*fpath)
		if dr.Err != nil {
			fmt.Printf("Error detecting file: %s", dr.Err)
			return
		}

		// вывод найденного языка
		fmt.Printf("Detected language: %s\n", dr.Lang)
		return
	}

	// приведение к общему формату
	pathS := filepath.FromSlash(*path)

	// синхронный режим работы
	if *sync {
		syncdetection.SyncDetect(pathS)
		return
	}

	// асинхронный режим работы
	asyncdetection.AsyncDetect(pathS)
}
