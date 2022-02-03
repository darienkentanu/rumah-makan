package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type pemain struct {
	dadu []int
}

var pemains []*pemain

var onlyOnce sync.Once
var dice = []int{1, 2, 3, 4, 5, 6}

func rollDice() int {

	onlyOnce.Do(func() {
		rand.Seed(time.Now().UnixNano()) // only run once
	})

	return dice[rand.Intn(len(dice))]
}

func rollSeveralDice(m int) (Dices []int) {
	var i int
	for i < m {
		i++
		Dices = append(Dices, rollDice())
	}
	return Dices
}

func main() {
	jumlahDadu := 4
	// jumlahPemain := 3
	// main := mainDadu(jumlahDadu, jumlahPemain)
	fmt.Println(rollSeveralDice(jumlahDadu))
}

func mainDadu(m, n int) {
	// jumlahPemain := n
	// jumlahDadu := m
	// for i := 0; i < jumlahPemain; i++ {
	// 	pemains = append(pemains, &pemain{rollSeveralDice(jumlahDadu)})
	// }
	// for idx, v := range pemains {
	// 	sort.Ints(v.dadu)
	// 	for _, val := range v.dadu {
	// 		if val == 1 {
	// 			v.dadu = v.dadu[1 : len(v.dadu)-1]
	// 			if idx < len(pemains)-1 {
	// 				temp = 1
	// 			}
	// 			pemains[idx+1]
	// 		}
	// 	}
	// }
}
