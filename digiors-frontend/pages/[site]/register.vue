<template>
<v-container>
    <v-sheet
    class="pa-8"
    >
        <p class="text-h4">Lead Ambassador Site Visit Registration</p>
        <p class="text-h6">{{ site.name }}</p>
        <v-form @submit.prevent>
            <v-text-field
            v-model="name"
            label="Name"
            :rules="[v => !!v || 'Name is required']"
            required
            ></v-text-field>

            <div>Event Begin</div>
            <VueDatePicker 
            v-model="start_date"
            :dark="true"
            :is-24="false"
            ></VueDatePicker>

            <v-input
            v-model="start_date"
            :rules="[v => !!v || 'Date is required']"
            required
            >
            {{ start_date }}
            </v-input>

            <div>Event End</div>
            <VueDatePicker 
            v-model="end_date"
            :dark="true"
            :is-24="false"
            ></VueDatePicker>

            <v-input
            v-model="end_date"
            :rules="[v => !!v || 'Date is required']"
            required
            >
            {{ end_date }}
            </v-input>

            <v-btn type="submit" @click="submit" block class="mt-2" color="blue">Submit</v-btn>
        </v-form>
    </v-sheet>
</v-container>
</template>

<script>
import VueDatePicker from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css'


export default {
    setup () {
        const route = useRoute()
        return {
            route
        }
    },
    components: { 
        VueDatePicker
    },
    data() {
        return {
            site: {},

            name: null,
            start_date: null,
            end_date: null,
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
                })
            })
        },
        async submit() {
            console.log("sub")
            if (this.name != null && this.start_date != null && this.end_date != null) {
                fetch(`http://127.0.0.1:8000/api/stations/${this.route.params.site}/registrations`, {
                    method: "POST",
                    body: JSON.stringify({
                        "steward": this.name,
                        "dateBegin": this.start_date,
                        "dateEnd": this.end_date
                    })
                })
                navigateTo(`/${this.route.params.site}/info`)
            }
        }
    },
    created() {
        this.fetchSite()
    }
}
</script>