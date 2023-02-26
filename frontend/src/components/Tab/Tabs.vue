<template>
	<div class="tabs">
		<a href="#"
		
		   style="margin-left: -14px;
		    width:246px;
			line-height: 33px;
			height: 45px;
			/* background: linear-gradient(silver, white, white, white); */
			"

           class="border-start-0 btn btn-primary"
           data-bs-toggle="modal"
           data-bs-target="#tabModal"
           data-bs-whatever="@getbootstrap">
		   <i class="bi bi-plus-square-dotted"></i>&nbsp;Add Tab
        </a>
		<!-- <hr class="border border-secondary opacity-75" 
			style="width: 230px; margin-top: 31px;"> -->

		<ul class="list-group" style="margin-top: 12px;">
	      	<router-link 
	      		v-for="tab in getTabs"
				style="margin-left: -14px;"
	            class=" nav-link border-start-0 rounded"
	            :to="{name: 'Tab', params: {id: tab.id} }"
				@click="setTitle"
	            v-bind:key="tab.id">
	            <li class="list-group-item border-1" style="width: 100%;">
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
			removeTab(tabId) {
				const check = confirm("Are you sure?");
				if(check === false) {
					return false;
				}
				this.$store.dispatch('removeTab', tabId)
				.then(() => {
					this.$store.dispatch('loadTabs');
					this.$router.push('/');
				});
			},
			setTitle() {
				setTimeout(() => {
					const currentTab = JSON.parse(localStorage.getItem('tab'));
					document.querySelector('#title-' + currentTab.id).value = currentTab.title;
				}, 100);
			}
		},
		watch: {
			'$route': 'setTitle'
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
