<template>
    <div class="main">
        <div class="searchBar">
            <div class="groupe">
                <inputText class="bar" v-model="title" @input="search" placeholder="Rechercher un média"/>
                <div v-if="mediaLs.length == 0 && title.length >= 2" class="result">
                    <div class="media">
                        <h2>Aucun résulats</h2>
                    </div>
                </div>
                <div class="result" v-else-if="mediaLs.length > 0">
                    <RouterLink class="media" :to="'/media/' + media.mediaType + '/' + media.id" v-for="media in mediaLs" :key="media">
                        <h2>{{ media.title }}</h2><h3>{{ media.mediaType }}</h3>
                    </RouterLink>
                </div>
            </div>
            
        </div>
    </div>
</template>

<script setup>
import { ref } from "vue";
import { ax } from "../router/router";

const title = ref('');
const mediaLs = ref([]);

const search = async (event) => {
    if (title.value.length < 2) {
        mediaLs.value = [];
        return;
    }
    await ax.post("/media/search", {
        search: title.value,
    }).then((response) => {
        mediaLs.value = response.data.data;
    }).catch((error) => {
        console.log(error);
    })
}
</script>

<style scoped>
a {
    text-decoration: none;
    color: var(--color-text);
}

a:hover {
    background-color: #444444;
}

.main {
    height: 100%;
    width: 100%;
    background-color: var(--color-background-soft);
    display: flex;
    justify-content: center;
    align-items: center;
}

.searchBar {
    width: 600px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}

.bar {
    width: 100%;
    height: 50px;
    border: 1px solid var(--primary-color);
    border-radius: 25px;
    z-index: 10;
}

.groupe {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    position: relative;
}

.result {
    width: 100%;
    height: min-content;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding-top: 25px;
    transform: translate(0, -25px);
    border: 1px solid var(--primary-color-soft);
    border-top: none!important;
    border-bottom-left-radius: 25px;
    border-bottom-right-radius: 25px;
    z-index: 5;
    position: sticky;
}

.media:first-child {
    border: none;
}

.media:last-child {
    border-bottom-left-radius: 25px;
    border-bottom-right-radius: 25px;
}

.media {
    display: flex;
    justify-content: center;
    border-top: 1px solid var(--primary-color-soft);
    width: 100%;
    display: flex;
    justify-content: space-between;
    padding: 10px;
}

h2 {
    font-size: 1.5rem;
    font-weight: 700;
    height: 100%;
}

h3 {
    font-size: 1.2rem;
    font-weight: 500;
    height: 100%;
}
</style>