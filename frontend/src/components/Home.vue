<template>
  <div id="home">
  	No tab selected
  </div>
</template>

<script>
	export default {
	 	name: 'Home',
	 	mounted() {
			const currentContainer = JSON.parse(
				localStorage.getItem('container')
			)
			if(currentContainer && currentContainer.id !== '') {
				this.$store.dispatch('loadTabs');
			}
		}
  };
</script>
