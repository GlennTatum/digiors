<template>
    <v-container>
        <v-sheet
        class="pa-8"
        >
            <p class="text-h4">Choose an Event</p>
            <p class="text-h6">{{ site.name }}</p>
            <v-form @submit.prevent>
                <v-text-field
                v-model="name"
                label="Name"
                :rules="[v => !!v || 'Name is required']"
                required
                ></v-text-field>

                <v-select
                :items="registrations"
                item-title="dateBegin_fmt"
                item-value="id"
                v-model="registration"
                required
                ></v-select>

                <v-input
                v-model="registration"
                :rules="[v => !!v || 'Date is required']"
                required
                >
                </v-input>
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
            site: {},
            registrations: [],
            name: null,
            registration: null,
            select: null,
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
                    this.registrations = data.registrations
                })
            })
        },
        async submit() {
            if (this.name != null && this.registration != null) {
                fetch(`http://127.0.0.1:8000/api/stations/${this.route.params.site}/registrations/${this.registration}/users`, {
                    method: "POST",
                    body: JSON.stringify({
                        "name": this.name,
                    })
                })
                navigateTo(`/${this.route.params.site}/info`)
            }
        },
    },
    created() {
        this.fetchSite()
    }
}
</script>