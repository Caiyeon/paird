import Vue from 'vue'
import Router from 'vue-router'
import GetStarted from '@/components/Home'
import SignUp from '@/components/signup'
import FindMatch from '@/components/findmatch'
import Profile from '@/components/profile'
import ContactUs from '@/components/contactus'
import Main from '@/components/main'
import TechStack from '@/components/techstack'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Main',
      component: Main
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
    },
    {
      path: '/techstack',
      name: 'TechStack',
      component: TechStack
    },
    {
      path: '/getstarted',
      name: 'GetStarted',
      component: GetStarted
    }
  ]
})
