<template>

	<div class="container">

		<Flash :status=status :isError=isError :msg=msg />
		
		<div class="card shadow-sm">
			<div class="card-header">
				<div class="modal-header">
			        <!-- <h5 class="modal-title" id="ddd">New Note</h5> -->
			        <button type="button"
						class="btn-close" 
						aria-label="Close"
						@click="backToParentTab"></button>
	     		</div>
			</div>

			<div class="card-body">
				<form class="mt-2">

					<div class="form-group col-md-6">
						<input type="text" class="form-control" v-model="title" placeholder="Type title">
					</div>

					<div class="form-group mt-4">

						<froala :tag="'textarea'"
							:config="config"
							v-model="content">		
						</froala>
					</div>

					<div class="form-group mt-4">
						<button type="button" 
								class="btn btn-primary" 
								@click="newNote">Save</button>
					</div>

				</form>
			</div>
		</div>

	</div>

</template>

<script>
	import Flash from '../Flash';
	import { flashError, flashSuccess } from '../../helpers/utils';
	export default {
		name: "NoteForm",
		mounted() {
			console.log('The current tab is => ', this.$store.getters.getCurrentTab.id);
		},
		data() {
			return {
				config: {
			        events: {
			          'froalaEditor.initialized': function () {
			            console.log('initialized')
			          }
			        },
					height: '340px',
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
      			title: '',
				msg: '',
				status: false,
				isError: false,
      			content: ''
			}
		},
		components: {
			Flash
		},
		methods: {
			newNote() {

				const titleField = (this.title).replace(/\s/g, '');
				const contentField = (this.content).replace(/\s/g, '');

				if(titleField == "" || titleField == null) {
					flashError("Please fill the title field", this);
					return false;
				}

				if(contentField == "" || contentField == null) {
					flashError("Please fill the content field", this);
					return false;
				}

				const data = {
					title: this.title,
					content: this.content
				};

				this.$store.dispatch('createNote', data)
					.then(() => {
						this.$store.dispatch('loadNotes', this.$store.getters.getCurrentTab.id);
						this.title = '';
						this.content = '';
					})
					.then(() => {
						this.$router.push(`/tab/${this.$store.getters.getCurrentTab.id}`);
						flashSuccess("Note Created", this);
					})
					.catch((err) => {
						console.log(err);
					});
			},
			backToParentTab() {
				this.$router.push(`/tab/${this.$store.getters.getCurrentTab.id}`);
			}
		},
		computed: {
			currentTabID() {
				return this.$store.state.tab.tab.id;
			},
			uniqueID() {
				return `noteModal${this.currentTabID}`;
			},
			contentID() {
				return `content${this.currentTabID}`;
			},
			modalLabelUniqueID() {
				return `noteModalLabel${Math.floor(Math.random() * 10)}`
			}
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
