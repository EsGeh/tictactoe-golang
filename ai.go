package main

import (
	"math/rand"
)

func aiPlay(
	gameData game,
) (status string) {
	var nextMove *cellInfo
	scanAllLines(
		gameData,
		func(line []cellInfo) bool {
			player1, player2 := 0, 0
			var freeField cellInfo
			for _, cellInfo := range line {
				switch cellInfo.value {
				case Player1:
					player1++
				case Player2:
					player2++
				default:
					freeField = cellInfo
				}
			}
			if player1 == 2 && player2 == 0 {
				status = "Haha, loser!"
				nextMove = &freeField
				return true
			}
			return false
		},
	)
	if nextMove == nil {
		scanAllLines(
			gameData,
			func(line []cellInfo) bool {
				player1, player2 := 0, 0
				var freeField cellInfo
				for _, cellInfo := range line {
					switch cellInfo.value {
					case Player1:
						player1++
					case Player2:
						player2++
					default:
						freeField = cellInfo
					}
				}
				if player2 == 2 && player1 == 0 {
					status = "Haha, you wish!"
					nextMove = &freeField
					return true
				}
				return false
			},
		)
	}
	if nextMove == nil {
		freeDiagonals := make([]cellInfo, 0, 0)
		for _, row := range []int{0, 2} {
			for _, col := range []int{0, 2} {
				if gameData[row][col] == Empty {
					freeDiagonals = append(freeDiagonals, cellInfo{row, col, Empty})
				}
			}
		}
		if len(freeDiagonals) > 0 {
			i := rand.Intn(len(freeDiagonals))
			nextMove = &freeDiagonals[i]
			status = "Hmm. I'm playing diagonal..."
		}
	}
	if nextMove != nil {
		gameData[nextMove.row][nextMove.col] = 1
		return
	}
	status = "hmm, I pass..."
	return
}
