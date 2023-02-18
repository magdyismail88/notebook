<template>
	<div id="note" class="col-4 mt-4" style="margin-top: 36px !important;">
		<div class="card">
			<div class="card-header d-flex justify-content-between">
				{{ note.title | short }} <a href="#" @click="removeNote" class="ml-auto float-right">&times;</a>
			</div>
			<!-- <div class="card-body">{{ note.content | short }}</div> -->
			<div class="card-body">Navigate the detail</div>
			<div class="card-footer">

				<router-link 
					:to="{name: 'NoteEditForm', params: {id: note.id} }"
					class="mt-4"> 
					<i class="bi bi-eye"></i>	
				</router-link>

			</div>
		</div>
	</div>
</template>

<script>
	import NoteEditForm from "./NoteEditForm";
	export default {
		name: "Note",
		props: ['note'],
		mounted() {
			this.$store.dispatch('loadTabs');
		},
		methods: {
			removeNote() {

				let check = confirm("Are you sure?");

				if(!check) {
					return false;
				}

				this.$store.dispatch('destroyNote', this.note.id)
				.then((res) => this.$store.dispatch('loadNotes', this.note.tab_id));
			}
		},
		components: {
			NoteEditForm,
		},
		computed: {
			noteEditFormUniqueID() {
				return `#noteEditModal${this.note.id}`;
			},
			noteShowUniqueID() {
				return `#noteShowModal${this.note.id}`;
			}
		},
		filters: {
			short(str) {
				return str.slice(0, 30);
			}
		}
	};
</script>

<style scoped>
	a {
		text-decoration: none;
	}
</style>
