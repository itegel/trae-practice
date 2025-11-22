package main

import "fmt"

// PrintBoard 打印棋盘状态，用于调试和演示
func PrintBoard(board Board) {
    fmt.Println("棋盘状态:")
    for row := 0; row < 4; row++ {
        for col := 0; col < 4; col++ {
            var piece string
            switch board[row][col] {
            case Empty:
                piece = "·"
            case White:
                piece = "+"
            case Black:
                piece = "-"
            }
            fmt.Printf("%s ", piece)
        }
        fmt.Println()
    }
}

func main() {
    // 演示如何调用InitBoard函数
    board := InitBoard()
    PrintBoard(board)
}
