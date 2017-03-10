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

func getExportInfo(xmlContent string) (*ExportInfo, error) {
	r, err := regexp.Compile(`<(.+:)*([a-zA-Z]*\d*[a-zA-Z]*) schemeVersion=\"(.+?)\".*>`)
	if err != nil {
		return nil, err
	}

	results := r.FindStringSubmatch(xmlContent)
	if len(results) < 4 {
		return nil, fmt.Errorf("Wrong format export file, could not determine version and title")
	}

	return &ExportInfo{Title: results[2], Version: results[3]}, nil
}
