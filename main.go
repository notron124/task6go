package main

import (
	"fmt"
)

func RightSwitch(heights []int) (studOne, studTwo int) {
	var indexes []int
	oddOnEvenCnt, evenOnOdd := 0, 0

	//itterate over heights of students
	for i, height := range heights {
		//if student in wrong place
		if (i+1)%2 != height%2 {
			indexes = append(indexes, i+1)

			//to check if amount of them are equal
			if (i+1)%2 == 0 && height%2 != 0 {
				oddOnEvenCnt++
			} else if (i+1)%2 != 0 && height%2 == 0 {
				evenOnOdd++
			}
		}

	}

	//if not equal it is not possible
	if oddOnEvenCnt != evenOnOdd {
		return -1, -1
	}

	switch len(indexes) {
	//if indexes is empty, everyone is in rigth possition
	case 0:
		//if there more, than two studentsm we can switch two uneven ones and it will be rigth
		if len(heights) > 2 {
			return 1, 3
		} else { //if there only two stydents, we can't move them
			return -1, -1
		}
	case 2: //if two students in wrong places, switch them
		return indexes[0], indexes[1]
	default: //if one or more then two students in wrong position, switch is impossible
		return -1, -1
	}

}

func main() {
	var studentsAmount int

	fmt.Scan(&studentsAmount)

	heigths := make([]int, studentsAmount)

	for i := range heigths {
		fmt.Scan(&heigths[i])
	}

	fmt.Println(RightSwitch(heigths))
}
