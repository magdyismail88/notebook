<template>
  <div id="app">
    <Header />
    <ContainerForm />
    <ContainerChangeForm />
    <ContainerEditForm />
    <div class="row" style="margin-top: 100px !important;">
      <div class="col-2 app-tabs">
        <Tabs />
      </div>
      <div class="col-10">
        <router-view />
      </div>
    </div>
  </div>
</template>
<script>
// import VueFroala from 'vue-froala-wysiwyg';
import Header from "./components/partials/Header";
import Tabs from "./components/Tab/Tabs";
import ContainerForm from './components/Container/ContainerForm';
import ContainerChangeForm from './components/Container/ContainerChangeForm';
import ContainerEditForm from './components/Container/ContainerEditForm';

export default {
  name: 'App',
  data() {
    return {
      currentContainer: '',
      currentPath: ''
    }
  },
  components: {
    Header,
    Tabs,
    ContainerForm,
    ContainerChangeForm,
    ContainerEditForm
  },
  created() {
    const vm = this
    this.currentContainer = localStorage.getItem('container')
    this.currentPath = this.$route.path
    setTimeout(() => {
      // vm.$store.dispatch('setDefaultContainer');
      vm.$store.dispatch('loadContainers');
      vm.$store.dispatch('loadTabs');
    }, 1000)
    document.addEventListener('DOMContentLoaded', function() {
      setTimeout(() => {
        const currentTab = JSON.parse(localStorage.getItem('tab')) || null;
        if(currentTab) {
          vm.$router.push('/tab/' + currentTab.id)
        }
      }, 1000)
        // setTimeout(function() {
        //   vm.$router.push(vm.currentPath)
        //   // const containerID = JSON.parse(vm.currentContainer).id
        //   // this.$store.dispatch("changeContainer", containerID)
        //   // localStorage.setItem('container', vm.currentContainer)
        //   // vm.$store.dispatch('changeContainer')
        //   // vm.$store.dispatch('loadTabs');
        // }, 1000)
    }, false);
  }
};
</script>

<style>
  .app-tabs {
    min-height: 1000px;
  }

  body {
    background-color: #ecf0f1;
  }
</style>
