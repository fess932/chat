<template>
  <!--  <img alt="Vue logo" src="./assets/logo.png" />-->
  <div>{{ ok }}</div>
  <hr />
  <HelloWorld />
</template>

<script lang="ts">
import HelloWorld from './components/HelloWorld.vue'
import { ref } from 'vue'

export default {
  name: 'App',
  components: {
    HelloWorld,
  },
  setup() {
    const ok = ref('ok')

    return { ok }
  },
}
</script>
