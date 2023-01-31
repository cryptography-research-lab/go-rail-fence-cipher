package rail_fence_cipher

import (
	"errors"
	variable_parameter "github.com/golang-infrastructure/go-variable-parameter"
	"strings"
	"unicode/utf8"
)

// Encrypt 对文本进行加密
func Encrypt(plaintext string, options ...*Options) (string, error) {

	// 未传递参数的话则设置默认的参数
	options = variable_parameter.SetDefaultParamByFunc(options, func() *Options {
		return NewOptions()
	})

	// TODO 参数检查
	if options[0].PutEdgeDirection == options[0].TakeEdgeDirection {
		return "", errors.New("PutEdgeDirection can not equals TakeEdgeDirection")
	}

	// 铺设字符
	plaintextConsumer := NewTextConsumer(plaintext, options[0].FillCharacter)
	rowCount := (utf8.RuneCountInString(plaintext) + options[0].Columns - 1) / options[0].Columns
	table := NewTable(rowCount, options[0].Columns)
	table.VisitByEdgeDirection(options[0].PutEdgeDirection, func(table RailRenceTable, rowIndex, columnIndex int, character rune) {
		table[rowIndex][columnIndex] = plaintextConsumer.Take()
	})

	// 收集字符
	result := strings.Builder{}
	table.VisitByEdgeDirection(options[0].TakeEdgeDirection, func(table RailRenceTable, rowIndex, columnIndex int, character rune) {
		result.WriteRune(character)
	})

	return result.String(), nil
}

// ------------------------------------------------ ---------------------------------------------------------------------
