<template>
	<div id="note" class="col-3 mt-4" style="margin-top: 36px !important;">
		<div style="background-color:#fdcb6e;" class="shadow-sm">
			<div class="card-header d-flex justify-content-between">
				{{ note.title | short }} <a href="#" @click="removeNote" class="ml-auto float-right">&times;</a>
			</div>
			<!-- <div class="card-body">{{ note.content | short }}</div> -->
			<router-link 
				:to="{name: 'NoteEditForm', params: {id: note.id} }"
				class="mt-4">
				
				<div class="card-body" 
					:style="note.contentText.length > 220 ? 'height: 230px' : 'height:280px'">
					<!-- {{ note.contentText | get400Chars }} -->
					<p v-for="elem in note.contentText.slice(0, 220).split('.')">{{ elem }}</p>
				</div>
				<div class="card-body" v-if="note.contentText.length > 400">
					<router-link :to="{name: 'NoteEditForm', params: {id: note.id} }">Read more...</router-link>
				</div>
			</router-link>
			<!-- <div class="card-footer">
				
				<router-link 
					:to="{name: 'NoteEditForm', params: {id: note.id} }"
					class="mt-4"> 
					<i class="bi bi-eye"></i>  <span>Read</span>	
				</router-link>

			</div> -->
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
				.then((res) => this.$store.dispatch('loadNotes', this.note.tabId));
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
				return str.length > 20 ? str.slice(0, 20) + '...' : str;
			},
			get400Chars(str) {
				return str.slice(0, 400).split()
			}
		}
	};
</script>

<style scoped>
	a {
		text-decoration: none;
	}

	#note {
		/* background-color: yellow; */
		/* outline: auto; */
	}
</style>
