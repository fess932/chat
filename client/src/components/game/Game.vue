<template>
  <div ref="dick" class="mem">dick</div>

  <button @click="move">move</button>
  <button @click="stop">stop</button>
</template>

<script lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { Player } from './Player'
import ws from '../socket/socket'

export default {
  setup() {
    const { player, move, stop, dick } = game()

    return { player, move, stop, dick }
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

function game() {
  const dick = ref(null)

  onMounted(() => {
    ws.onclose = (ev: CloseEvent) => {
      console.log(ev, 'onclose')
    }
    ws.onerror = (ev: Event) => {
      console.log(ev, 'onerror')
    }
    ws.onmessage = (ev: MessageEvent) => {
      console.log(ev.data)
      const message = JSON.parse(ev.data)
      console.log(message)
      if (message.type === 'command') {
        console.log('is command!')
        dick.value.style.top = '250px'
        if (message.body.move === 'stop') {
          dick.value.style.left = '30px'
        }
        if (message.body.move === 'right') {
          dick.value.style.left = '400px'
        }
      }
    }
    ws.onopen = (ev: Event) => {
      console.log(ev, 'onopen')
    }
  })

  const player = reactive(Player(ws))

  function move() {
    player.move()
  }

  function stop() {
    player.stop()
  }

  return { player, move, stop, dick }
}
</script>

<style>
.mem {
  font-size: 50px;
  position: absolute;
  left: 100px;
}
</style>
