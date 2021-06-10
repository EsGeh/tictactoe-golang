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
			player1Counter, player2Counter := 0, 0
			var freeField cellInfo
			for _, cellInfo := range line {
				switch cellInfo.value {
				case player1:
					player1Counter++
				case player2:
					player2Counter++
				default:
					freeField = cellInfo
				}
			}
			if player1Counter == 2 && player2Counter == 0 {
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
				player1Counter, player2Counter := 0, 0
				var freeField cellInfo
				for _, cellInfo := range line {
					switch cellInfo.value {
					case player1:
						player1Counter++
					case player2:
						player2Counter++
					default:
						freeField = cellInfo
					}
				}
				if player2Counter == 2 && player1Counter == 0 {
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
				if gameData[row][col] == empty {
					freeDiagonals = append(freeDiagonals, cellInfo{row, col, empty})
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
