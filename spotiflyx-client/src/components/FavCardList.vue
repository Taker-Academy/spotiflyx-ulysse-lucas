<template>
    <Suspense>
        <template #default>
            <div class="cardList">
                <template v-if="cards.length > 0">
                    <RouterLink :to="'/media/' + card.mediaType + '/' + card.id" class="card" v-for="card in cards" :key="card.id">
                        <div class="img" :style="{ backgroundImage: `url(${card.imgUrl})`}" alt="media" />
                        <div class="cardInfo">
                            <h3>{{ card.title }}</h3>
                        </div>
                    </RouterLink>
                </template>
                <template v-else>
                    <h3>Aucune {{ mediaType }} à afficher</h3>
                </template>
            </div>
        </template>
        <template #fallback>
            <div class="cardList">
                <h2>Loading ...</h2>
            </div>
        </template>
    </Suspense>
</template>

<script setup>
import { ref, onMounted, defineProps } from 'vue';
import { ax } from '../router/router';
import { RouterLink } from 'vue-router';

ax.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token');

const props = defineProps({
    type: String
});

const mediaType = ref('');
if (props.type == "music") {
    mediaType.value = 'musique';
} else {
    mediaType.value = 'vidéo';
}

const cards = ref([]);
onMounted(async () => {
    try {
        const data = await ax.get('/me/favorites');
        if (data.data.data != null) {
            for (let i = 0; i < data.data.data.length; i++) {
                if (props.type == data.data.data[i].mediaType) {
                    cards.value.push(data.data.data[i]);
                }
            }
        }
    } catch (error) {
        console.log(error);
    }
    console.log("cards: ", cards.value);
});
</script>

<style scoped>
.cardList {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-evenly;
    align-items: center;
    width: 100%;
    height: 100%;
    padding: 1rem;
    background-color: var(--color-background-soft);
}

.title {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: 1rem;
}

.title h2 {
    font-size: 2rem;
    font-weight: 700;
}

.card {
    width: 250px;
    height: 300px;
    margin: 1rem;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    align-items: center;
    text-decoration: none;
    color: var(--text-color);
    border-radius: 10px;
    background-color: var(--color-background-soft);
    box-shadow: 0px 0px 20px 0px var(--primary-color);
}

.cardInfo {
    padding: 10px;
    align-items: center;
    justify-content: center;
}

.img {
    width: 100%;
    height: 70%;
    background-size: cover;
    background-position: center;
    border-radius: 10px 10px 0 0;
}
</style>