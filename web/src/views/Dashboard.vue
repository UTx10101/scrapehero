<template>
  <div>
    <CRow>
      <CCol col="7">
        <transition name="fade">
          <CCard>
            <CWidgetProgress class="bg-primary">
			  <div class="card-header-actions">
				<CLink class="card-header-action btn-minimize" @click="isCollapsed_pr = !isCollapsed_pr">
				  <CIcon :name="`cil-chevron-${isCollapsed_pr ? 'bottom' : 'top'}`" height="40" style="color: white"/>
			    </CLink>
			  </div>
			  <div class="h1" style="color: white"><CIcon name="cil-layers" height="40"/> {{ totalPrjs }} Projects</div>
			</CWidgetProgress>
            <CCollapse :show="!isCollapsed_pr" :duration="400">
              <CCardBody>
				  <CCardGroup>
					  
						<CWidgetDropdown v-for="prj in prjItems" v-bind:key="prj.pid" :color="clrSchemes[getRand(5)]" :header="prj.pid" :text="prj.pname" style="margin: 1em 1em 1em 1em">
							<template #default>
							  <CDropdown
								color="transparent p-0"
								placement="bottom-end"
							  >
								<template #toggler-content>
								  <CIcon name="cil-settings"/>
								</template>
								<CDropdownItem><CLink :to="`/projects/${prj.pid}/view`">View</CLink></CDropdownItem>
								<CDropdownItem><CLink :to="`/projects/${prj.pid}/edit`">Edit</CLink></CDropdownItem>
								<CDropdownItem><CLink :to="`/projects/${prj.pid}/delete`">Delete</CLink></CDropdownItem>
							  </CDropdown>
							</template>
							<template #footer>
							  <span style="margin: 1em 1.5em 1em 1.5em">Last Modified: {{ prj.last_modified }}</span>
							</template>
						</CWidgetDropdown>
					  
				  </CCardGroup>
              </CCardBody>
            </CCollapse>
          </CCard>
        </transition>
	  </CCol>
	  <CCol>
        <transition name="fade">
          <CCard>
            <CWidgetProgress class="bg-primary">
			  <div class="card-header-actions">
				<CLink class="card-header-action btn-minimize" @click="isCollapsed_api = !isCollapsed_api">
				  <CIcon :name="`cil-chevron-${isCollapsed_pr ? 'bottom' : 'top'}`" height="40" style="color: white"/>
			    </CLink>
			  </div>
			  <div class="h1" style="color: white"><CIcon name="cil-puzzle" height="40"/> {{ totalApiks }} API Keys</div>
			</CWidgetProgress>
            <CCollapse :show="!isCollapsed_api" :duration="400">
              <CCardBody>
                <CDataTable
				  class="mb-0 table-outline"
				  hover
				  :items="apiItems"
				  :fields="apiFields"
				  head-color="light"
				  no-sorting
				>
				  <td slot="api_key" slot-scope="{item}">
					<div>{{item.api_key}}</div>
				  </td>
				  <td
					slot="status"
					slot-scope="{item}"
					class="text-center"
				  >
					<CLink @click="upStatusAPIKey(item)">
					  <CIcon v-if="item.status=='active'" name="cil-chevron-top" style="color: green" v-c-tooltip="'Active, Click to Deactivate'" height="25"/>
					  <CIcon v-else name="cil-chevron-bottom" style="color: orange" v-c-tooltip="'Inactive, Click to Activate'" height="25"/>
					</CLink>
				  </td>
				  <td
					slot="actions"
					slot-scope="{item}"
					class="text-center"
				  >
					<CRow>
						<CCol col="3">
							<CLink v-clipboard="item.api_key" v-c-tooltip="'Copy API Key to Clipboard'" v-clipboard:success="copySuccessF" v-clipboard:error="copyFailedF">
							  <CIcon name="cil-clipboard" style="color: dodgerblue" height="25"/>
							</CLink>
						</CCol>
						<CCol col="3">
							<CLink @click="deleteAPIKey(item)">
							  <CIcon name="cil-trash" style="color: red" v-c-tooltip="'Delete API Key'" height="25"/>
							</CLink>
						</CCol>
					</CRow>
				  </td>
				</CDataTable>
              </CCardBody>
            </CCollapse>
          </CCard>
        </transition>
	  </CCol>
    </CRow>
	<CModal
      title="Copied Successfully"
      color="success"
      :show.sync="copySuccess"
    >
      Your API Key has been copied to clipboard please keep it safe because it is used to access the data from projects externally.
    </CModal>
    <CModal
      title="Copy Failed"
      color="danger"
      :show.sync="copyFailed"
    >
      Failed to copy the API Key due to some error though this error is not so critical but you may report to developer if your are seeing this.
    </CModal>
  </div>
</template>

<script>
export default {
  name: 'Dashboard',
  data () {
    return {
	  clrSchemes: ['primary', 'info', 'success', 'warning', 'danger'],
	  isCollapsed_pr: false,
	  isCollapsed_api: false,
	  copySuccess: false,
	  copyFailed: false,
	  totalPrjs: 4,
	  totalApiks: 4,
	  prjItems: [
        {
		  pid: '1',
          pname: 'SCPh',
          last_modified: 'Today'
        },
		{
		  pid: '2',
          pname: 'SCPh',
          last_modified: 'Today'
        },
		{
		  pid: '3',
          pname: 'SCPh',
          last_modified: 'Today'
        },
		{
		  pid: '4',
          pname: 'SCPh',
          last_modified: 'Today'
        }
      ],
      apiItems: [
        {
		  id: 1,
          api_key: 'afaigjklnaiuefnajkehgkwghwungowipgjw84tujp',
          status: 'active'
        },
		{
		  id: 2,
          api_key: 'afaigjklnaiuefnajkehgkwghwungowipgjw84tujp',
          status: 'inactive'
        },
		{
		  id: 3,
          api_key: 'afaigjklnaiuefnajkehgkwghwungowipgjw84tujp',
          status: 'inactive'
        },
		{
		  id: 4,
          api_key: 'afaigjklnaiuefnajkehgkwghwungowipgjw84tujp',
          status: 'active'
        }
      ],
      apiFields: [
        { key: 'api_key', label: 'API Key', _classes: 'text-center' },
        { key: 'status', _classes: 'text-center' },
        { key: 'actions', _classes: 'text-center' }
      ]
    }
  },
  methods: {
    upStatusAPIKey: function(id){
		return
	},
	copySuccessF: function(){
		this.copySuccess = true;
	},
	copyFailedF: function(){
		this.copyFailed = true;
	},
	getRand: function(x){
		return Math.floor(Math.random() * x)
    }
  }
}
</script>
