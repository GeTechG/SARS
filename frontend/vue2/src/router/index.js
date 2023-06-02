import { createRouter, createWebHistory } from 'vue-router'
import AuthPage from "@/pages/AuthPage";
import MainPage from "@/pages/MainPage";
import AttendancePage from "@/pages/AttendancePage";

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
  {
    path: '/attendance/:classId',
    name: 'Attendance',
    component: AttendancePage,
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
