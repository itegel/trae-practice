package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// 全局游戏状态
var currentGameState GameState

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

// GetGameStateHandler 获取当前游戏状态的API接口
func GetGameStateHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	// 返回当前游戏状态
	json.NewEncoder(w).Encode(currentGameState)
}

// MovePieceHandler 移动棋子的API接口
func MovePieceHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// 处理OPTIONS请求
	if r.Method == "OPTIONS" {
		return
	}

	// 解析请求参数
	var moveRequest struct {
		FromRow   int       `json:"fromRow"`
		FromCol   int       `json:"fromCol"`
		Direction Direction `json:"direction"`
	}

	err := json.NewDecoder(r.Body).Decode(&moveRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 移动棋子
	newState, err := MovePiece(currentGameState, moveRequest.FromRow, moveRequest.FromCol, moveRequest.Direction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 更新全局游戏状态
	currentGameState = newState

	// 返回新的游戏状态
	json.NewEncoder(w).Encode(currentGameState)
}

func main() {
	// 初始化游戏状态
	currentGameState = InitGameState()
	fmt.Println("游戏初始化完成，白方先走")
	PrintGameState(currentGameState)

	// 设置API路由
	http.HandleFunc("/api/game-state", GetGameStateHandler)
	http.HandleFunc("/api/move-piece", MovePieceHandler)

	// 启动HTTP服务器
	fmt.Println("服务器启动，监听端口 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
