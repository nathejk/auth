import Vue from 'vue'
import Router from 'vue-router'
import qs from 'qs'

const routes = [
  { path: '/', component: 'Frontpage' },
  { path: '/spejder', component: 'Signup', props: { team: "patrulje" } },
	// Notfound
  { path: '*',  component: 'NotFound' },
].map(route => {
  return {
    ...route,
    component: () => import(`@/views/${route.component}.vue`),
  }
})

Vue.use(Router)

const router = new Router({
  mode: 'history',
  routes,
  parseQuery: (query) => {
    return qs.parse(query);
  },
  stringifyQuery(query) {
    let result = qs.stringify(query, { encode: false });

    return result ? ('?' + result) : '';
  }
})

router.setPermissions = function (userPermissions) {
  this.userPermissions = userPermissions
}

router.beforeEach((to, from, next) => {
  const route = router.options.routes.find(r => r.path === to.path)
  const pathPermissions = route && route.meta && route.meta.permissions ? route.meta.permissions : []
  const userPermissions = router.userPermissions
  if (pathPermissions.length == 0) return next()
  if (!userPermissions) return next()
  if (userPermissions.some(el => pathPermissions.includes(el))) return next()
  next({ path: '/' })
})

// Make sure navbar collapses when changing route.
router.beforeEach((to, from, next) => {
  $('#navbarNav').collapse('hide')
  next()
})

// Make sure navbar collapses when changing route.
router.afterEach((to, from) => {
  //window.Vue.$store.commit('app/refreshCurrentUrl', to.fullPath)
  window.scrollTo({top: 0})
})

export default router
