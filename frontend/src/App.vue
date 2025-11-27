<template>
  <div class="game-container">
    <h1 class="game-title">4x4 棋盘游戏</h1>
    
    <div class="game-info">
      <div class="current-turn">
        当前轮次：<span :class="currentPlayerClass">{{ currentPlayerText }}</span>
      </div>
    </div>
    
    <div class="board">
      <div 
        v-for="(row, rowIndex) in gameState.board" 
        :key="rowIndex"
        class="board-row"
      >
        <div 
          v-for="(piece, colIndex) in row" 
          :key="colIndex"
          class="board-cell"
          :class="getCellClass(piece, rowIndex, colIndex)"
          @click="handleCellClick(rowIndex, colIndex)"
        >
          <div class="piece" v-if="piece !== 0">
            {{ piece === 1 ? '白' : '黑' }}
          </div>
        </div>
      </div>
    </div>
    
    <div class="controls" v-if="selectedPiece">
      <h3>选择移动方向：</h3>
      <div class="direction-buttons">
        <button @click="movePiece('up')" class="direction-btn">↑ 上</button>
        <button @click="movePiece('down')" class="direction-btn">↓ 下</button>
        <button @click="movePiece('left')" class="direction-btn">← 左</button>
        <button @click="movePiece('right')" class="direction-btn">→ 右</button>
      </div>
      <button @click="cancelMove" class="cancel-btn">取消</button>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'App',
  data() {
    return {
      gameState: {
        board: [],
        currentTurn: 1
      },
      selectedPiece: null
    }
  },
  computed: {
    currentPlayerText() {
      return this.gameState.currentTurn === 1 ? '白方' : '黑方'
    },
    currentPlayerClass() {
      return this.gameState.currentTurn === 1 ? 'white' : 'black'
    }
  },
  mounted() {
    this.loadGameState()
  },
  methods: {
    async loadGameState() {
      try {
        const response = await axios.get('/api/game-state')
        this.gameState = response.data
      } catch (error) {
        console.error('加载游戏状态失败：', error)
        alert('加载游戏状态失败，请刷新页面重试')
      }
    },
    getCellClass(piece, row, col) {
      const classes = ['cell']
      if (piece === 1) classes.push('white')
      if (piece === 2) classes.push('black')
      if (this.selectedPiece && this.selectedPiece.row === row && this.selectedPiece.col === col) {
        classes.push('selected')
      }
      return classes.join(' ')
    },
    handleCellClick(row, col) {
      const piece = this.gameState.board[row][col]
      
      // 如果没有选中棋子，尝试选中当前棋子
      if (!this.selectedPiece) {
        // 只有当前轮次的棋子可以被选中
        if (piece === this.gameState.currentTurn) {
          this.selectedPiece = { row, col }
        } else if (piece !== 0) {
          alert('这不是当前轮次的棋子！')
        }
        return
      }
      
      // 如果已经选中棋子，点击其他棋子则切换选中
      if (piece === this.gameState.currentTurn) {
        this.selectedPiece = { row, col }
        return
      }
    },
    async movePiece(direction) {
      if (!this.selectedPiece) return
      
      try {
        const directionMap = {
          'up': 1,
          'down': 2,
          'left': 3,
          'right': 4
        }
        
        const response = await axios.post('/api/move-piece', {
          fromRow: this.selectedPiece.row,
          fromCol: this.selectedPiece.col,
          direction: directionMap[direction]
        })
        
        // 更新游戏状态
        this.gameState = response.data
        this.selectedPiece = null
      } catch (error) {
        console.error('移动棋子失败：', error)
        alert(error.response?.data || '移动失败，请检查移动是否合法')
      }
    },
    cancelMove() {
      this.selectedPiece = null
    }
  }
}
</script>

<style scoped>
.game-container {
  text-align: center;
  background-color: white;
  padding: 30px;
  border-radius: 10px;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.1);
}

.game-title {
  margin-bottom: 20px;
  color: #333;
}

.game-info {
  margin-bottom: 20px;
  font-size: 18px;
}

.current-turn {
  padding: 10px;
  background-color: #f5f5f5;
  border-radius: 5px;
}

.current-turn span {
  font-weight: bold;
  padding: 5px 10px;
  border-radius: 3px;
}

.current-turn span.white {
  background-color: #fff;
  color: #333;
  border: 2px solid #ccc;
}

.current-turn span.black {
  background-color: #333;
  color: #fff;
}

.board {
  display: inline-block;
  border: 3px solid #333;
  border-radius: 5px;
  margin-bottom: 20px;
}

.board-row {
  display: flex;
}

.board-cell {
  width: 80px;
  height: 80px;
  border: 1px solid #ccc;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  background-color: #fafafa;
  transition: all 0.3s ease;
}

.board-cell:hover {
  background-color: #e0e0e0;
}

.board-cell.selected {
  background-color: #ffeb3b;
  border-color: #ff9800;
}

.board-cell.white {
  background-color: #fff;
}

.board-cell.black {
  background-color: #f0f0f0;
}

.piece {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 24px;
  font-weight: bold;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.3);
}

.board-cell.white .piece {
  background-color: #fff;
  color: #333;
  border: 2px solid #ccc;
}

.board-cell.black .piece {
  background-color: #333;
  color: #fff;
}

.controls {
  margin-top: 20px;
  padding: 20px;
  background-color: #f5f5f5;
  border-radius: 5px;
}

.controls h3 {
  margin-bottom: 15px;
  color: #333;
}

.direction-buttons {
  display: flex;
  justify-content: center;
  gap: 10px;
  margin-bottom: 15px;
}

.direction-btn {
  padding: 10px 20px;
  font-size: 16px;
  background-color: #2196f3;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.direction-btn:hover {
  background-color: #1976d2;
}

.cancel-btn {
  padding: 8px 16px;
  font-size: 14px;
  background-color: #f44336;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.cancel-btn:hover {
  background-color: #d32f2f;
}
</style>
