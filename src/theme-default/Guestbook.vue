<template>
  <section>
    Guestbook
    <Comment />
  </section>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import Comment from '@/comment/'

@Component({
  components: { Comment }
  })
export default class Guestbook extends Vue {}
</script>
