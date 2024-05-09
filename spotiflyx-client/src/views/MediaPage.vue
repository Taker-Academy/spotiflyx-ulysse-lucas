<script>
import NavBar from '../components/NavBar.vue'
import Footer from '../components/Footer.vue'

export default {
    components: {
        NavBar,
        Footer
    },
    name: 'Media'
}
</script>

<template>
    <NavBar />
    <main>
        <div class="video iframShadow" v-if="media.mediaType == 'video'">
            <iframe width="560" height="315" :src="media.url" title="media" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
        </div>
        <div class="music iframShadow" v-else>
            <div id="embed-iframe"><h2>Chargement . . . (si le chargement persiste, rechargez la page)</h2></div>
        </div>
        <div class="mediaInfos">
            <h1>{{ media.title }}</h1>
            <h2>{{ media.author }}</h2>
            <div class="interact">
                <Button @click="Like()" class="like">
                    <h2>{{ media.likes }} Likes</h2>
                    <i v-if="media.liked == false" class="pi pi-thumbs-up"></i>
                    <i v-else class="pi pi-thumbs-up-fill"></i>
                </Button>
                <Button @click="Favoris()" class="favorite">
                    <i v-if="media.favorite == false" class="pi pi-bookmark"></i>
                    <i v-else class="pi pi-bookmark-fill"></i>
                </Button>
            </div>
            
        </div>
    </main>
    <Footer />
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { ax } from '../router/router';

const route = useRoute();
const router = useRouter();
const id = ref(route.params.id);
const type = ref(route.params.type);

ax.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token');

const media = ref({});
onMounted(async () => {
    try {
        const data = await ax.get('/media/' + type.value + '/' + id.value);
        media.value = data.data.data;
    } catch (error) {
        console.log(error);
        router.push('/error');
    }
    console.log("media: ", media.value);
    if (type.value == 'music') {
        window.onSpotifyIframeApiReady = (IFrameAPI) => {
            const element = document.getElementById('embed-iframe');
            const options = {
                uri: media.value.url
                };
            const callback = (EmbedController) => {};
            IFrameAPI.createController(element, options, callback);
        };
    }
});


const Favoris = () => {
    if (media.value.favorite == false) {
        addFavori();
    } else {
        rmFavori();
    }
}

const Like = () => {
    if (media.value.liked == false) {
        addLike();
    } else {
        rmLike();
    }
}

const addLike = async () => {
    try {
        var log = await ax.post('/me/like/' + id.value);
        console.log(log);
        media.value.likes += 1;
        media.value.liked = true;
    } catch (error) {
        console.log(error);
    }
}

const rmLike = async () => {
    try {
        var log = await ax.delete('/me/like/' + id.value);
        console.log(log);
        media.value.likes -= 1;
        media.value.liked = false;
    } catch (error) {
        console.log(error);
    }
}

const addFavori = async () => {
    try {
        var log = await ax.post('/me/save/' + id.value);
        console.log(log);
        media.value.favorite = true;
    } catch (error) {
        console.log(error);
    }
}

const rmFavori = async () => {
    try {
        var log = await ax.delete('/me/save/' + id.value);
        console.log(log);
        media.value.favorite = false;
    } catch (error) {
        console.log(error);
    }
}
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

.mediaInfos {
    display: flex;
    flex-direction: column;
    justify-content: center;
    width: 100%;
    height: min-content;
    padding: 2rem;
}

.interact {
    display: flex;
    justify-content: start;
    width: 100%;
    margin-top: 1rem;
}

.like {
    min-width: min-content;
    display: flex;
    align-items: center;
}

.like i {
    margin-left: 1rem;
    font-size: 1.8rem;
}

.favorite {
    display: flex;
    align-items: center;
    justify-content: center;
    margin-left: 2rem;
    min-width: min-content;
}

.favorite i {
    margin: 0 1.5rem;
    font-size: 1.8rem;
}

.video {
    width: 100%;
    padding: 2rem;
    padding-top: 3rem;
    height: min-content;
    border-radius: 10px;
}

.video.iframeShadow iframe {
    height: 60vh;
    width: auto;
}

.music {
    width: 100%;
    padding: 2rem;
    padding-top: 3rem;
    height: min-content;
    border-radius: 10px;
}
</style>
