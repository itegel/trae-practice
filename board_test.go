package main

import (
    "testing"
)

// TestInitBoard 测试棋盘初始化功能
func TestInitBoard(t *testing.T) {
    board := InitBoard()
    
    // 检查白方棋子是否在正确位置（第一行）
    for col := 0; col < 4; col++ {
        if board[0][col] != White {
            t.Errorf("期望位置 (0,%d) 是白方棋子，但实际是 %d", col, board[0][col])
        }
    }
    
    // 检查黑方棋子是否在正确位置（第四行）
    for col := 0; col < 4; col++ {
        if board[3][col] != Black {
            t.Errorf("期望位置 (3,%d) 是黑方棋子，但实际是 %d", col, board[3][col])
        }
    }
    
    // 检查中间两行是否为空
    for row := 1; row < 3; row++ {
        for col := 0; col < 4; col++ {
            if board[row][col] != Empty {
                t.Errorf("期望位置 (%d,%d) 是空的，但实际是 %d", row, col, board[row][col])
            }
        }
    }
}

// TestBoardSize 测试棋盘大小是否为4x4
func TestBoardSize(t *testing.T) {
    board := InitBoard()
    
    // 检查行数是否为4
    if len(board) != 4 {
        t.Errorf("期望棋盘有4行，但实际有 %d 行", len(board))
    }
    
    // 检查每行是否有4列
    for row := 0; row < 4; row++ {
        if len(board[row]) != 4 {
            t.Errorf("期望第 %d 行有4列，但实际有 %d 列", row, len(board[row]))
        }
    }
}

// TestInitGameState 测试游戏状态初始化
func TestInitGameState(t *testing.T) {
    state := InitGameState()
    
    // 检查棋盘是否正确初始化
    for col := 0; col < 4; col++ {
        if state.Board[0][col] != White {
            t.Errorf("期望位置 (0,%d) 是白方棋子，但实际是 %d", col, state.Board[0][col])
        }
        if state.Board[3][col] != Black {
            t.Errorf("期望位置 (3,%d) 是黑方棋子，但实际是 %d", col, state.Board[3][col])
        }
    }
    
    // 检查当前玩家是否为白方
    if state.CurrentTurn != White {
        t.Errorf("期望当前玩家是白方，但实际是 %d", state.CurrentTurn)
    }
}

// TestMovePiece 测试棋子移动
func TestMovePiece(t *testing.T) {
    // 初始化游戏状态
    state := InitGameState()
    
    // 测试1：白方棋子向下移动
    newState, err := MovePiece(state, 0, 0, Down)
    if err != nil {
        t.Errorf("期望移动成功，但实际失败了：%v", err)
    }
    if newState.Board[1][0] != White {
        t.Errorf("期望位置 (1,0) 是白方棋子，但实际是 %d", newState.Board[1][0])
    }
    if newState.Board[0][0] != Empty {
        t.Errorf("期望位置 (0,0) 是空的，但实际是 %d", newState.Board[0][0])
    }
    if newState.CurrentTurn != Black {
        t.Errorf("期望当前玩家是黑方，但实际是 %d", newState.CurrentTurn)
    }
    
    // 测试2：黑方棋子向上移动
    newState, err = MovePiece(newState, 3, 0, Up)
    if err != nil {
        t.Errorf("期望移动成功，但实际失败了：%v", err)
    }
    if newState.Board[2][0] != Black {
        t.Errorf("期望位置 (2,0) 是黑方棋子，但实际是 %d", newState.Board[2][0])
    }
    if newState.Board[3][0] != Empty {
        t.Errorf("期望位置 (3,0) 是空的，但实际是 %d", newState.Board[3][0])
    }
    if newState.CurrentTurn != White {
        t.Errorf("期望当前玩家是白方，但实际是 %d", newState.CurrentTurn)
    }
    
    // 测试3：尝试移动对方的棋子
    _, err = MovePiece(newState, 2, 0, Up)
    if err == nil {
        t.Errorf("期望移动失败，因为尝试移动对方的棋子，但实际成功了")
    }
    
    // 测试4：尝试移动到超出棋盘范围的位置
    _, err = MovePiece(newState, 0, 0, Up)
    if err == nil {
        t.Errorf("期望移动失败，因为目标位置超出棋盘范围，但实际成功了")
    }
    
    // 测试5：尝试移动到已经被占用的位置
    // 先移动一个棋子到 (1,1) 位置
    newState, err = MovePiece(newState, 0, 1, Down)
    if err != nil {
        t.Errorf("期望移动成功，但实际失败了：%v", err)
    }
    // 再尝试移动另一个棋子到 (1,1) 位置
    _, err = MovePiece(newState, 0, 2, Down)
    if err == nil {
        t.Errorf("期望移动失败，因为目标位置已经被占用，但实际成功了")
    }
}

