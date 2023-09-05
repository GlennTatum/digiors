<template>
    <v-app>
        <v-navigation-drawer
        v-model="drawer"
        temporary
        >
        <v-list nav>
                <v-list-item 
                :prepend-icon="mdiHome" 
                title="Home" 
                value="home"
                @click="navigateTo('/')"
                >
                </v-list-item>            
        </v-list>
        </v-navigation-drawer>

        <v-app-bar>
        <v-app-bar-nav-icon  @click="drawer = !drawer"></v-app-bar-nav-icon>

        <v-app-bar-title>Billion Oyster Project Digital Platform</v-app-bar-title>
        </v-app-bar>


        <v-main class="pl-8 pr-8">
            <slot />
        </v-main>
    </v-app>
</template>

<script setup>
import {mdiHome} from '@mdi/js'
</script>

<script>
export default {
    data() {
        return {
            drawer: false
        }
    }
}
</script>