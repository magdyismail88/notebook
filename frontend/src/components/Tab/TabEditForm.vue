<template>
	<div class="modal fade" :id="uniqueID" tabindex="-1" :aria-labelledby="modalLabelUniqueID" aria-hidden="true">
	  <div class="modal-dialog">
	    <div class="modal-content">
	      <div class="modal-header">
	        <h5 class="modal-title" :id="modalLabelUniqueID">Edit Tab</h5>
	        <button type="button" id="tab-edit-btn-close" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
	      </div>
	      <div class="modal-body">

			<Flash :status=status :isError=isError :msg=msg />

	        <form>
	          <div class="mb-3">
	            <label :for='titleId' class="col-form-label">Title</label>
	            <input type="text" class="form-control" :id="titleId" 
	            	v-model="title">
	          </div>
	        </form>
	      </div>
	      <div class="modal-footer">
	        <!-- <button type="button" class="btn btn-secondary" data-bs-dismiss="modal"><i class="bi bi-x"></i></button> -->
	        <button type="button" class="btn btn-primary" @click="updatingTab">
	        	<i class="bi bi-check2"></i>&nbsp;Save
	        </button>
	      </div>
	    </div>
	  </div>
	</div>
</template>

<script>
	import Flash from '../Flash';
	import { flashError, flashSuccess } from '../../helpers/utils';

	export default {
		name: 'TabEditForm',
		props: ['t'],
		data() {
			return {
				msg: '',
				status: false,
				isError: false,
				title: ''
			}
		},
		components: {
			Flash
		},
		methods: {
			updatingTab() {

				const titleField = (this.title).replace(/\s/g, '');

				if(titleField == "" || titleField == null) {
					flashError("Fill the title field", this);
					return false;
				}
				
				const currentTab = JSON.parse(localStorage.getItem('tab'));
				currentTab.title = this.title;

				this.$store.dispatch('updateTab', currentTab)
					.then(() => {
						this.$store.dispatch('loadTabs');
						this.t.title = this.title
					})
					.then(() => {
						document.querySelector('#tab-edit-btn-close').click();
					})
					.catch((err) => console.log(err));
			}
		},
		computed: {
			uniqueID() {
				return `tabEditModal${this.t.id}`;
			},
			titleId() {
				return 'title-' + this.t.id;
			},
			modalLabelUniqueID() {
				return `tabEditModalLabel${Math.floor(Math.random() * 10)}`
			}
		}
	}

</script>
