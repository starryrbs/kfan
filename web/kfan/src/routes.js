import Home from './views/Home.vue'
import About from './views/About.vue'
import House from './views/house/index.vue'
import NotFound from './views/NotFound.vue'

/** @type {import('vue-router').RouterOptions['routes']} */
export const routes = [
  { path: '/', component: Home, meta: { title: 'Home' } },
  {
    path: '/about',
    meta: { title: 'About' },
    component: About,
    // example of route level code-splitting
    // this generates a separate chunk (About.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    // component: () => import('./views/About.vue')
  },
  {
    path: '/house',
    meta: { title: '房源' },
    component: House,
  },
  { path: '/:path(.*)', component: NotFound },
]
