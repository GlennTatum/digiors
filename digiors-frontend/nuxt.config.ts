// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({

  devtools: { enabled: true },
  css: [
    'vuetify/lib/styles/main.sass',
  ],
  build: {
    transpile: ['vuetify', '@vuepic/vue-datepicker'],
  },
  vite: {
    define: {
      'process.env.DEBUG': false,
    },
  },
})
