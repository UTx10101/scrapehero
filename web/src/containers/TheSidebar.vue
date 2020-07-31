<template>
  <CSidebar 
    fixed 
    :minimize="minimize"
    :show="show"
    @update:show="(value) => $store.commit('set', ['sidebarShow', value])"
  >
    <CSidebarBrand class="d-md-down-none" to="/">
      <CIcon 
        class="c-sidebar-brand-full" 
        name="logo" 
        size="custom-size" 
        :height="35" 
        viewBox="0 0 640 640"
      />
      <CIcon 
        class="c-sidebar-brand-minimized" 
        name="logo" 
        size="custom-size" 
        :height="35" 
        viewBox="0 0 640 640"
      />
    </CSidebarBrand>

    <CRenderFunction flat :content-to-render="nav"/>
    <CSidebarMinimizer
      class="d-md-down-none"
      @click.native="$store.commit('set', ['sidebarMinimize', !minimize])"
    />
  </CSidebar>
</template>

<script>
export default {
  name: 'TheSidebar',
  data(){
	  return {
  nav: [{
	_name: 'CSidebarNav',
	_children: [
	  {
		_name: 'CSidebarNavItem',
		name: 'Dashboard',
		to: '/dashboard',
		icon: { name: 'cil-speedometer', class: 'text-white' }
	  },
	  {
		_name: 'CSidebarNavTitle',
		_children: ['Project']
	  },
	  {
		_name: 'CSidebarNavItem',
		name: 'View',
		to: '/projects/'+this.$route.params.id+'/view',
		icon: { name: 'cil-bar-chart', class: 'text-white' }
	  },
	  {
		_name: 'CSidebarNavItem',
		name: 'Edit',
		to: '/projects/'+this.$route.params.id+'/edit',
		icon: { name: 'cil-pencil', class: 'text-white' }
	  },
	  {
		_name: 'CSidebarNavItem',
		name: 'Delete Project',
		to: '/projects/'+this.$route.params.id+'/delete',
		icon: { name: 'cil-trash', class: 'text-white' },
		_class: 'bg-danger text-white'
	  }
	]
  }]
	  }
  },
  computed: {
    show () {
      return this.$store.state.sidebarShow 
    },
    minimize () {
      return this.$store.state.sidebarMinimize 
    }
  }
}
</script>
