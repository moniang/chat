import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

const routes = [
  {
    path: '/',
    name: 'Index',
    component: () => import(/* webpackChunkName: "index" */ '@/views/Index.vue'),
    meta: {
      title: '登录'
    }
},  {
    path: '/chat',
    name: 'chat',
    component: () => import(/* webpackChunkName: "index" */ '@/views/Chat/Index.vue'),
    meta: {
      title: '聊天'
    }
  }
]

const router = new Router({
  routes
})

export default router
