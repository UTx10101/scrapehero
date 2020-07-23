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
      path: '/',
      redirect: '/dashboard',
      name: 'Home',
      component: () => import('@/containers/TheContainer'),
      children: [
        {
          path: 'dashboard',
          name: 'Dashboard',
          component: () => import('@/views/Dashboard')
        },
        {
          path: 'login',
          name: 'Login',
          component: () => import('@/views/Login')
        },
        {
          path: 'api',
          name: 'API',
          component: () => import('@/views/api/API')
        },
        {
          path: 'projects',
          meta: {
            label: 'Projects'
          },
          component: {
            render(c) {
              return c('router-view')
            }
          },
          children: [
            {
              path: '',
              name: 'Projects',
              component: () => import('@/views/projects/Home')
            },
            {
              path: ':id',
              meta: {
                label: 'Project Details'
              },
              component: {
                render(c) {
                  return c('router-view')
                }
              },
              children: [
                {
                  path: '',
                  name: 'Project',
                  component: () => import('@/views/projects/Project')
                },
                {
                  path: 'edit',
                  name: 'EditProject',
                  component: () => import('@/views/projects/EditProject')
                },
              ]
            }
          ]
        },
        {
          path: '404',
          name: 'Err404',
          component: () => import('@/views/misc/Err404')
        },
        {
          path: '500',
          name: 'Err500',
          component: () => import('@/views/misc/Err500')
        }
      ]
    }
  ]
}

