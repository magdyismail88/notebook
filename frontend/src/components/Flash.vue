<template>
  <div class="alert" style="display: none;"
    v-bind:class="[ status ? 'div-block' : '', isError ? 'alert-danger' : 'alert-success']">
    {{ msg }}
  </div>
</template>

<script>
  export default {
    name: "Flash",
    props: ["status", "isError", "msg"]
  };
</script>
