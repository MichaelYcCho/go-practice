package grammar

import "fmt"

func DimensionalSlice() {
	jb := []string{"James", "Bond", "Chocolate", "martini"}

	mp := []string{"Miss", "Moneypenny", "Strawberry", "hazelnut"}

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
