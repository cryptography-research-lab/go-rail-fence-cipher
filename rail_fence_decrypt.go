package rail_fence_cipher

import (
	"errors"
	variable_parameter "github.com/golang-infrastructure/go-variable-parameter"
	"strings"
	"unicode/utf8"
)

// TODO
func Decrypt(encryptText string, options ...*Options) (string, error) {

	// 未传递参数的话则设置默认的参数
	options = variable_parameter.SetDefaultParamByFunc(options, func() *Options {
		return NewOptions()
	})

	// TODO 参数检查
	if options[0].PutEdgeDirection == options[0].TakeEdgeDirection {
		return "", errors.New("PutEdgeDirection can not equals TakeEdgeDirection")
	}

	rowCount := (utf8.RuneCountInString(encryptText) + options[0].Columns - 1) / options[0].Columns
	table := NewTable(rowCount, options[0].Columns)
	encrypttextConsumer := NewTextConsumer(encryptText, options[0].FillCharacter)

	// 按照取的方式放
	table.VisitByEdgeDirection(options[0].TakeEdgeDirection, func(table RailRenceTable, rowIndex, columnIndex int, character rune) {
		table[rowIndex][columnIndex] = encrypttextConsumer.Take()
	})

	// 按照放的方式取
	result := strings.Builder{}
	table.VisitByEdgeDirection(options[0].PutEdgeDirection, func(table RailRenceTable, rowIndex, columnIndex int, character rune) {
		result.WriteRune(character)
	})

	return result.String(), nil
}
