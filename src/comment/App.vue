<template>
  <section>
    <Editor v-if="editable" />
    <Comments />
  </section>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator'
import Comments from './Comments.vue'
import Editor from './Editor.vue'
import { getAPI, getAuth } from './index'

@Component({
  components: { Comments, Editor }
  })
export default class Comment extends Vue {
  @Prop({ default: location.href }) uuid!: string | number
  @Prop({ default: true }) editable!: boolean

  API = getAPI()
  auth = getAuth()
}
</script>
