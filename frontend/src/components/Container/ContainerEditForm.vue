<template>
	<div class="modal fade" id="containerEditModal" tabindex="-1" aria-labelledby="containerEditModalLabel"
	 aria-hidden="true">
	  <div class="modal-dialog">
	    <div class="modal-content">
	      <div class="modal-header">
	        <h5 class="modal-title" id="containerEditModalLabel">Edit Container</h5>
	        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
	      </div>
	      <div class="modal-body">
				<Flash :status=status :isError=isError :msg=msg />
	        <form>
	          <div class="mb-3">
	            <label for="titleContainerEditModalField" class="col-form-label">Title</label>
	            <input type="text" class="form-control" id="titleContainerEditModalField" 
	            	v-model="selectedContainer" @focus="getTheLast">
	          </div>
	        </form>
	      </div>
	      <div class="modal-footer">
	        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal"><i class="bi bi-x"></i></button>
	        <button type="button" 
	                class="btn btn-primary"
	                @click="updatingContainer" id="btnContainerSave">
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
		// emits: ['prepare'],
		components: {
			Flash
		},
		// mounted() {
		// 	this.container = this.$store.getters.getContainer;
		// },
		data() {
			return {
				container: {},
				msg: '',
				status: false,
				isError: false
			}
		},
		methods: {
			updatingContainer() {
				const titleContainerField = document.getElementById('titleContainerEditModalField');
				const titleField = (titleContainerEditModalField.value).replace(/\s/g, '');
				
				if(titleField == null || titleField == "") {
					flashError("Fill the title field", this);
					return false;
				}

				this.container.title = titleContainerField.value;

				this.$store.dispatch("updateContainer", this.container)
				.then(() => {
					flashSuccess("Success Updated!", this);
					this.$store.dispatch('loadContainers');
					this.$store.dispatch('changeContainer', this.container.id);
				})
				.catch((err) => {
					console.log(err);
					return false;
				})

				return;

			},
			getTheLast() {
				this.container = this.$store.getters.getContainer;
			}
		},
		computed: {
			selectedContainer() {
        		return this.$store.state.container.container.title;
      		}
		}
	};
</script>