<template>
  <div
      class="flex-grow-1 ma-2"
      style="max-width: 360px; min-width: 360px"
  >
      <v-card
        class="rounded-lg"
        @click="navigateTo(`${slug}/info`)"
      >
        <v-img
          :src="image"
          style="max-width: 360px; max-height: 240px;"
          cover
        ></v-img>
        <v-card-item>
          <div class="text-h6">
            <p class="overflow-x-auto">{{ name }}</p>
          </div>
        </v-card-item>

        <v-card-subtitle>
          <p>{{ geo }}</p>
        </v-card-subtitle>

        <v-card-item>
          <div v-if="status == 'open'">
            <div class="pa-2 rounded-lg" style="background-color: #263226;">
              <p class="font-weight-medium text-body-2" style="color: #4CAF50;">Open</p>
            </div>
          </div>
          <div v-if="status == 'closed'">
            <div class="pa-2 rounded-lg" style="background-color: #321C20;">
              <p class="font-weight-medium text-body-2" style="color: #d4062c;">Closed</p>
            </div>
          </div>
          <div v-if="status == 'support'">
            <div class="pa-2 rounded-lg" style="background-color: #3C2E1C;">
              <p class="font-weight-medium text-body-2" style="color: #FB8C00;">Maintenance Required</p>
            </div>
          </div>
        </v-card-item>

        <v-card-text>
        </v-card-text>

        <v-card-actions>
        </v-card-actions>
      </v-card>
  </div>
</template>


<script>
export default {
    props: {
        image: String,
        name: String,
        slug: String,
        geo: String,
        status: String
    },
}
</script>