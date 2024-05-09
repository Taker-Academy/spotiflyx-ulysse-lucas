<script>
import SearchBar from '../components/SearchBar.vue'
import NavBar from '../components/NavBar.vue'
import CardList from '../components/CardList.vue'
import Popup from '../components/popup.vue'
import Footer from '../components/Footer.vue'

export default {
    components: {
        SearchBar,
        NavBar,
        CardList,
        Popup,
        Footer
    },
    name: 'Home'
}
</script>

<template>
    <Popup />
    <navBar />
    <main>
        <div class="search">
            <SearchBar />
        </div>
        <div class="splitter">
            <div class="leftPanel">
                <CardList :type="video" />
            </div>
            <div class="rightPanel">
                <CardList :type="music" />
            </div>
        </div>
    </main>
    <Footer />
</template>

<script setup>
import { ref } from 'vue';

const music = ref('music');
const video = ref('video');
</script>

<style scoped>
main {
    display: flex;
    flex-direction: column;
    height: min-content;
    min-height: 90vh;
    padding: 2.5rem 0;
    width: 100%;
    background-color: var(--color-background-soft);
}

.search {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    height: min-content;
    padding: 1.5rem 4rem;
    transition-duration: .5s;
    animation-duration: .5s;
}

.splitter {
    display: flex;
    height: min-content;
    width: 100%;
    border-top: 1px solid var(--primary-color);
    margin-top: 2rem;
}

.leftPanel {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    width: 50%;
    height: 100%;
}

.rightPanel {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;
    width: 50%;
}
</style>
