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
    turnInfo.push({ map, players, bombs })
  }
  return turnInfo
}

export class Game {
  ctx: CanvasRenderingContext2D
  turnMax: number
  turnInfo: TurnInfo[]

  constructor(log: string, canvas: HTMLCanvasElement) {
    this.ctx = canvas.getContext('2d')
    this.turnInfo = parseLog(log)
    this.turnMax = this.turnInfo.length - 1
  }

  draw(turn: number) {
    this.ctx.clearRect(0, 0, 720, 624)
    const { map, players, bombs } = this.turnInfo[turn]
    map.forEach((row, i) => row.map((cell, j) => {
      let [sx, sy] = {
        '#': [0, 0],
        '+': [IMAGE_SIZE, 0],
        '.': [IMAGE_SIZE * 2, 0],
        '*': [IMAGE_SIZE * 3, 0],
        '|': [IMAGE_SIZE * 4, 0],
        '-': [IMAGE_SIZE * 5, 0],
        'f': [IMAGE_SIZE, IMAGE_SIZE],
        'b': [IMAGE_SIZE * 2, IMAGE_SIZE],
      }[cell]
      this.ctx.drawImage(mapSprite, sx, sy, IMAGE_SIZE, IMAGE_SIZE,
        j * DRAW_SIZE, i * DRAW_SIZE, DRAW_SIZE, DRAW_SIZE)
    }))
    bombs.forEach(b => {
      this.ctx.drawImage(mapSprite, 0, IMAGE_SIZE, IMAGE_SIZE, IMAGE_SIZE,
        b[1] * DRAW_SIZE, b[0] * DRAW_SIZE, DRAW_SIZE, DRAW_SIZE)
      if (b[4] === turn + 1 && turn + 1 <= this.turnMax) {
        const cmap = this.turnInfo[turn].map
        const nmap = this.turnInfo[turn + 1].map
        nmap[b[0]][b[1]] = '*'
        ;[[-1, 0], [1, 0], [0, -1], [0, 1]].forEach(dp => {
          for (let f = 1; f <= b[3]; f++) {
            if (cmap[b[0] + dp[0] * f][b[1] + dp[1] * f] === '#') {
              break
            }
            // TODO: 他の爆弾考慮
            nmap[b[0] + dp[0] * f][b[1] + dp[1] * f] = dp[0] === 0 ? '-' : '|'
            if (cmap[b[0] + dp[0] * f][b[1] + dp[1] * f] === '+') {
              break
            }
          }
        })
      }
    })
    players.filter(p => !p[7]).forEach(p => {
      this.ctx.drawImage(charaSprite, CHARA_SIZE * 3 * p[0], CHARA_SIZE * (p[3] === 0 ? 1 : p[3] - 1), CHARA_SIZE, CHARA_SIZE,
        p[2] * DRAW_SIZE, p[1] * DRAW_SIZE, DRAW_SIZE, DRAW_SIZE)
    })
  }
}
