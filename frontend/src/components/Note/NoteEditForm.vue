<template>

	<div class="container">

		<Flash :status=status :isError=isError :msg=msg />
		
		<div class="card shadow-sm">
			
				<div class="modal-header">
			        <!-- <h5 class="modal-title"></h5> -->
			        <button type="button" 
			        		class="btn-close" 
			        		aria-label="Close" 
			        	    @click="backToParentTab"></button>
			    </div>
		    

			<div class="card-body">
				<form class="mt-2">

					<div class="form-group col-md-6">
						<input type="text" 
						       class="form-control" 
						       v-model="note.title" placeholder="Type title">
					</div>

					<a href="javascript:void(0)" 
						style="text-decoration:none;margin-top: 10px;" 
						@click="updatingNote"
						><i class="bi bi-save"></i> Save</a>

					<div class="form-group mt-4">
						<froala :tag="'textarea'"
							:config="config"
							v-model="note.content">		
						</froala>
					</div>

					<!-- <div class="form-group mt-4">
						<input
							class="form-check-input"
							type="checkbox"
							:id="isTextualID"
							v-model="note.is_textual" />

						 	<label class="form-check-label" :for="isTextualID">
						    	textual content
						 	</label>
					</div> -->

					<div class="form-group mt-4">
						<button type="button" class="btn btn-primary"
						 @click="updatingNote">Save</button>
					</div>

				</form>
			</div>
		</div>

	</div>

</template>

<script>
	import Flash from '../Flash';
	import { flashError, flashSuccess, center } from '../../helpers/utils';
	export default {
		name: "NoteForm",
		mounted() {
			this.$store.dispatch('loadNote', this.$route.params.id)
			.then(() => {
				this.note = this.$store.getters.getNote;
			})
			.catch((err) => {
				console.log(err);
			})
		},
		data() {
			return {
				config: {
			        events: {
			          'froalaEditor.initialized': function () {
			            console.log('initialized')
			          }
			        },
					height: '800px',
			        imageUpload: true,
			        imageUploadMethod: 'POST',
			        imageUploadParam: 'file',
			        imageUploadURL: 'http://localhost:8888/api/upload',
			        tableInsertHelper: true,
			        fileUpload: true,
			        colorsHEXInput: true,
			        colorsBackground: ['#61BD6D', '#1ABC9C', '#54ACD2', 'REMOVE'],
			        colorsButtons: ["colorsBack", "|", "-"],
			        dragInline: true,
        			charCounterCount: false
      			},
				msg: '',
				status: false,
				isError: false,
				note: {}
			}
		},
		components: {
			Flash
		},
		methods: {
			updatingNote() {
				const titleField = (this.note.title).replace(/\s/g, '');
				const contentField = (this.note.content).replace(/\s/g, '');

				// Validation Section

				if(titleField == "" || titleField == null) {
					flashError("Please fill the title field", this);
					return false;
				}

				if(contentField == "" || contentField == null) {
					flashError("Please fill the content field", this);
					return false;
				}

				this.$store.dispatch('updateNote', this.note)
					.then(() => {
						this.$store.dispatch('loadNotes', this.note.tabId);
					})
					.then(() => {
						flashSuccess("Note Updated", this);
					})
					.catch((err) => {
						console.log(err);
					});
			},
			loadingNote() {
				this.$store.dispatch('loadNote', this.$route.params.id)
					.then(() => {
						this.note = this.$store.getters.getNote;
					})
					.catch((err) => {
						console.log(err);
					})
			},
			backToParentTab() {
				this.$router.push(`/tab/${this.note.tabId}`);
			}
			
		},
		computed: {
			contentID() {
				return `content${this.currentTabID}`;
			},
			isTextualID() {
				return `isTextual${this.currentTabID}`;
			}
		},
		watch: {
			'$route': 'loadingNote'
		}
	};
</script>

<style scoped>
	.div-block {
		display: block !important;
	}

	.modal-dark {
		background: #212529 !important;
	}

</style>
