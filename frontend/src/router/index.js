import Vue from 'vue'
import VueRouter from 'vue-router'
import Host from '../views/Host.vue'
import Memory from '../views/Memory.vue'
import Docker from '../views/Docker.vue'
import Connection from '../views/Connection.vue'
import CPU from '../views/CPU.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Cpu',
    component: CPU
  },
  {
    path: '/host',
    name: 'Host',
    component:Host
  },
  {
    path: '/memory',
    name: 'Memory',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component:Memory
  },
  {
    path: '/docker',
    name: 'Docker',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component:Docker
  },
  {
    path: '/connection',
    name: 'Connection',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: Connection
  },
 
]

const router = new VueRouter({
  
  base: "/",
  routes
})

export default router
