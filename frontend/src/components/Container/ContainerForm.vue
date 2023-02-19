<template>
	<div class="modal fade" id="containerModal" tabindex="-1" aria-labelledby="containerModalLabel" aria-hidden="true">
	  <div class="modal-dialog">
	    <div class="modal-content">
	      <div class="modal-header">
	        <h5 class="modal-title" id="containerModalLabel">new container</h5>
	        <button type="button" id="container-btn-close" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
	      </div>
	      <div class="modal-body">

			<Flash :status=status :isError=isError :msg=msg />

	        <form>
	          <div class="mb-3">
	            <label for="titleContainerModalField" class="col-form-label">title:</label>
	            <input type="text" class="form-control" id="titleContainerModalField" v-model="title">
	          </div>
	        </form>
	      </div>
	      <div class="modal-footer">
	        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal"><i class="bi bi-x"></i></button>
	        <button type="button" 
	                class="btn btn-primary" 
	                @click="creatingContainer">

	                <i class="bi bi-check2"></i>&nbsp;save
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
		name: "ContainerForm",
		components: {
			Flash
		},
		data() {
			return {
				title: '',
				msg: '',
				status: false,
				isError: false
			}
		},
		methods: {
			creatingContainer() {

				// alert('Demo');

				const titleField = (this.title).replace(/\s/g, '');

				if(titleField == null || titleField == "") {
					flashError("Fill the title field", this);
					return false;
				}

				this.$store.dispatch("createContainer", { title: this.title })
				.then(() => {
					this.title = '';
					flashSuccess("Success Created!", this);
					this.$store.dispatch('loadContainers');
				})
				.then(() => {
					document.querySelector('#container-btn-close').click()
				})
				.catch((err) => {
					console.log(err);
					return false;
				})

				return;

			}
		}
	}
</script>