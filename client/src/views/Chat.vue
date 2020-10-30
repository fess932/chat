<template>
  <div>
    {{ messages }}
    <div v-for="(message, key) in messages" :key="key">
      <div class="message_container">
        {{ message.message }}
      </div>
    </div>
    <hr />

    <div class="input">
      <textarea
        v-model="newMessage"
        placeholder="Пишите..."
        @keyup.enter.exact="sendMessage"
      ></textarea>
      <button @click="sendMessage">></button>
    </div>
  </div>
</template>

<script lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
const serverUrl = 'ws://localhost:8080/ws'

export default {
  setup() {
    onMounted(() => {
      connectToWebSocket()
    })

    return {
      ws,
      messages,
      newMessage,
      sendMessage,
    }
  },
}

const ws = ref<WebSocket>()

function connectToWebSocket() {
  ws.value = new WebSocket(serverUrl)
  ws.value.addEventListener('open', (e) => {
    console.log('web soc op!', e)
  })

  ws.value.addEventListener('message', (e) => {
    console.log('message!', e)
    handleNewMessage(e)
  })
}

const messages = ref([''])
function handleNewMessage(e: MessageEvent) {
  let data = e.data
  data = data.split(/\r?\n/)
  for (let i = 0; i < data.length; i++) {
    let msg = JSON.parse(data[i])
    messages.value.push(msg)
  }
}

const newMessage = ref('')
function sendMessage() {
  if (newMessage.value !== '' && ws.value !== undefined) {
    const msg = JSON.stringify({ message: newMessage.value })
    ws.value.send(msg)
    newMessage.value = ''
  }
}
</script>
