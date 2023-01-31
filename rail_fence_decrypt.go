package rail_fence_cipher

import (
	"errors"
	variable_parameter "github.com/golang-infrastructure/go-variable-parameter"
	"strings"
	"unicode/utf8"
)

// ------------------------------------------------ ---------------------------------------------------------------------

func Decrypt(encryptText string, options ...*Options) (string, error) {

	// 未传递参数的话则设置默认的参数
	options = variable_parameter.SetDefaultParamByFunc(options, func() *Options {
		return NewOptions()
	})

	// 参数检查
	if options[0].PutEdgeDirection == options[0].TakeEdgeDirection {
		return "", errors.New("PutEdgeDirection can not equals TakeEdgeDirection")
	}

	rowCount := (utf8.RuneCountInString(encryptText) + options[0].Columns - 1) / options[0].Columns
	table := NewTable(rowCount, options[0].Columns)
	encrypttextConsumer := NewTextConsumer(encryptText, options[0].FillCharacter)

	// 按照取的方式放
	table.VisitByEdgeDirection(options[0].TakeEdgeDirection, func(table Table, rowIndex, columnIndex int, character rune) {
		table[rowIndex][columnIndex] = encrypttextConsumer.Take()
	})

	//fmt.Println(table.String())

	// 按照放的方式取
	result := strings.Builder{}
	table.VisitByEdgeDirection(options[0].PutEdgeDirection, func(table Table, rowIndex, columnIndex int, character rune) {
		result.WriteRune(character)
	})

	return result.String(), nil
}

// ------------------------------------------------ ---------------------------------------------------------------------

// DecryptW W型的栅栏解密
func DecryptW(ciphertext string, options ...*Options) (string, error) {

	// 未传递参数的话则设置默认的参数
	options = variable_parameter.SetDefaultParamByFunc(options, func() *Options {
		return NewOptions()
	})

	// 参数检查
	if options[0].Rows < 3 {
		return "", errors.New("rows min 3")
	}

	table := NewTable(options[0].Rows, utf8.RuneCountInString(ciphertext))

	// 标记W路径上的字符
	table.VisitByW(func(table Table, rowIndex, columnIndex int, character rune) {
		table[rowIndex][columnIndex] = 1
	})

	// 遍历设置W路径上的字符
	ciphertextConsumer := NewTextConsumer(ciphertext, '.')
	table.VisitByEdgeDirection(EdgeDirectionLeftTop2Right, func(table Table, rowIndex, columnIndex int, character rune) {
		if character == 1 {
			table[rowIndex][columnIndex] = ciphertextConsumer.Take()
		}
	})

	//fmt.Println(table.String())

	// 收集W路径上的字符
	result := strings.Builder{}
	table.VisitByW(func(table Table, rowIndex, columnIndex int, character rune) {
		result.WriteRune(character)
	})

	return result.String(), nil
}

// ------------------------------------------------ ---------------------------------------------------------------------
