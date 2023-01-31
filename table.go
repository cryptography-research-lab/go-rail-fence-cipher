package rail_fence_cipher

// RailRenceTable 用于表示一张栅栏表，封装了一些表的操作方法
type RailRenceTable [][]rune

// NewTable 创建一张指定大小的二维表格
// rowCount: 表的行数
// columnCount: 表的列数
func NewTable(rowCount, columnCount int) RailRenceTable {
	table := make([][]rune, rowCount)
	for rowIndex := range table {
		table[rowIndex] = make([]rune, columnCount)
	}
	return table
}

// RowCount 栅栏表的行数
func (x RailRenceTable) RowCount() int {
	return len(x)
}

// ColumnCount 栅栏表的列数
func (x RailRenceTable) ColumnCount() int {
	if x.RowCount() == 0 {
		return 0
	} else {
		return len(x[0])
	}
}

// VisitByEdgeDirection 根据给定的边的方向遍历栅栏表
// edgeDirection: 要遍历的方向
// visitFunc: 访问时对每个元素的访问方法，每个单元格都会调用一次此方法
func (x RailRenceTable) VisitByEdgeDirection(edgeDirection EdgeDirection, visitFunc func(table RailRenceTable, rowIndex, columnIndex int, character rune)) {

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
