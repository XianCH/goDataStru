package main

import (
	"fmt"
)

const (
	Empty = iota
	Player1
	Player2
)

type Game struct {
	board [15][15]int
	turn  int
}

func (g *Game) DrawBoard() {
	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			switch g.board[i][j] {
			case Empty:
				fmt.Print("- ")
			case Player1:
				fmt.Print("X ")
			case Player2:
				fmt.Print("O ")
			}
		}
		fmt.Println()
	}
}

func (g *Game) MakeMove(row, col int) bool {
	if row < 0 || row >= 15 || col < 0 || col >= 15 || g.board[row][col] != Empty {
		return false
	}

	g.board[row][col] = g.turn
	g.turn = 3 - g.turn
	return true
}

func (g *Game) CheckWin(row, col int) bool {
	player := g.board[row][col]

	// Check horizontally
	count := 1
	for i := col + 1; i < 15 && g.board[row][i] == player; i++ {
		count++
	}
	for i := col - 1; i >= 0 && g.board[row][i] == player; i-- {
		count++
	}
	if count >= 5 {
		return true
	}

	// Check vertically
	count = 1
	for i := row + 1; i < 15 && g.board[i][col] == player; i++ {
		count++
	}
	for i := row - 1; i >= 0 && g.board[i][col] == player; i-- {
		count++
	}
	if count >= 5 {
		return true
	}

	// Check diagonally (top-left to bottom-right)
	count = 1
	for i, j := row-1, col-1; i >= 0 && j >= 0 && g.board[i][j] == player; i, j = i-1, j-1 {
		count++
	}
	for i, j := row+1, col+1; i < 15 && j < 15 && g.board[i][j] == player; i, j = i+1, j+1 {
		count++
	}
	if count >= 5 {
		return true
	}

	// Check diagonally (top-right to bottom-left)
	count = 1
	for i, j := row-1, col+1; i >= 0 && j < 15 && g.board[i][j] == player; i, j = i-1, j+1 {
		count++
	}
	for i, j := row+1, col-1; i < 15 && j >= 0 && g.board[i][j] == player; i, j = i+1, j-1 {
		count++
	}
	if count >= 5 {
		return true
	}

	return false
}

func (g *Game) IsFull() bool {
	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			if g.board[i][j] == Empty {
				return false
			}
		}
	}
	return true
}

func main() {
	game := Game{}

	for {
		game.DrawBoard()
		var row, col int
		fmt.Printf("Player %d's turn. Enter row (0-14) and column (0-14): ", game.turn)
		fmt.Scan(&row, &col)

		if game.MakeMove(row, col) {
			if game.CheckWin(row, col) {
				game.DrawBoard()
				fmt.Printf("Player %d wins!\n", game.turn)
				break
			} else if game.IsFull() {
				game.DrawBoard()
				fmt.Println("It's a draw!")
				break
			}
		}
	}
}
