<template>
<v-container style="max-width: 75vw;">
    <div>
        <div class="pt-4"></div>
        <v-img
        :src="site.image"
        style="max-height: 80vh;"
        ></v-img>
        <div class="pt-4"></div>

        <p class="text-h4">{{ site.name }}</p>

        <div class="pt-2"></div>

        <p class="font-weight-light">{{ site.geo }}</p>

        <div class="pt-4"></div>
        <span>
            <v-btn
            color="blue"
            >
            <NuxtLink :to="`register`" style="text-decoration: none; color: white;">
                <p>Register</p>
            </NuxtLink>
            </v-btn>
        </span>
        <span class="pa-2"></span>
        <span>
            <v-btn
            color="purple"
            >
            <NuxtLink :to="`signup`" style="text-decoration: none; color: white;">
                <p>Sign Up</p>
            </NuxtLink>
            </v-btn>
        </span>

        <div class="pt-4"></div>
        {{ site.description }}
        <div class="pt-4"></div>
        <p class="text-h4">Information</p>
        <div class="pt-4"></div>
        <v-table
            fixed-header
        >
            <thead>
            <tr>
                <th class="text-left">
                    Date
                </th>
                <th class="text-left">
                    Lead Ambassador
                </th>
                <th class="text-left">
                    Ambassadors
                </th>
            </tr>
            </thead>
            <tbody>
            <tr
            v-for="event in site.registrations"
            >
                <td>
                    {{ `${event.dateBegin_fmt} - ${event.dateEnd_fmt}` }}
                </td>
                <td>
                    {{ event.steward }}
                </td>
                <td>
                <span v-for="member in event.users">
                    {{ `${member.name} ` }}
                </span>
                </td>
            </tr>
            </tbody>
        </v-table>
    </div>
</v-container>
</template>



<script>
// import {mdiHome} from '@mdi/js'
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
            dialog: false,
            isEditingDescription: false
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
        }
    },
    created() {
        this.fetchSite()
    }
}
</script>