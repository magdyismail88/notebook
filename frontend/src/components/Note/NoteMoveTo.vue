<template>
	<div class="modal fade" id="noteMoveToModal" tabindex="-1" 
	:aria-labelledby="modalLabelUniqueID" aria-hidden="true">
	  <div class="modal-dialog">
	    <div class="modal-content">
	      <div class="modal-header">
	        <h5 class="modal-title" :id="modalLabelUniqueID">Move to</h5>
	        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
	      </div>
	      <div class="modal-body">

	        <form>
	          <div class="mb-3">
				<select class="form-control">
					<option>Container > Tab</option>
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
		name: "NoteMoveTo",
		data() {
			return {
				containerId: null,
				tabId
			}
		},
		computed: {
			...mapGetters(['getContainers', 'getTabs']),
			modalLabelUniqueID() {
				return `noteMoveToModalLabel${Math.floor(Math.random() * 10)}`;
			}
		},
		methods: {
			moveTo() {
				return
			}
		}
	};
</script>