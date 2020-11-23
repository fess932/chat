<template>
  <div class="message-container">
    <div class="messages">
      <div v-for="(v, k) in messages" :key="k">
        {{ v }}
      </div>
    </div>
    <div class="send-message">
      <input
        v-model="message"
        class="send-message_text"
        type="text"
        placeholder="message..."
        @keyup.prevent.enter="sendMessage"
      />
      <input
        @click="sendMessage"
        class="send-message_button"
        type="button"
        value="send"
      />
    </div>
  </div>
</template>

<script lang="ts">
import ws from './socket/socket'
import { ref, onMounted } from 'vue'
import alarm from '../assets/sound.wav'

import { Actions } from '../models/commands'

export default {
  setup() {
    const { messages, message, sendMessage } = chat()

    return { messages, message, sendMessage }
  },
}

function chat() {
  const messages = ref([''])
  const message = ref('')
  const alarm2 = new Audio(alarm)

  async function sendMessage() {
    if (message.value === '') {
      return
    }

    const messageSample = {
      action: Actions.SendMessage,
      message: message.value,
    }

    const msg = JSON.stringify(messageSample)
    console.log(msg)

    ws.send(JSON.stringify(messageSample))
    message.value = ''
  }

  onMounted(() => {
    ws.addEventListener('open', (ev: Event) => {
      console.log('onopen')
    })

    ws.addEventListener('close', (ev: CloseEvent) => {
      console.log('onclose')
    })

    ws.addEventListener('error', (ev: Event) => {
      console.log(ev, 'onerror')
    })

    ws.addEventListener('message', (ev: MessageEvent) => {
      const message = JSON.parse(ev.data)
      console.log(message)

      if (message.action === Actions.SendMessage) {
        console.log('is message!')
        messages.value.push(message.message)
        alarm2.play()
      }
    })
  })

  return { messages, message, sendMessage }
}
</script>

<style scoped>
.message-container {
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  justify-content: space-between;
}

.send-message {
  height: 50px;
  width: 100%;
  display: flex;
  justify-content: space-between;
}

.send-message_text {
  height: 100%;
  font-size: 1.3em;
}

.send-message_button {
  font-size: 1.3em;
}

.messages {
  margin-left: auto;
  margin-right: auto;
}
</style>
