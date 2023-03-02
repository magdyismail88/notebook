<template>
	<div id="note" class="col-3 mt-4" style="margin-top: 36px !important;">
		<div style="background-color:#fdcb6e;" class="shadow-sm">
			<div class="card-header d-flex justify-content-between">
				{{ note.title | short }}
				<!-- <a href="#" @click="removeNote" class="ml-auto float-right">&times;</a> -->
				<div class="dropdown">
  <button class="btn btn-sm" type="button" data-bs-toggle="dropdown" aria-expanded="false">
    <!-- <i class="bi bi-gear-fill"></i> -->
	<!-- <i class="bi bi-sliders"></i> -->
	<i class="bi bi-three-dots"></i>
  </button>
  <ul class="dropdown-menu">
    <li>
		<a href="#"
            class="nav-link text-primary"
            data-bs-toggle="modal"
            data-bs-target="#noteMoveToModal"
            data-bs-whatever="@getbootstrap" @click="setNote">
			<i class="bi bi-arrows-move"></i> Move
		</a>
	
	</li>
    <li><a href="javascript:void(0)" @click="removeNote" class="dropdown-item text-danger"><i class="bi bi-trash"></i> Remove</a></li>
  </ul>
</div>
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
				const check = confirm("Are you sure?");
				if(!check) {
					return false;
				}
				this.$store.dispatch('destroyNote', this.note.id)
					.then((res) => this.$store.dispatch('loadNotes', this.note.tabId));
			},
			setNote() {
				localStorage.setItem('note', JSON.stringify(this.note))
				this.$store.dispatch('loadContainersAndTabs')
					.then(() => {
						document.querySelector('#note-move-select').selectedIndex = 0;
						console.log('Containers and tabs have loaded!');
					})
					.catch(err => {
						console.log(err);
					})
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
