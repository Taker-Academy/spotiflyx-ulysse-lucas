<script>
import NavBar from '../components/NavBar.vue'

export default {
    components: {
        NavBar
    },
    name: 'ErrorPage'
}
</script>
<template>
    <navBar />
    <main>
        <h1>404 page not found</h1>
        <h3>There is nothing here</h3>
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
