<template>
  <div class="container">
    <div class="game">
      <Game />
    </div>
    <div class="chat">
      <Chat />
    </div>
  </div>
</template>

<script lang="ts">
import Game from './components/game/Game.vue'
import Chat from './components/Chat.vue'

import { ref } from 'vue'

export default {
  name: 'App',
  components: {
    Game,
    Chat,
  },
  setup() {
    return {}
  },
}
</script>

<style scoped>
.container {
  display: grid;
  grid-template-columns: 1fr 1fr;
}

.game {
  background-color: chocolate;
  width: 500px;
  height: 500px;
  border: 3px solid black;
}

.chat {
  background-color: blanchedalmond;
  border: 3px solid grey;
}
</style>
