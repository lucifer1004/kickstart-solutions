const readline = require('readline')
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
  terminal: false,
})

let t = 0
let n = 0
let p = 0
let s = []
let tIndex = 0

const main = () => {
  rl.on('line', function(line) {
    if (t === 0) {
      t = Number(line)
      return
    }

    if (n === 0) {
      ;[n, p] = line.split(' ').map(char => Number(char))
      return
    }

    if (s.length === 0) {
      const sRaw = line.split(' ')
      for (let i = 0; i < sRaw.length; i++) {
        s.push(Number(sRaw[i]))
      }

      s.sort((a, b) => b - a)

      const accuSum = [0]
      for (let i = 0; i < s.length; i++) {
        accuSum.push(accuSum[accuSum.length - 1] + s[i])
      }

      let minHours = -1
      for (let i = 0; i < n - p + 1; i++) {
        currentSkill = accuSum[i + p] - accuSum[i]

        targetSkill = p * s[i]

        hours = targetSkill - currentSkill

        if (hours < minHours || minHours < 0) {
          minHours = hours
          continue
        }
      }

      tIndex++
      console.log(`Case #${tIndex}: ${minHours}`)
      n = 0
      p = 0
      s = []
    }

    if (tIndex === t) {
      rl.close()
      return
    }
  })
}

main()
