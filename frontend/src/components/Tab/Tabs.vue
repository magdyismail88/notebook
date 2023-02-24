<template>
	<div class="tabs">
		<a href="#"
		   style="margin-left: -14px; width:246px; line-height: 33px; height: 45px; background-color: #34495e;"
           class="nav-link border-start-0 btn text-white border-white rounded-pill"
           data-bs-toggle="modal"
           data-bs-target="#tabModal"
           data-bs-whatever="@getbootstrap">
		   <i class="bi bi-plus-square-dotted"></i>&nbsp;Add Tab
        </a>
		<hr class="border border-primary border-1 opacity-75" 
			style="width: 230px; margin-top: 50px;">

		<ul class="list-group">
	      	<router-link 
	      		v-for="tab in getTabs"
				style="margin-left: -14px;"
	            class="nav-link border-start-0 rounded"
	            :to="{name: 'Tab', params: {id: tab.id} }"
	            v-bind:key="tab.id">
	            <li class="list-group-item shadow-sm" style="width: 100%;">
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
