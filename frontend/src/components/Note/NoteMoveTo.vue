<template>
	<div class="modal fade" id="noteMoveToModal" tabindex="-1" 
	:aria-labelledby="modalLabelUniqueID" aria-hidden="true">
	  <div class="modal-dialog">
	    <div class="modal-content">
	      <div class="modal-header">
	        <h5 class="modal-title" :id="modalLabelUniqueID">Move to</h5>
	        <button type="button" id="note-move-close-btn" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
	      </div>
	      <div class="modal-body">

	        <form>
	          <div class="mb-3">
				<select @change="moveTo" id="note-move-select" class="form-control" v-model="tabId">
					<option></option>
					<option v-for="elem in getContainersAndTabs"
						:value="elem.tabId" 
						>{{ elem.containerTitle }} > {{ elem.tabTitle }}</option>
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
				tabId: ''
			}
		},
		computed: {
			...mapGetters(['getContainersAndTabs']),
			modalLabelUniqueID() {
				return `noteMoveToModalLabel${Math.floor(Math.random() * 10)}`;
			}
		},
		methods: {
			moveTo() {
				const note = JSON.parse(localStorage.getItem('note'));
				this.$store.dispatch('moveTo', {noteId: note.id, tabId: this.tabId})
					.then(() => {
						this.$store.dispatch('loadNotes', note.tabId);
						document.querySelector('#note-move-select').selectedIndex = 0;
						document.querySelector('#note-move-close-btn').click();
					})
					.catch((err) => {
						console.log(err);
					});

			}
		}
	};
</script>