package rail_fence_cipher

import "strings"

// Table 用于表示一张栅栏表，封装了一些表的操作方法
type Table [][]rune

// NewTable 创建一张指定大小的二维表格作为栅栏
// rowCount: 表的行数
// columnCount: 表的列数
func NewTable(rowCount, columnCount int) Table {
	table := make([][]rune, rowCount)
	for rowIndex := range table {
		table[rowIndex] = make([]rune, columnCount)
	}
	return table
}

// RowCount 栅栏表的行数
func (x Table) RowCount() int {
	return len(x)
}

// ColumnCount 栅栏表的列数
func (x Table) ColumnCount() int {
	if x.RowCount() == 0 {
		return 0
	} else {
		return len(x[0])
	}
}

type VisitFunc func(table Table, rowIndex, columnIndex int, character rune)

// VisitByEdgeDirection 根据给定的边的方向遍历栅栏表
// edgeDirection: 要遍历的方向
// visitFunc: 访问时对每个元素的访问方法，每个单元格都会调用一次此方法
func (x Table) VisitByEdgeDirection(edgeDirection EdgeDirection, visitFunc VisitFunc) {

	switch edgeDirection {

	// EdgeDirectionLeftTop2Right 从左上往右
	case EdgeDirectionLeftTop2Right:
		for i := 0; i < x.RowCount(); i++ {
			for j := 0; j < x.ColumnCount(); j++ {
				visitFunc(x, i, j, x[i][j])
			}
		}
	// EdgeDirectionLeftTop2Bottom 从左上往下
	case EdgeDirectionLeftTop2Bottom:
		for i := 0; i < x.ColumnCount(); i++ {
			for j := 0; j < x.RowCount(); j++ {
				visitFunc(x, j, i, x[j][i])
			}
		}
	// EdgeDirectionLeftBottom2Top 从左下往上
	case EdgeDirectionLeftBottom2Top:
		for i := 0; i < x.ColumnCount(); i++ {
			for j := x.RowCount() - 1; j >= 0; j-- {
				visitFunc(x, j, i, x[j][i])
			}
		}
	// EdgeDirectionLeftBottom2Right 从左下往右
	case EdgeDirectionLeftBottom2Right:
		for i := x.RowCount() - 1; i >= 0; i-- {
			for j := 0; j < x.ColumnCount(); j++ {
				visitFunc(x, i, j, x[i][j])
			}
		}
	// DirectionRightTop2Left 从右上往左
	case EdgeDirectionRightTop2Left:
		for i := x.RowCount() - 1; i >= 0; i-- {
			for j := x.ColumnCount() - 1; j >= 0; j-- {
				visitFunc(x, i, j, x[i][j])
			}
		}
	// EdgeDirectionRightTop2Bottom 从右上往下
	case EdgeDirectionRightTop2Bottom:
		for i := x.ColumnCount() - 1; i >= 0; i-- {
			for j := 0; j < x.RowCount(); j++ {
				visitFunc(x, j, i, x[j][i])
			}
		}
	// EdgeDirectionRightBottom2Top 从右下往上
	case EdgeDirectionRightBottom2Top:
		for i := x.ColumnCount() - 1; i >= 0; i-- {
			for j := x.RowCount() - 1; j >= 0; j-- {
				visitFunc(x, j, i, x[j][i])
			}
		}
	// EdgeDirectionRightBottom2Left 从右下往左
	case EdgeDirectionRightBottom2Left:
		for i := x.RowCount() - 1; i >= 0; i-- {
			for j := x.ColumnCount() - 1; j >= 0; j-- {
				visitFunc(x, i, j, x[i][j])
			}
		}
	}
}

// VisitByW W型遍历表格
func (x Table) VisitByW(visitFunc VisitFunc) {
	step := 1
	rowIndex := 0
	for columnIndex := 0; columnIndex < len(x[0]); columnIndex++ {

		visitFunc(x, rowIndex, columnIndex, x[rowIndex][columnIndex])

		rowIndex += step

		// 触底回弹
		if rowIndex < 0 {
			rowIndex += 2
			step = 1
		} else if rowIndex >= len(x) {
			// 触顶回弹
			rowIndex -= 2
			step = -1
		}
	}
}

// 把加密使用的表格转为字符串返回，用于观察表格长啥样
// 返回数据样例：
//
//	 [
//		[ I, C, L, O, M ]
//		[ P, H, D, R, Z ]
//		[ U, V, F, Y, B ]
//		[ G, X, T, Q, E ]
//		[ S, N, K, W, A ]
//	]
func (x Table) String() string {
	sb := strings.Builder{}
	sb.WriteString("[\n")
	for _, line := range x {
		sb.WriteString("\t[ ")
		for index, cellValue := range line {
			// 2023-2-1 02:32:09 为了打印时美观，将其替换为可视化的字符
			if cellValue == 0 {
				cellValue = '.'
			}
			sb.WriteRune(cellValue)
			if index+1 != len(line) {
				sb.WriteString(",")
			}
			sb.WriteString(" ")
		}
		sb.WriteString("]\n")
	}
	sb.WriteString("]")
	return sb.String()
}
