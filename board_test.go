package main

import "testing"

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
