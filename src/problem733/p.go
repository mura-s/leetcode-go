package problem733

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	curColor := image[sr][sc]
	if curColor == color {
		return image
	}

	image[sr][sc] = color
	if sr-1 >= 0 && image[sr-1][sc] == curColor {
		floodFill(image, sr-1, sc, color)
	}
	if sc-1 >= 0 && image[sr][sc-1] == curColor {
		floodFill(image, sr, sc-1, color)
	}
	if sr+1 < len(image) && image[sr+1][sc] == curColor {
		floodFill(image, sr+1, sc, color)
	}
	if sc+1 < len(image[0]) && image[sr][sc+1] == curColor {
		floodFill(image, sr, sc+1, color)
	}

	return image
}
