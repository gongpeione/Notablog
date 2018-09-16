import Vue from 'vue'
import NotaBlog from './NotaBlog.vue'
import router from './router'
import store from './store/'
import './utils/registerServiceWorker'
import { setAuth } from '@/comment/'
import Config from './config.json'

setAuth(Config.comment.auth)

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(NotaBlog)
}).$mount('#app')
