package main

import "fmt"

// PrintGameState 打印游戏状态，用于调试和演示
func PrintGameState(state GameState) {
    fmt.Println("当前棋盘状态:")
    for row := 0; row < 4; row++ {
        for col := 0; col < 4; col++ {
            var piece string
            switch state.Board[row][col] {
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
    
    currentPlayer := "白方"
    if state.CurrentTurn == Black {
        currentPlayer = "黑方"
    }
    fmt.Printf("当前玩家: %s\n\n", currentPlayer)
}

func main() {
    // 初始化游戏状态
    state := InitGameState()
    fmt.Println("游戏初始化完成，白方先走")
    PrintGameState(state)
    
    // 白方移动棋子：(0,0) 向下移动
    fmt.Println("白方移动棋子：(0,0) 向下移动")
    newState, err := MovePiece(state, 0, 0, Down)
    if err != nil {
        fmt.Printf("移动失败：%v\n\n", err)
    } else {
        state = newState
        PrintGameState(state)
    }
    
    // 黑方移动棋子：(3,0) 向上移动
    fmt.Println("黑方移动棋子：(3,0) 向上移动")
    newState, err = MovePiece(state, 3, 0, Up)
    if err != nil {
        fmt.Printf("移动失败：%v\n\n", err)
    } else {
        state = newState
        PrintGameState(state)
    }
    
    // 尝试移动对方的棋子：白方尝试移动黑方的棋子(2,0) 向上移动
    fmt.Println("白方尝试移动黑方的棋子：(2,0) 向上移动")
    _, err = MovePiece(state, 2, 0, Up)
    if err != nil {
        fmt.Printf("移动失败：%v\n\n", err)
    } else {
        state = newState
        PrintGameState(state)
    }
}
