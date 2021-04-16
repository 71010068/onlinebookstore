import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login.vue'
import Admin from '../views/Admin.vue'

// 页面子路由组件
import Index from '../components/admin/Index.vue'
import AddBook from '../components/book/AddBook.vue'
import BookList from '../components/book/BookList.vue'
import CateList from '../components/category/CateList.vue'
import Cart from '../components/cart/Cart.vue'
import UserList from '../components/user/UserList.vue'
// import Personal from '../components/user/Personal.vue'


Vue.use(VueRouter)

const routes = [
  {
    // path: '/login',
    path: '/',

    name: 'login',
    component: Login,
  }, 
  {
    // path: '/admin',  // 配置了vue.config.js所以不用了
    path: '/',
    name: 'Admin',
    component: Admin,
    children:[
      {path:'index',component:Index},
      {path:'addbook',component:AddBook},
      {path:'addbook/:id',component:AddBook,props:true},
      {path:'booklist',component:BookList},
      {path:'catelist',component:CateList},
      {path:'cart',component:Cart},
      {path:'userlist',component:UserList},
      // {path:'userlist/:id',component:Personal},
    ]
  }
]

const router = new VueRouter({
  routes
})

// 在导出之前给一个路由导航守卫
router.beforeEach((to, from, next) => {
  const token = window.sessionStorage.getItem('token')
  if (to.path === '/login') return next()
  if (!token && to.path === '/admin') {
    next('/login')
  } else {
    next()
  }
})

export default router
