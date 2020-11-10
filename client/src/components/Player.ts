class Brain {
  stack: Array<Function>

  constructor() {
    this.stack = new Array<Function>()

    setInterval(() => {
      this.update()
    }, 1000)
  }

  update(): void {
    console.log('before update:', this.stack)
    const currentStateFunc = this.getCurrentState()

    if (currentStateFunc != null) {
      currentStateFunc()
    }
    console.log('after update', this.stack)
  }

  popState(): Function {
    return this.stack.pop
  }

  pushState(state: Function): void {
    if (this.getCurrentState() != state) {
      this.stack.push(state)
    }
  }

  getCurrentState(): Function {
    return this.stack.length > 0 ? this.stack[this.stack.length - 1] : null
  }
}

const commandMoveRight = {
  type: 'command',
  body: {
    move: 'right',
  },
}

const commandStop = {
  type: 'command',
  body: {
    move: 'stop',
  },
}

const messageSample = {
  type: 'message',
  body: 'FIRE!',
}

export function Player(ws: WebSocket) {
  const brain = new Brain()

  function stopState() {
    ws.send(JSON.stringify(commandStop))
  }

  function moveState() {
    ws.send(JSON.stringify(commandMoveRight))
  }

  function stop() {
    console.log(' push stop...')
    brain.pushState(stopState)
  }

  function move() {
    console.log('push move...')
    brain.pushState(moveState)
  }

  return { brain, move, stop }
}
