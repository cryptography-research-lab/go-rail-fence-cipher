package rail_fence_cipher

// ------------------------------------------------ ---------------------------------------------------------------------

// EdgeDirection 用于表示栅栏的起点和方向
type EdgeDirection int

const (

	// EdgeDirectionLeftTop2Right 从左上往右
	EdgeDirectionLeftTop2Right EdgeDirection = iota

	// EdgeDirectionLeftTop2Bottom 从左上往下
	EdgeDirectionLeftTop2Bottom

	// EdgeDirectionLeftBottom2Top 从左下往上
	EdgeDirectionLeftBottom2Top

	// EdgeDirectionLeftBottom2Right 从左下往右
	EdgeDirectionLeftBottom2Right

	// EdgeDirectionRightTop2Left 从右上往左
	EdgeDirectionRightTop2Left

	// EdgeDirectionRightTop2Bottom 从右上往下
	EdgeDirectionRightTop2Bottom

	// EdgeDirectionRightBottom2Top 从右下往上
	EdgeDirectionRightBottom2Top

	// EdgeDirectionRightBottom2Left 从右下往左
	EdgeDirectionRightBottom2Left
)

// ------------------------------------------------ ---------------------------------------------------------------------

// Options 加密或者解密时通用的一些选项
type Options struct {

	// 往栅栏中放入字符的时候的方向，不能与TakeEdgeDirection相同
	PutEdgeDirection EdgeDirection

	// 从栅栏中拿字符的时候的方向，不能与TakeEdgeDirection相同
	TakeEdgeDirection EdgeDirection

	// 栅栏是几列
	Columns int
	// 栅栏是几行
	Rows int

	// 长度不够的时候拿什么字符来填充对齐
	FillCharacter rune
}

// NewOptions 创建一个选项，使用默认的参数配置
func NewOptions() *Options {
	return &Options{
		PutEdgeDirection:  EdgeDirectionLeftTop2Right,
		TakeEdgeDirection: EdgeDirectionLeftTop2Bottom,
		Columns:           2,
		Rows:              3,
		FillCharacter:     'x',
	}
}

// ------------------------------------------------ ---------------------------------------------------------------------
