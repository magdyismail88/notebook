<template>
	<div class="modal fade" :id="uniqueID" tabindex="-1" aria-labelledby="noteEditModalLabel" aria-hidden="true">
	  <div class="modal-dialog modal-fullscreen">
	    <div class="modal-content">
	      <div class="modal-header">
	        <h3 class="modal-title" id="noteShowModalLabel">
	        	{{ n.title }}
	    	</h3>
	        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
	      </div>
	      <div class="modal-body">
	      	<!-- <p class="lead text-center" style="width: 800px !important; margin: 0 auto">
	      		<! {{ n.content }} -->


	      	<!-- </p> -->

	      	<!-- <span v-if="n.is_textual == false" v-html="n.content"></span> -->

					<froalaView v-if="n.is_textual == false" v-model="n.content"></froalaView>

	      	<textarea
	      		rows="20"
	      		class="form-control lead"
	      	 	v-if="n.is_textual == true"
	      	 	readonly
	      	 	disabled>{{ theContent }}</textarea>
	      </div>
	      <div class="modal-footer">
	        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
	      </div>
	    </div>
	  </div>
	</div>
</template>

<script>
	export default {
		name: "NoteShow",
		props: ['n'],
		computed: {
			uniqueID() {
				return `noteShowModal${this.n.id}`;
			},

			theContent() {
				return this.n.content;
			}
		}
	};
</script>

<style scoped>
	textarea {
  		resize: none;
  		width: 800px;
  		margin: 0 auto;
  	}

</style>