// TestMovePieceWithComplexBoard 测试在复杂棋盘布局下的棋子移动
func TestMovePieceWithComplexBoard(t *testing.T) {
    // 创建一个复杂的棋盘布局
    var board Board
    board[0][0] = White
    board[0][1] = White
    board[0][2] = Black
    board[0][3] = Black
    board[1][0] = White
    board[1][1] = Empty
    board[1][2] = Empty
    board[1][3] = Black
    board[2][0] = White
    board[2][1] = Empty
    board[2][2] = Empty
    board[2][3] = Black
    board[3][0] = White
    board[3][1] = White
    board[3][2] = Black
    board[3][3] = Black
    
    // 创建游戏状态，当前玩家为白方
    state := GameState{
        Board:       board,
        CurrentTurn: White,
    }
    
    // 测试1：白方棋子向右移动
    newState, err := MovePiece(state, 1, 0, Right)
    if err != nil {
        t.Errorf("期望移动成功，但实际失败了：%v", err)
    }
    if newState.Board[1][1] != White {
        t.Errorf("期望位置 (1,1) 是白方棋子，但实际是 %d", newState.Board[1][1])
    }
    if newState.Board[1][0] != Empty {
        t.Errorf("期望位置 (1,0) 是空的，但实际是 %d", newState.Board[1][0])
    }
    if newState.CurrentTurn != Black {
        t.Errorf("期望当前玩家是黑方，但实际是 %d", newState.CurrentTurn)
    }
    
    // 测试2：黑方棋子向左移动
    newState, err = MovePiece(newState, 1, 3, Left)
    if err != nil {
        t.Errorf("期望移动成功，但实际失败了：%v", err)
    }
    if newState.Board[1][2] != Black {
        t.Errorf("期望位置 (1,2) 是黑方棋子，但实际是 %d", newState.Board[1][2])
    }
    if newState.Board[1][3] != Empty {
        t.Errorf("期望位置 (1,3) 是空的，但实际是 %d", newState.Board[1][3])
    }
    if newState.CurrentTurn != White {
        t.Errorf("期望当前玩家是白方，但实际是 %d", newState.CurrentTurn)
    }
    
    // 测试3：白方棋子向下移动
    newState, err = MovePiece(newState, 0, 0, Down)
    if err != nil {
        t.Errorf("期望移动成功，但实际失败了：%v", err)
    }
    if newState.Board[1][0] != White {
        t.Errorf("期望位置 (1,0) 是白方棋子，但实际是 %d", newState.Board[1][0])
    }
    if newState.Board[0][0] != Empty {
        t.Errorf("期望位置 (0,0) 是空的，但实际是 %d", newState.Board[0][0])
    }
    if newState.CurrentTurn != Black {
        t.Errorf("期望当前玩家是黑方，但实际是 %d", newState.CurrentTurn)
    }
    
    // 测试4：黑方棋子向上移动
    newState, err = MovePiece(newState, 3, 2, Up)
    if err != nil {
        t.Errorf("期望移动成功，但实际失败了：%v", err)
    }
    if newState.Board[2][2] != Black {
        t.Errorf("期望位置 (2,2) 是黑方棋子，但实际是 %d", newState.Board[2][2])
    }
    if newState.Board[3][2] != Empty {
        t.Errorf("期望位置 (3,2) 是空的，但实际是 %d", newState.Board[3][2])
    }
    if newState.CurrentTurn != White {
        t.Errorf("期望当前玩家是白方，但实际是 %d", newState.CurrentTurn)
    }
}
