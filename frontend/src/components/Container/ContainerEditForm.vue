<template>
	<div class="modal fade"
		id="containerEditModal"
		tabindex="-1"
		aria-labelledby="containerEditModalLabel"
	 	aria-hidden="true">

	  <div class="modal-dialog">
	    <div class="modal-content">
	      <div class="modal-header">
	        <h5 class="modal-title" id="containerEditModalLabel">Edit Container</h5>
	        <button type="button" 
	        	id="containerEditModalCloseBtn" 
				class="btn-close" 
				data-bs-dismiss="modal" 
				aria-label="Close"></button>
	      </div>
	      <div class="modal-body">
				<Flash :status=status :isError=isError :msg=msg />
	        <form>
	          <div class="mb-3">
	            <label for="titleContainerEditModalField" class="col-form-label">Title</label>
	            <input type="text" class="form-control" 
					id="titleContainerEditModalField" 
	            	v-model="container.title">
	          </div>
	        </form>
	      </div>
	      <div class="modal-footer">
	        <button type="button" 
	                class="btn btn-primary"
	                @click="updateContainer" id="btnContainerSave">
	                <i class="bi bi-check2"></i>&nbsp;Save
	        </button>
	      </div>
	    </div>
	  </div>
	</div>
</template>

<script>
	import Flash from "../Flash";
	import { flashError, flashSuccess } from '../../helpers/utils';
	export default {
		name: "ContainerEditForm",
		emits: ['prepareContainerEdit'],
		components: {
			Flash
		},
		data() {
			return {
				container: {
					id: JSON.parse(localStorage.getItem('container')).id,
					title: ''
				},
				msg: '',
				status: false,
				isError: false
			}
		},
		methods: {
			updateContainer() {
				const titleField = (this.container.title).replace(/\s/g, '');

				if(titleField == null || titleField == "") {
					flashError("Fill the title field", this);
					return false;
				}

				this.$store.dispatch("updateContainer", this.container)
					.then(() => {
						flashSuccess("Success Updated!", this);
						this.$store.dispatch('loadContainers');
						this.$store.dispatch('changeContainer', this.container.id);
					})
					.then(() => {
						document.querySelector('#containerEditModalCloseBtn').click()
					})
					.catch((err) => {
						console.log(err);
						return false;
					})

				return;
			}
		}
	};
</script>