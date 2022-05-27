<template>
	<div class="modal fade" :id="uniqueID" tabindex="-1" :aria-labelledby="modalLabelUniqueID" aria-hidden="true">
	  <div class="modal-dialog">
	    <div class="modal-content">
	      <div class="modal-header">
	        <h5 class="modal-title" :id="modalLabelUniqueID">edit tab</h5>
	        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
	      </div>
	      <div class="modal-body">

			<Flash :status=status :isError=isError :msg=msg />

	        <form>
	          <div class="mb-3">
	            <label :for="titleUniqueID" class="col-form-label">title: </label>
	            <input type="text" class="form-control" id="titleUniqueID" 
	            v-model="t.title">
	          </div>
	        </form>
	      </div>
	      <div class="modal-footer">
	        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal"><i class="bi bi-x"></i></button>
	        <button type="button" class="btn btn-primary" @click="updatingTab">
	        	<i class="bi bi-check2"></i>&nbsp;save
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
				isError: false
			}
		},
		components: {
			Flash
		},
		methods: {
			updatingTab() {

				let titleField = (this.t.title).replace(/\s/g, '');

				if(titleField == "" || titleField == null) {
					flashError("Fill the title field", this);
					return false;
				}

				this.$store.dispatch('updateTab', this.t)
				.then(() => {
					this.$store.dispatch('loadTabs');
					flashSuccess("Updated Successfully", this);
				})
				.catch((err) => console.log(err));
			}
		},
		computed: {
			uniqueID() {
				return `tabEditModal${this.t.id}`;
			},
			titleUniqueID() {
				return `titleField${this.t.id}`;
			},
			modalLabelUniqueID() {
				return `tabEditModalLabel${Math.floor(Math.random() * 10)}`
			}
		}
	}

</script>
