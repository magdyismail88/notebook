<template>
	<div class="tabs">

		<nav class="nav flex-colum">

	      	<router-link 
	      		v-for="tab in getTabs"
	            class="nav-link"
	            :to="{name: 'Tab', params: {id: tab.id} }"
	            v-bind:key="tab.id">
	            <h5>{{ tab.title }}&nbsp;
	            <a href="#" @click="removeTab(tab.id)"><span>&times;</span></a></h5>
	      	</router-link>

	    </nav>

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
				// event.preventDefault();

				let check = confirm("Are you sure?");

				if(check === false) {
					return false;
				}

				this.$store.dispatch('removeTab', tab_id)
				// .then(() => this.$route.push('/'))
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
