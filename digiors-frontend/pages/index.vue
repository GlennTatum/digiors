<template>
    <div class="d-flex flex-wrap pt-8">
        <LocationCard
        v-for="site in sites"
        :image="site.image"
        :name="site.name"
        :slug="site.slug"
        :geo="site.geo"
        :status="site.status"
        ></LocationCard>
    </div>
</template>
  
<script>
  export default {
    data() {
      return {
        sites: {}
      }
    },
    methods: {
      async fetchSites() {
        fetch("http://127.0.0.1:8000/api/stations", {
          method: "GET"
        })
        .then((response) => {
          response.json()
          .then((data) => {
            this.sites = data
          })
        })
      }
    },
    created() {
      this.fetchSites()
    }
  }
</script>
  