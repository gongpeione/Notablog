<template>
  <section>
    <Articles />
  </section>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import Articles from './Articles.vue'
import { setAuth, getAuth } from '@/comment/'

@Component({
  components: {
  Articles
  }
  })
export default class App extends Vue {}
</script>
