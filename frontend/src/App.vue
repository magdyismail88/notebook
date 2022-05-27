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
import VueFroala from 'vue-froala-wysiwyg';
import Header from "./components/partials/Header";
import Tabs from "./components/Tab/Tabs";
import ContainerForm from './components/Container/ContainerForm';
import ContainerChangeForm from './components/Container/ContainerChangeForm';
import ContainerEditForm from './components/Container/ContainerEditForm';

export default {
  name: 'App',
  components: {
    Header,
    Tabs,
    ContainerForm,
    ContainerChangeForm,
    ContainerEditForm
  },
  created() {
    this.$store.dispatch('setDefaultContainer');
    this.$store.dispatch('loadContainers');
    this.$store.dispatch('loadTabs');
  }
};
</script>

<style>
  .app-tabs {
    min-height: 1000px;
  }
</style>
