package Mario

import "fmt"

func Pyramid(height int) {
	for {
		_, response := fmt.Scan(&height)
		if response != nil {
			continue
		}

		if height < 1 || height > 8 {
			continue
		}

		for i := 0; i < height; i++ {
			for j := 0; j < height-i-1; j++ {
				fmt.Print(" ")
			}
			for k := 0; k <= i; k++ {
				fmt.Print("#")
			}
			fmt.Println()
		}

		break
	}
}
