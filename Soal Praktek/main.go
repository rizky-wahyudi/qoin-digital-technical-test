package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Point   int
	Dice    int
	Numbers []int
	AddDice int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var N, M, step, one, six int

	fmt.Print("Jumlah Pemain: ")
	fmt.Scanln(&N)
	fmt.Print("Jumlah Dadu: ")
	fmt.Scanln(&M)

	players := make([]Player, N)
	step = 1

	for i := range players {
		players[i].Dice = M
	}

	for gameEnd(players) {
		fmt.Println("========================")
		fmt.Printf("Giliran %d Lempar Dadu:\n", step)
		for i := range players {
			players[i].Numbers, one, six = rollDice(players[i])
			fmt.Printf("          Pemain #%v(%v): %v\n", i+1, players[i].Point, players[i].Numbers)
			players[i].Numbers = removeOneAndSix(players[i].Numbers)
			players[i].Point += six
			players[i].Dice -= (one + six)
			players[nextPlayer(players, i)].AddDice += one
		}
		fmt.Printf("Setelah Evaluasi:\n")
		for i := range players {
			players[i].Dice += players[i].AddDice
			fmt.Printf("          Pemain #%v(%v): %v\n", i+1, players[i].Point, append(players[i].Numbers, populateOne(players[i].AddDice)...))
			players[i].AddDice = 0
		}
		step += 1
	}
	remaining, winner := checkResult(players)
	fmt.Println("========================")
	fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu\n", remaining+1)
	fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya", winner+1)

}

// Function to check the player who still have dice and the player with the largest score
func checkResult(players []Player) (remaining int, winner int) {
	for i, p := range players {
		if p.Point > players[winner].Point {
			winner = i
		}
		if p.Dice != 0 {
			remaining = i
		}
	}
	return
}

// Function to populate dice number 1
func populateOne(count int) (res []int) {
	for i := 0; i < count; i++ {
		res = append(res, 1)
	}
	return
}

// Function to remove dice number 1 and 6
func removeOneAndSix(numbers []int) (res []int) {
	for _, v := range numbers {
		if v != 1 && v != 6 {
			res = append(res, v)
		}
	}
	return
}

// Function to return the next player
func nextPlayer(player []Player, i int) int {
	if i+1 >= len(player) {
		return 0
	}
	return i + 1
}

// Function to check if the condition meet to end the game
func gameEnd(players []Player) bool {
	count := 0
	for _, p := range players {
		if p.Dice == 0 {
			count += 1
		}
		if count >= len(players)-1 {
			return false
		}
	}
	return true
}

// Function to roll dice based on the player's dice
func rollDice(player Player) (numbers []int, one int, six int) {
	for i := 0; i < player.Dice; i++ {
		num := rand.Intn(6) + 1
		numbers = append(numbers, num)
		switch num {
		case 1:
			one += 1
		case 6:
			six += 1
		}
	}
	return
}
