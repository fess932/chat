import { Ref, ref } from 'vue'

class Room {
  newMessage: string
  name: string
  messages: Array<string>

  constructor() {
    this.newMessage = ''
    this.name = ''
    this.messages = []
  }
}

class Rooms {
  rooms: Ref<Array<Room>>
  ws: WebSocket

  constructor(ws: WebSocket) {
    this.rooms = ref(new Array<Room>())
    this.ws = ws
  }

  findRoom(roomName: string): Room {
    this.rooms.value.push(new Room())

    for (let i = 0; i < this.rooms.value.length; i++) {
      if (this.rooms.value[i].name === roomName) {
        return this.rooms.value[i]
      }
    }
    return new Room()
  }

  sendMessage(room: Room) {
    if (room.newMessage !== '') {
      this.ws.send(
        JSON.stringify({
          action: 'send-message',
          message: room.newMessage,
          target: room.name,
        })
      )
      room.newMessage = ''
    }
  }

  handleNewMessage(e: MessageEvent) {
    let data = e.data
    data = data.split(/\r?\n/)

    for (let i = 0; i < data.length; i++) {
      let msg = JSON.parse(data[i])

      const room = this.findRoom(msg.target)
      if (typeof room !== 'undefined') {
        room.messages.push(msg)
      }
    }
  }
}

export { Rooms }
