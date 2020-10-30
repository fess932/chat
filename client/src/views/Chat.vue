<template>
  <div>
    <button @click="rooms.findRoom">Click</button>
    <hr />
    <textarea @keyup.enter.exact="rooms.sendMessage"></textarea>
    <hr />
    {{ rooms }}
  </div>
</template>

<script lang="ts">
import { Rooms } from './room'
import { ref, reactive, computed, onMounted, Ref } from 'vue'
const serverUrl = 'ws://localhost:8000/ws'

export default {
  setup() {
    const rooms = connectToWebSocket()

    return { rooms }
  },
}

function connectToWebSocket(): Rooms {
  const ws = new WebSocket(`${serverUrl}?name=${user.name}`)
  const rooms = new Rooms(ws)

  ws.addEventListener('open', (e) => {
    console.log('web soc op!', e)
  })

  ws.addEventListener('message', (e) => {
    rooms.handleNewMessage(e)
  })

  return rooms
}

// rooms

const roomInput = null
const user = {
  name: 'sample',
}
</script>
