package leetcode

import (
	"fmt"
	"math"
)

func robotSim(commands []int, obstacles [][]int) int {
	result, curDir, mObstacles := float64(0), 0, make(map[string]bool) // 结果 当前方向 障碍物哈希
	x, y, stepX, stepY := 0, 0, []int{0, 1, 0, -1}, []int{1, 0, -1, 0} // 当前位置 以及 xy轴上各个方向移动的大小
	for _, v := range obstacles {
		mObstacles[fmt.Sprintf("%d-%d", v[0], v[1])] = true
	}

	for _, v := range commands {
		switch v {
		case -1:
			curDir = (curDir + 1) % 4
		case -2:
			curDir = (curDir + 3) % 4
		default:
			for i := 0; i < v; i++ {
				tempX, tempY := x+stepX[curDir], y+stepY[curDir]
				// 碰到障碍物，就不要移动了
				if mObstacles[fmt.Sprintf("%d-%d", tempX, tempY)] {
					break
				}
				x, y = tempX, tempY
				result = math.Max(float64(x*x+y*y), result)
			}
		}
	}
	return int(result)
}
