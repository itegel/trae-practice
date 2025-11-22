package main

// Piece 表示棋子类型
type Piece int

const (
    Empty Piece = iota // 空位置
    White              // 白方棋子
    Black              // 黑方棋子
)

// Board 表示棋盘，是一个4x4的矩阵
type Board [4][4]Piece

// InitBoard 初始化棋盘，返回一个新的棋盘实例
// 白方初始位置：(0,0), (0,1), (0,2), (0,3)
// 黑方初始位置：(3,0), (3,1), (3,2), (3,3)
func InitBoard() Board {
    var board Board
    
    // 初始化白方棋子（第一行）
    for col := 0; col < 4; col++ {
        board[0][col] = White
    }
    
    // 初始化黑方棋子（第四行）
    for col := 0; col < 4; col++ {
        board[3][col] = Black
    }
    
    // 中间两行保持为空，不需要额外处理
    
    return board
}
