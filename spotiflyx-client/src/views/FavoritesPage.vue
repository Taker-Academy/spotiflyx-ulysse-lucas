<script>
import NavBar from '../components/NavBar.vue'
import Footer from '../components/Footer.vue'
import FavCardList from '../components/FavCardList.vue'

export default {
    components: {
        NavBar,
        Footer
    },
    name: 'FavoritesPage'
}
</script>

<template>
    <navBar />
    <main>
        <h1>Mes favoris</h1>
        <TabView class="ls">
            <TabPanel header="Musiques">
                <FavCardList :type="'music'"/>
            </TabPanel>
            <TabPanel header="Vidéos">
                <FavCardList :type="'video'"/>
            </TabPanel>
        </TabView>
    </main>
    <Footer />
</template>

<script setup>
import { ref } from 'vue';

const music = ref('');
</script>

<style scoped>
main {
    display: flex;
    flex-direction: column;
    height: min-content;
    min-height: 90vh;
    padding: 2.5rem 0;
    width: 100%;
    background-color: #18181b;
}

main h1 {
    color: white;
    text-align: center;
    margin: 1rem 0;
    font-size: 2.8rem;
    font-weight: 700;
    text-shadow: -5px 2px 20px var(--primary-color);
}

.ls {
    display: flex;
    flex-direction: column;
    width: 100%;
    height: max-content;
}
</style>