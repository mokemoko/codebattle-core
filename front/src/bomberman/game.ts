const DRAW_SIZE = 48
const IMAGE_SIZE = 16
const CHARA_SIZE = 24

const mapSprite = new Image()
const charaSprite = new Image()
mapSprite.src = 'bomberman/sprite.png'
charaSprite.src = 'bomberman/chara.png'

interface TurnInfo {
  map: string[][]
  players: number[]
  bombs: number[]
}

function parseLog(log): TurnInfo[] {
  const turnInfo: TurnInfo[] = []
  const lines = log.split('\n')
  while (lines.length > 0) {
    const [turn, turnMax] = lines.shift().split(' ').map(v => parseInt(v))
    if (isNaN(turnMax)) {
      break
    }
    const [height, width] = lines.shift().split(' ').map(v => parseInt(v))
    const map = [...Array(height)].map(() => lines.shift().split(''))
    const players = [...Array(parseInt(lines.shift()))].map(() => lines.shift().split(' ').map(v => parseInt(v)))
    const bombs = [...Array(parseInt(lines.shift()))].map(() => lines.shift().split(' ').map(v => parseInt(v)))
    bombs.forEach(b => map[b[0]][b[1]] = '@')
    turnInfo.push({ map, players, bombs })
  }
  return turnInfo
}

export class Game {
  ctx: CanvasRenderingContext2D
  turnMax: number
  turnInfo: TurnInfo[]
  playerNames: string[]

  constructor(log: string, canvas: HTMLCanvasElement, playerNames: string[]) {
    this.ctx = canvas.getContext('2d')
    this.ctx.font = '16px "Segoe Print"'
    this.ctx.lineWidth = 4
    this.ctx.textAlign = 'center'
    this.ctx.fillStyle = '#fff'
    this.ctx.strokeStyle = '#000'

    this.turnInfo = parseLog(log)
    this.turnMax = this.turnInfo.length - 1

    this.playerNames = playerNames
  }

  draw(turn: number) {
    this.ctx.clearRect(0, 0, 720, 624)
    const { map, players, bombs } = this.turnInfo[turn]
    map.forEach((row, i) => row.map((cell, j) => {
      const [sx, sy] = {
        '#': [0, 0],
        '+': [IMAGE_SIZE, 0],
        '.': [IMAGE_SIZE * 2, 0],
        '*': [IMAGE_SIZE * 3, 0],
        '|': [IMAGE_SIZE * 4, 0],
        '-': [IMAGE_SIZE * 5, 0],
        '@': [0, IMAGE_SIZE],
        'f': [IMAGE_SIZE, IMAGE_SIZE],
        'b': [IMAGE_SIZE * 2, IMAGE_SIZE],
      }[cell]
      this.ctx.drawImage(mapSprite, sx, sy, IMAGE_SIZE, IMAGE_SIZE,
        j * DRAW_SIZE, i * DRAW_SIZE, DRAW_SIZE, DRAW_SIZE)
    }))
    // TODO: 爆発処理は初期化時に
    bombs.forEach(b => {
      if (b[4] === turn + 1 && turn + 1 <= this.turnMax) {
        const cmap = this.turnInfo[turn].map
        const nmap = this.turnInfo[turn + 1].map
        nmap[b[0]][b[1]] = '*'
        ;[[-1, 0], [1, 0], [0, -1], [0, 1]].forEach(dp => {
          for (let f = 1; f <= b[3]; f++) {
            const [y, x] = [b[0] + dp[0] * f, b[1] + dp[1] * f]
            if (cmap[y][x] === '#') {
              break
            }
            nmap[y][x] = dp[0] === 0 ? '-' : '|'
            if (cmap[y][x] === '+' || cmap[y][x] === '@') {
              break
            }
          }
        })
      }
    })
    players.filter(p => !p[7]).forEach(p => {
      this.ctx.drawImage(charaSprite, CHARA_SIZE * 3 * p[0], CHARA_SIZE * (p[3] === 0 ? 1 : p[3] - 1), CHARA_SIZE, CHARA_SIZE,
        p[2] * DRAW_SIZE, p[1] * DRAW_SIZE, DRAW_SIZE, DRAW_SIZE)
      if (this.playerNames) {
        this.drawText(this.playerNames[p[0]], p[2], p[1])
      }
    })
  }

  drawText(text: string, x: number, y: number) {
    this.ctx.strokeText(text, (x + 0.5) * DRAW_SIZE, y * DRAW_SIZE)
    this.ctx.fillText(text, (x + 0.5) * DRAW_SIZE, y * DRAW_SIZE)
  }
}
