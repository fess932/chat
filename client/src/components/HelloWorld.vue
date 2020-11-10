<template>
  <div>{{ player }}</div>
  <button @click="move">move</button>
  <button @click="stop">stop</button>
</template>

<script lang="ts">
import { ref, reactive } from 'vue'
import { Player } from './Player'

export default {
  setup() {
    const { player, move, stop } = chat()

    return { player, move, stop }
  },
}

function sample() {
  const msg = ref('lel')

  const count = ref(2)

  function add() {
    count.value++
  }

  return { msg, count, add }
}

const host = 'ws://localhost:8000/v1/ws'

function chat() {
  const ws = new WebSocket(host)

  {
    ws.onclose = (ev: CloseEvent, b, c) => {
      console.log(ev, b, c, 'onclose')
    }
    ws.onerror = (ev: Event, b, c) => {
      console.log(ev, b, c, 'onerror')
    }
    ws.onmessage = (ev: MessageEvent, b, c) => {
      // console.log(ev, b, c, 'onmessage')
      console.log(ev.data)
    }
    ws.onopen = (ev: Event) => {
      console.log(ev, 'onopen')
    }
  }

  const player = reactive(Player(ws))

  function move() {
    player.move()
  }

  function stop() {
    player.stop()
  }

  return { player, move, stop }
}
</script>
