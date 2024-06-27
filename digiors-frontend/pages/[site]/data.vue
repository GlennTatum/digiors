<template>
<v-container>
    <v-sheet
    class="pa-8"
    >
        <p class="text-h4">Data Submission</p>
        <p class="text-h6">{{ site.name }}</p>
        <v-form @submit.prevent>
            <v-text-field
            v-model="name"
            label="School Organization/Name"
            :rules="[v => !!v || 'Name is required']"
            required
            ></v-text-field>

            <v-text-field
            v-model="name"
            label="Email of Primary Steward"
            :rules="[v => !!v || 'Email is required']"
            required
            ></v-text-field>

            <div class="pb-2">Name of Site</div>
            <v-select
                :items="registrations"
                item-title="date_fmt"
                item-value="id"
                v-model="registration"
                required
            ></v-select>

            <v-text-field
            v-model="name"
            label="ORS Tag Number"
            :rules="[v => !!v || 'Email is required']"
            required
            ></v-text-field>

            <v-file-input
                accept="image/*"
                multiple
                label="Data Sheets"
            ></v-file-input>


            <v-btn type="submit" @click="submit" block class="mt-2" color="blue">Submit</v-btn>
        </v-form>
    </v-sheet>
</v-container>
</template>

<script>
export default {
    setup () {
        const route = useRoute()
        return {
            route
        }
    },
    data() {
        return {
            site: {}
        }
    },
    methods: {
        async fetchSite() {
            fetch(`http://127.0.0.1:8000/api/stations/${this.route.params.site}`, {
                method: "GET"
            })
            .then((response) => {
                response.json()
                .then((data) => {
                    this.site = data
                })
            })
        },
    },
    created() {
        this.fetchSite()
    }
}
</script>