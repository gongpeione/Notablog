export default [
  {
    path: '/article',
    name: 'Article',
    component: () => import(/* webpackChunkName: "Article" */ './Article.vue')
  }
]
