package detection

import (
	"os"
	"path/filepath"

	"github.com/pemistahl/lingua-go"
)

// структура данных о результатах анализа
type DetectResult struct {
	Fpath string
	Lang  lingua.Language
	Err   error
}

func DetectFile(fpath string) DetectResult {
	dr := DetectResult{
		Fpath: fpath,
	}

	// создание объекта определителя языка с настройками
	detector := lingua.NewLanguageDetectorBuilder().
		FromAllLanguages().
		WithPreloadedLanguageModels().
		Build()

	// открытие файла по пути флага fpath и чтение
	text, err := os.ReadFile(filepath.FromSlash(fpath))
	// ошибка: файл не найден
	if os.IsNotExist(err) {
		dr.Err = err
		return dr
	} else if err != nil {
		// ошибка: чтение
		dr.Err = err
		return dr
	}

	// определение языка в строке
	lang, exist := detector.DetectLanguageOf(string(text))
	dr.Lang = lang
	// язык не может быть определен
	if !exist {
		dr.Err = err
		return dr
	}
	return dr
}
