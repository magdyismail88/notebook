<template>
  <nav class="navbar navbar-expand-lg navbar-light bg-light fixed-top shadow-sm">
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
                class="text-white"
                data-bs-toggle="modal"
                data-bs-target="#containerEditModal"
                data-bs-whatever="@getbootstrap"
                @click="preparing()"
               >
              <i class="bi bi-pencil-square"></i>
            </a>
          </span> 
        </li>

        &nbsp;&nbsp;

        <li class="nav-item">
          <a href="#"
           class="nav-link text-primary"
           data-bs-toggle="modal"
           data-bs-target="#containerChangeModal"
           data-bs-whatever="@getbootstrap">
           <i class="bi bi-caret-down"></i>&nbsp;change
          </a>
        </li>


        <li class="nav-item">
          <a href="#"
           class="nav-link text-success"
           data-bs-toggle="modal"
           data-bs-target="#containerModal"
           data-bs-whatever="@getbootstrap">
           <i class="bi bi-plus-circle-dotted"></i>&nbsp;add
          </a>
        </li>


        <li class="nav-item">
          <a href="#" class="nav-link text-danger" @click="deleteCurrentContainer">
            <i class="bi bi-x-circle"></i>&nbsp;delete
          </a>
        </li>


      </ul>
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
            localStorage.removeItem('tab')
            this.$store.dispatch('loadContainers');
            this.$store.dispatch('setDefaultContainer')
            this.$store.dispatch('changeContainer', 0);
            this.$router.push('/');
          })
          .catch((err) => console.log(err));
        }

      },
      preparing() {
        console.log('Prepare for edit container');
      }
    },
    computed: {
      selectedContainer() {
        return this.$store.state.container.container.title || "containerless"
      }
    }
  };
</script>
