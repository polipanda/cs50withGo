package week1

import (
	"fmt"
	"strings"
)

func MarioStair(height int) string {
	var stair string
	for i := 1; i < height+1; i++ {
		stair += fmt.Sprintf("%s%s\n",
			strings.Repeat(" ", height-i),
			strings.Repeat("#", i))
	}

	return stair
}

func MarioStairsWithGap(height int) string {
	var stair string
	for i := 1; i < height+1; i++ {
		stair += fmt.Sprintf("%s%s %s\n",
			strings.Repeat(" ", height-i),
			strings.Repeat("#", i),
			strings.Repeat("#", i))
	}

	return stair
}
