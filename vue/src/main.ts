import { createApp } from 'vue'
import './style.css'
import App from './App.vue'

import { createWebHistory, createRouter } from 'vue-router'


import NewPost from './components/NewPost.vue'
import Post from './components/Post.vue'
import UpdatePost from './components/UpdatePost.vue'
import Catalog from './components/Catalog.vue'


const routes = [
    { path: '/', component: Catalog },
    { path: '/newPost', component: NewPost },
    { path: '/post/:id', component: Post },
    { path: '/updatePost/:id', component: UpdatePost },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})



createApp(App).use(router).mount('#app')
