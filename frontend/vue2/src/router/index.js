import { createRouter, createWebHistory } from 'vue-router'
import AuthPage from "@/pages/AuthPage";
import MainPage from "@/pages/MainPage";

const routes = [
  {
    path: '/',
    name: 'Auth',
    component: AuthPage
  },
  {
    path: '/main',
    name: 'Main',
    component: MainPage
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
