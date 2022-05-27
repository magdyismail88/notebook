import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import Tab from '@/components/Tab/Tab'
import NoteForm from '@/components/Note/NoteForm'
import NoteEditForm from '@/components/Note/NoteEditForm'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
    	path: '/tab/:id',
    	name: 'Tab',
      component: Tab
    },
    {
      path: '/note/new',
      name: 'NoteForm',
      component: NoteForm
    },
    {
      path: '/note/:id/edit',
      name: 'NoteEditForm',
      component: NoteEditForm
    }
  ],
  mode: "history",
  base: process.env.BASE_URL,
})
