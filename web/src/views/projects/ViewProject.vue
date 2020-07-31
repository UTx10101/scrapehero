<template>
  <CRow>
    <CCol col="12" lg="6">
      <CCard>
        <CCardHeader>
          User id:  {{ $route.params.id }}
        </CCardHeader>
      </CCard>
    </CCol>
  </CRow>
</template>

<script>
export default {
  name: 'ViewProject',
  beforeRouteEnter(to, from, next) {
    next(vm => {
      vm.usersOpened = from.fullPath.includes('projects')
    })
  }
}
</script>
