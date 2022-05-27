<template>
	<div class="modal fade" id="containerChangeModal" tabindex="-1" 
	:aria-labelledby="modalLabelUniqueID" aria-hidden="true">
	  <div class="modal-dialog">
	    <div class="modal-content">
	      <div class="modal-header">
	        <h5 class="modal-title" :id="modalLabelUniqueID">change container</h5>
	        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
	      </div>
	      <div class="modal-body">

	        <form>
	          <div class="mb-3">
	            <select class="form-control" 
	            		v-model="container_id" 
	            		@change="changingContainer">

	            	<option v-for="container in getContainers" 
	            		    v-bind:key="container.id"
	            		    :value="container.id">{{ container.title }}
	            	</option>

	            </select>
	          </div>
	        </form>
	      </div>
	   
	    </div>
	  </div>
	</div>
</template>

<script>
	import { mapGetters } from 'vuex';
	export default {
		name: "ContainerChangeForm",
		data() {
			return {
				container_id: null
			}
		},
		computed: {
			...mapGetters(['getContainers']),
			modalLabelUniqueID() {
				return `containerChangeModalLabel${Math.floor(Math.random() * 10)}`;
			}
		},
		methods: {
			changingContainer() {
				this.$store.dispatch("changeContainer", this.container_id)
				.then(() => {
					
					if(this.$route.path != '/') {
						this.$router.push('/');
					}

					this.$store.dispatch('loadTabs');
					
				})
				.catch((err) => {
					console.log(err);
				})
			}
		}
	};
</script>