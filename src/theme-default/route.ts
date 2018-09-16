export default [
  {
    path: '/article',
    name: 'Article',
    component: () => import(/* webpackChunkName: "Article" */ './Article.vue')
  },
  {
    path: '/guestbook',
    name: 'Guestbook',
    component: () => import(/* webpackChunkName: "Guestbook" */ './Guestbook.vue')
  }
]
