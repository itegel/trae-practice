package main

import "fmt"

// Piece 表示棋子类型
type Piece int

const (
    Empty Piece = iota // 空位置
    White              // 白方棋子
    Black              // 黑方棋子
)

// Direction 表示移动方向
type Direction int

const (
    Up Direction = iota    // 向上移动
    Down                   // 向下移动
    Left                   // 向左移动
    Right                  // 向右移动
)

// Board 表示棋盘，是一个4x4的矩阵
type Board [4][4]Piece

// GameState 表示游戏状态
type GameState struct {
    Board       Board  // 当前棋盘状态
    CurrentTurn Piece  // 当前轮次的玩家（White或Black）
}

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

// InitGameState 初始化游戏状态，返回一个新的游戏状态实例
// 默认白方先走
func InitGameState() GameState {
    return GameState{
        Board:       InitBoard(),
        CurrentTurn: White,
    }
}

// MovePiece 移动棋子
// 参数：
//   state: 当前游戏状态
//   fromRow: 棋子当前所在行
//   fromCol: 棋子当前所在列
//   direction: 移动方向
// 返回值：
//   newState: 移动后的游戏状态
//   err: 移动失败时的错误信息
func MovePiece(state GameState, fromRow, fromCol int, direction Direction) (GameState, error) {
    // 检查起始位置是否在棋盘范围内
    if fromRow < 0 || fromRow >= 4 || fromCol < 0 || fromCol >= 4 {
        return state, fmt.Errorf("起始位置 (%d,%d) 超出棋盘范围", fromRow, fromCol)
    }
    
    // 检查起始位置是否有当前玩家的棋子
    if state.Board[fromRow][fromCol] != state.CurrentTurn {
        return state, fmt.Errorf("起始位置 (%d,%d) 没有当前玩家的棋子", fromRow, fromCol)
    }
    
    // 计算目标位置
    toRow, toCol := fromRow, fromCol
    switch direction {
    case Up:
        toRow--
    case Down:
        toRow++
    case Left:
        toCol--
    case Right:
        toCol++
    }
    
    // 检查目标位置是否在棋盘范围内
    if toRow < 0 || toRow >= 4 || toCol < 0 || toCol >= 4 {
        return state, fmt.Errorf("目标位置 (%d,%d) 超出棋盘范围", toRow, toCol)
    }
    
    // 检查目标位置是否为空
    if state.Board[toRow][toCol] != Empty {
        return state, fmt.Errorf("目标位置 (%d,%d) 已经被占用", toRow, toCol)
    }
    
    // 创建新的游戏状态
    newState := state
    
    // 移动棋子
    newState.Board[toRow][toCol] = newState.CurrentTurn
    newState.Board[fromRow][fromCol] = Empty
    
    // 切换玩家
    if newState.CurrentTurn == White {
        newState.CurrentTurn = Black
    } else {
        newState.CurrentTurn = White
    }
    
    return newState, nil
}
