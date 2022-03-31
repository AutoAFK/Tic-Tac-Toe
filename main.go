package main

import (
	"fmt"
	"os"
	"os/exec"
)

func clear() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	clear()
	fmt.Println("Tic Tac Toe")
	fmt.Println("===========")
	gameLoop := true
	board := [3][3]string{
		{"0", "1", "2"},
		{"3", "4", "5"},
		{"6", "7", "8"}}
	player := "X"
	turn := 0
	for gameLoop {
		printBoard(&board)
		play(&board, &player, &gameLoop)
		turn++
		if turn == 9 {
			fmt.Println("That's a draw!")
			break
		}
	}
}

func printBoard(board *[3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(board[i][j] + " ")
		}
		fmt.Println()
		fmt.Println("-----")
	}
}

func play(board *[3][3]string, player *string, gameLoop *bool) {
	var input string
	for {
		// Take the player input
		fmt.Print("Enter position: ")
		fmt.Scanln(&input)
		// Compare the player input to the board cells
		if input != "X" && input != "O" {
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if input == board[i][j] {
						board[i][j] = *player
						checkForWinner(board, player, gameLoop)
						if *player == "X" {
							*player = "O"
						} else {
							*player = "X"
						}
						return
					}
				}
			}
		}
		fmt.Println("Please choose other place.")
	}
}

func checkForWinner(board *[3][3]string, player *string, gameLoop *bool) {
	matchRow, matchCol := 0, 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if *player == board[i][j] {
				matchRow += 1
			}
			if *player == board[j][i] {
				matchCol += 1
			}
		}
		if matchRow == 3 || matchCol == 3 {
			fmt.Println("Player " + *player + " is the winner!")
			*gameLoop = false
			return
		} else {
			matchRow = 0
			matchCol = 0
		}
	}
	if *player == board[0][0] && *player == board[1][1] && *player == board[2][2] || *player == board[0][2] && *player == board[1][1] && *player == board[2][0] {
		fmt.Println("Player " + *player + " is the winner!")
		*gameLoop = false
		return
	}
}
