import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  linkActiveClass: 'active',
  scrollBehavior: () => ({ x: 0, y: 0 }),
  routes: configRoutes()
})

function configRoutes () {
  return [
	{
	  path: '/login',
	  name: 'Login',
	  meta: {
	    access: 'guest'
	  },
	  component: () => import('@/views/Login')
	},
	{
      path: '/',
      redirect: '/dashboard',
      name: 'Home',
	  meta: {
		access: 'admin',
	  },
      component: () => import('@/containers/TheContainer1'),
      children: [
        {
          path: 'dashboard',
          name: 'Dashboard',
          component: () => import('@/views/Dashboard')
        }
      ]
    },
	{
	  path: '/projects',
	  redirect: '/dashboard',
	  name: 'Project',
	  meta: {
		access: 'admin'
	  },
	  component: () => import('@/containers/TheContainer2'),
	  children: [
		{
		  path: ':id/view',
		  name: 'View',
		  component: () => import('@/views/projects/ViewProject'),
		  props: true
		},
		{
		  path: ':id/edit',
		  name: 'Edit',
		  component: () => import('@/views/projects/EditProject'),
		  props: true
		},
		{
		  path: ':id/delete',
		  name: 'Delete',
		  component: () => import('@/views/projects/DeleteProject'),
		  props: true
		}
	  ]
	},
	{
	  path: '*',
	  name: '404',
	  component: () => import('@/views/misc/404')
	}
  ]
}

