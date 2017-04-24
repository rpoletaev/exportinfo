package exportinfo

import (
	"fmt"
	"regexp"
)

// ExportInfo структура, которая содержит информацию
// о заголовке документа экспорта и версию схемы данных
// по заголовку документа определяется имя коллекции в которой будет сохранен документа
// так же по заголовку документа определяется какой именно документ из структуры экспорта мы должны получить,
// т.к. документ экспорта включает в себя все типы документов, но каждый xml содержит только один документ
// из всех возможных описаных в структуре экспорта
type ExportInfo struct {
	Title   string
	Version string
}

var (
	findTagRegex, infoRegex *regexp.Regexp
)

func init() {
	findTagRegex, _ = regexp.Compile(`.+(<.+ schemeVersion=".+".*>)`)
	infoRegex, _ = regexp.Compile(`<([a-z]+\d?[:]{1})?([a-zA-Z]*\d*[a-zA-Z]*[\S]).*(\sschemeVersion=\"(.+?)\").*>`)
}

// GetExportInfo извлекает версию и тип документа
func GetExportInfo(xml string) (*ExportInfo, error) {
	tag := findTagRegex.FindStringSubmatch(xml)
	if len(tag) < 2 {
		return nil, fmt.Errorf("Неверный формат файла. Не удалось определить тег с версией и типом документа")
	}
	results := infoRegex.FindStringSubmatch(tag[1])
	if len(results) < 5 {
		return nil, fmt.Errorf("Неверный формат файла. Не удалось определить версию и номер документа")
	}

	return &ExportInfo{Title: results[2], Version: results[4]}, nil
}
