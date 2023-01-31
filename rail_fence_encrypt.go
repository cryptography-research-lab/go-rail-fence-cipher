package rail_fence_cipher

import (
	"errors"
	variable_parameter "github.com/golang-infrastructure/go-variable-parameter"
	"strings"
	"unicode/utf8"
)

// ------------------------------------------------ ---------------------------------------------------------------------

// Encrypt 对明文进行加密
func Encrypt(plaintext string, options ...*Options) (string, error) {

	// 未传递参数的话则设置默认的参数
	options = variable_parameter.SetDefaultParamByFunc(options, func() *Options {
		return NewOptions()
	})

	// 参数检查
	if options[0].PutEdgeDirection == options[0].TakeEdgeDirection {
		return "", errors.New("PutEdgeDirection can not equals TakeEdgeDirection")
	}

	// 铺设字符
	plaintextConsumer := NewTextConsumer(plaintext, options[0].FillCharacter)
	rowCount := (utf8.RuneCountInString(plaintext) + options[0].Columns - 1) / options[0].Columns
	table := NewTable(rowCount, options[0].Columns)
	table.VisitByEdgeDirection(options[0].PutEdgeDirection, func(table Table, rowIndex, columnIndex int, character rune) {
		table[rowIndex][columnIndex] = plaintextConsumer.Take()
	})

	// 收集字符
	result := strings.Builder{}
	table.VisitByEdgeDirection(options[0].TakeEdgeDirection, func(table Table, rowIndex, columnIndex int, character rune) {
		result.WriteRune(character)
	})

	return result.String(), nil
}

// EncryptW W型的栅栏加密
func EncryptW(plaintext string, options ...*Options) (string, error) {
	// 未传递参数的话则设置默认的参数
	options = variable_parameter.SetDefaultParamByFunc(options, func() *Options {
		return NewOptions()
	})

	// 参数检查
	if options[0].Rows < 3 {
		return "", errors.New("rows min 3")
	}

	table := NewTable(options[0].Rows, utf8.RuneCountInString(plaintext))
	plaintextRuneSlice := []rune(plaintext)
	table.VisitByW(func(table Table, rowIndex, columnIndex int, character rune) {
		table[rowIndex][columnIndex] = plaintextRuneSlice[columnIndex]
	})

	//fmt.Println(table.String())

	result := strings.Builder{}
	table.VisitByEdgeDirection(EdgeDirectionLeftTop2Right, func(table Table, rowIndex, columnIndex int, character rune) {
		if character == 0 {
			return
		}
		result.WriteRune(character)
	})

	return result.String(), nil
}

// ------------------------------------------------ ---------------------------------------------------------------------
