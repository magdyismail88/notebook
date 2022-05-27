<template>
	<div id="tab">

		<div>

			<table class="table">
				<tr>
					<td>
						<span class="h2">{{ currentTab.title }}</span>

						<a  href="#"
							data-bs-toggle="modal"
							:data-bs-target="tabEditUniqueID"
							data-bs-whatever="@getbootstrap">
							&nbsp;&nbsp;<i class="bi bi-pencil-square"></i>
						</a>
					</td>
					<td>
						
						<router-link to="/note/new">
							<i class="bi bi-plus-square-dotted"></i> add note
						</router-link>

					</td>
				</tr>
			</table>

			<hr>

			<TabEditForm :t=currentTab />

		</div>

		<div class="row p-2">
			<Note v-for="note in notes" v-bind:key="note.id" :note=note />
		</div>

		

	</div>
</template>

<script>

// this.$route.params.id
	import { mapGetters } from 'vuex';
	import Note from "../Note/Note";
	import NoteForm from '../Note/NoteForm';
	import TabEditForm from './TabEditForm';
	export default {
		name: "Tab",
		data() {
			return {
				currentTab: {}
			}
		},
		created() {

			this.getCurrentTab();
			this.$store.dispatch('loadNotes', this.$route.params.id);
		},
		computed: {
			...mapGetters({notes: "getNotes"}),
			tabEditUniqueID() {
				return `#tabEditModal${this.$route.params.id}`;
			},
			noteForm() {
				return `#noteModal${this.currentTab.id}`;
			}
		},
		methods: {
			getNotesForCurrentTab() {
				this.$store.dispatch('loadTabs');
				this.$store.dispatch('loadNotes', this.$route.params.id);
			},
			getCurrentTab() {
				this.$store.dispatch('loadCurrentTab', this.$route.params.id)
				.then(() => {
					this.currentTab = this.$store.getters.getCurrentTab
				})
			}
		},
		components: {
			Note,
			NoteForm,
			TabEditForm
		},
		watch: {
			"$route": ["getNotesForCurrentTab", "getCurrentTab"],
			// "$route": 
		}
	};
</script>

<style scoped>
	#tab {
		z-index: -1;
		padding: 20px;
		min-height: 1000px;
		border: 1px solid silver;
	}
	a {
		text-decoration: none;
	}
</style>
