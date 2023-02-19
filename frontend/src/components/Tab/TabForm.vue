<template>
	<div class="modal fade" id="tabModal" tabindex="-1" 
	:aria-labelledby="modalLabelUniqueID" aria-hidden="true">
	  <div class="modal-dialog">
	    <div class="modal-content">
	      <div class="modal-header">
	        <h5 class="modal-title" :id="modalLabelUniqueID">new tab</h5>
	        <button type="button" id="tab-btn-close" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
	      </div>
	      <div class="modal-body">

			<Flash :status=status :isError=isError :msg=msg />

	        <form>
	          <div class="mb-3">
	            <label :for="titleUniqueID" class="col-form-label">title: </label>
	            <input type="text" class="form-control" :id="titleUniqueID" v-model="title">
	          </div>
	        </form>
	      </div>
	      <div class="modal-footer">
	        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal"><i class="bi bi-x"></i></button>
	        <button type="button" class="btn btn-primary" @click="newTab">
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
		name: 'TabForm',
		data() {
			return {
				title: '',
				msg: '',
				status: false,
				isError: false
			}
		},
		components: {
			Flash
		},
		methods: {
			newTab() {
				const titleField = (this.title).replace(/\s/g, '');

				if(titleField == "" || titleField == null) {
					flashError("Fill the title field", this);
					return false;
				}

				this.$store.dispatch('createTab', { title: this.title })
				.then(() => {
					this.$store.dispatch('loadTabs');
					flashSuccess("Created Successfully", this);
					this.title = '';
				})
				.then(() => {
					document.querySelector('#tab-btn-close').click()
				})
				.catch((err) => console.log(err));
			}
		},
		computed: {
			titleUniqueID() {
				return `titleField${Math.floor(Math.random() * 10)}`;
			},
			modalLabelUniqueID() {
				return `tabModalLabel${Math.floor(Math.random() * 10)}`
			}
		}
	}

</script>
