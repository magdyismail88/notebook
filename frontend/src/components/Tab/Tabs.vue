<template>
	<div class="tabs">
		<a href="#"
		   style="margin-left: 5px;"
           class="nav-link"
           data-bs-toggle="modal"
           data-bs-target="#tabModal"
           data-bs-whatever="@getbootstrap">
		   <i class="bi bi-plus-square-dotted"></i>&nbsp;Add Tab
        </a>
		<hr class="border border-primary border-1 opacity-75">

		

		<ul class="list-group">
	      	<router-link 
	      		v-for="tab in getTabs"
	            class="nav-link"
	            :to="{name: 'Tab', params: {id: tab.id} }"
	            v-bind:key="tab.id">
	            <li class="list-group-item" style="width: 100%;">
					<p style="padding: 5px;  margin: 0;">{{ tab.title }}&nbsp; </p>
				</li>
	            
	      	</router-link>
	    </ul>
	    <TabForm />
	</div>
</template>

<script>
	import { mapGetters } from 'vuex';
	import TabForm from './TabForm'
	export default {
		name: "Tabs",
		computed: {
			...mapGetters(['getTabs'])
		},
		components: {
			TabForm
		},
		mounted() {
			this.$store.dispatch('loadTabs');
		},
		methods: {
			removeTab(tab_id) {
				const check = confirm("Are you sure?");
				if(check === false) {
					return false;
				}
				this.$store.dispatch('removeTab', tab_id)
				.then(() => {
					this.$store.dispatch('loadTabs');
					this.$router.push('/');
				})
			}
		}
	};
</script>

<style scoped>
	.tabs {
		z-index: 999999;
	}

	a {
		text-decoration: none;
	}

</style>
