import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import SignUp from '@/components/signup'
import FindMatch from '@/components/findmatch'
import Profile from '@/components/profile'
import ContactUs from '@/components/contactus'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/signup',
      name: 'SignUp',
      component: SignUp
    },
    {
      path: '/findmatch',
      name: 'FindMatch',
      component: FindMatch
    },
    {
      path: '/profile',
      name: 'Profile',
      component: Profile
    },
    {
      path: '/contactus',
      name: 'ContactUs',
      component: ContactUs
    }
  ]
})
