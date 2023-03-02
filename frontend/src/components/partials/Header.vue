<template>
  <nav id="top-navbar"
    class="navbar navbar-expand-lg navbar-light bg-light fixed-top shadow-sm" 
    data-bs-theme="dark">
  <!-- <nav class="navbar navbar-expand-lg navbar-dark bg-dark fixed-top shadow-sm"> -->
  <div class="container-fluid">
    <router-link class="navbar-brand" to="/">
      <i class="h3 bi bi-journals"></i>&nbsp;
      <!-- <strong>Notebook</strong> -->
    </router-link>


    <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarSupportedContent">
      <span>Container</span>
      <ul class="navbar-nav me-auto mb-2 mb-lg-0">
        <li class="nav-item">
          <span class="nav-link">
            <i class="bi bi-chevron-bar-right"></i> {{ selectedContainer }}
            <a  href="#"
                id="container-edit-btn"
                data-bs-toggle="modal"
                data-bs-target="#containerEditModal"
                data-bs-whatever="@getbootstrap"
                @click="prepareContainerEdit()"
               >
              <i class="bi bi-pencil-square"></i>
            </a>
          </span> 
        </li>

      </ul>

      <span class="navbar-text">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <a href="#"
            class="nav-link text-primary"
            data-bs-toggle="modal"
            data-bs-target="#containerChangeModal"
            data-bs-whatever="@getbootstrap">
            <i class="bi bi-caret-down"></i>&nbsp;Change
            </a>
          </li>

        <li class="nav-item">
          <a href="#"
           class="nav-link text-primary"
           data-bs-toggle="modal"
           data-bs-target="#containerModal"
           data-bs-whatever="@getbootstrap">
           <i class="bi bi-plus-square-dotted"></i>&nbsp;Add
          </a>
        </li>

        <li class="nav-item">
          <a href="#" id="container-remove-btn" class="nav-link text-danger" @click="deleteCurrentContainer">
            <i class="bi bi-trash"></i>&nbsp;Remove
          </a>
        </li>
        </ul>
      </span>

    </div>
  </div>
</nav>
</template>

<script>
  export default {
    name: "Header",
    methods: {
      deleteCurrentContainer() {
        const check = confirm("Are you sure ?");
        if(check) {
          this.$store.dispatch('deleteContainer')
            .then(() => {
              localStorage.removeItem('tab');
              this.$store.dispatch('loadContainers');
              this.$store.dispatch('setDefaultContainer');
              this.$store.dispatch('changeContainer', 0);
              document.querySelector('#tabs').style.display = 'none'
              document.querySelector('#container-edit-btn').style.display = 'none'
              document.querySelector('#container-remove-btn').style.display = 'none'
              this.$router.push('/');
            })
            .catch((err) => console.log(err));
        }

      },
      prepareContainerEdit() {
        setTimeout(() => {
          const currentContainer = JSON.parse(localStorage.getItem('container'));
          document.querySelector('#container-title-edit').value = currentContainer.title;
        }, 100)
      }

    },
    computed: {
      selectedContainer() {
        return this.$store.state.container.container.title || "containerless"
      }
    }
  };
</script>
<style scoped>
  #top-navbar {
    background: linear-gradient(#FFF, #FFF, #e3f2fd)
  }
</style>