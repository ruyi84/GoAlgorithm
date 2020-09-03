package main

import (
	"testing"
)

/*
使用两个数标记最终位置，如果最终位置为0，0则说明这个机器人回到了原点
*/
func judgeCircle(moves string) bool {
	// 首先判处移动次数为奇数的，如果回到原点肯定能够为偶数
	if len(moves)%2 != 0 {
		return false
	}
	//分别记录垂直和水平位置
	v, l := 0, 0
	for k := range moves {
		switch moves[k] {
		case 85:
			v++
		case 68:
			v--
		case 76:
			l++
		case 82:
			l--
		}
	}
	// 回到起点的情况为两者为0

	return v == l && v == 0

}

/*
U 85
D 68
L 76
R 82
*/
func Test_T657(t *testing.T) {
	str := "UDLR"
}
