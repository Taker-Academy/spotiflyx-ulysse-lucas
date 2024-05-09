<script>
import NavBar from '../components/NavBar.vue'
import Footer from '../components/Footer.vue'

export default {
    components: {
        NavBar,
        Footer
    },
    name: 'ErrorPage'
}
</script>
<template>
    <navBar />
    <main>
        <h1>404 page non trouvé</h1>
        <h3>Il n'y a rien ici</h3>
    </main>
    <Footer />
</template>

<style scoped>
main {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    height: 90vh;
    width: 100%;
    background-color: var(--color-background-soft);
}
</style>
