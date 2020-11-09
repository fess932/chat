<template>
  <div>
    <!--    <button @click="dial">lel</button>-->
    <!--    <div>default</div>-->
    <!--    <div>{{ clientRef }}</div>-->
    <!--    <div>{{ dial }}</div>-->
  </div>
</template>

<script lang="ts">
// import { Client } from './neffos'
import { ref, onMounted, Ref } from 'vue'
import neffos from 'neffos.js'

export default {
  setup() {
    // const client =
    //
    // const titleRef = ref(null)

    neffos
      .dial('ws://localhost:8000/ws', {})
      .then((client) => {
        console.log(client)
      })
      .catch((e) => {
        console.log(e)
      })

    // onMounted(async () => {
    //   const clientRef = await neffos.dial('ws://localhost:8000/ws', {})
    //   clientRef.console.log('titleRef: ', titleRef.value)
    // })
    //
    // const { clientRef, dial } = useMessages()
    // return { clientRef, dial }
  },
}

// function useMessages() {
//   let clientRef: Ref<neffos.Conn>
//   const clietnRef = ref(null)
//
//   async function dial() {
//     console.log('lel?', clientRef)
//
//     if (clientRef.value.isClosed()) {
//       console.log('closed')
//     }
//
//     if (clientRef) {
//       try {
//         clientRef.value = await neffos.dial('ws://localhost:8000/ws', {})
//       } catch (e) {
//         console.log(e)
//       }
//     }
//   }
//
//   return { clientRef, dial }
// }
//
// function useSendPing() {
//   const messages = ref('пусто')
//
//   console.log('lel')
//
//   return { messages }
// }
//
// const { messages } = useSendPing()
//
// function lel(e: PointerEvent, u: any) {
//   console.log('1', e, '2: ', u)
// }

// client = await client.dial('ws://localhost:8000/ws')

// const client = new Client()
//
// client.dial('ws://localhost:8000/ws')
// const ping = client.ping
//
// const messages = ref('')
</script>
