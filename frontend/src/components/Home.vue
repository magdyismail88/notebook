<template>
  <div id="home">
  	Navigate notes accross tabs

  </div>
</template>

<script>

	export default {
	 	name: 'Home',
	 	mounted() {
	 		this.$store.dispatch('loadTabs');
		}
  };
</script>
