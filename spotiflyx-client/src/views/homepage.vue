<script>
import SearchBar from '../components/SearchBar.vue'
import NavBar from '../components/NavBar.vue'

export default {
    components: {
        SearchBar,
        NavBar
    },
    name: 'Home'
}
</script>

<template>
    <navBar />
    <main>
        <h1>Homepage</h1>
        <h3>bienvuenue</h3>
        <searchBar />
    </main>
</template>

<style scoped>
main {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    height: 100vh;
    width: 100%;
    background-color: var(--color-background-soft);
}
</style>
