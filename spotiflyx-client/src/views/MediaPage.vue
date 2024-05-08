<script>
import NavBar from '../components/NavBar.vue'

export default {
    components: {
        NavBar
    },
    name: 'Media'
}
</script>

<template>
    <navBar />
    <main>
        <h1>media: {{ id }}</h1>
    </main>
</template>

<script setup>
import { ref } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();
// get id from route after /media/
const id = ref(route.params.id);
</script>

<style scoped>
main {
    display: flex;
    flex-direction: column;
    height: 100vh;
    width: 100%;
    background-color: var(--color-background-soft);
}
</style>
