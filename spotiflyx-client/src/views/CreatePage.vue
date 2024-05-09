<script>
import Footer from '../components/Footer.vue';
import NavBar from '../components/NavBar.vue';
import Popup from '../components/popup.vue';

export default {
    components: {
        NavBar,
        Popup,
        Footer
    },
    name: 'CreatePage'
}
</script>

<template>
    <Popup />
    <navBar />
    <main>
        <Panel class="panel">
            <template #header>
                <h1 style="font-weight: 700; font-size: 1.8rem; text-shadow: -5px 2px 20px var(--primary-color);">Créer un média :</h1>
            </template>
            <InputText v-model="name" style="width: 600px;" placeholder="Nom du média" maxlength="60"/>
            <div class="type">
                <h3>Type de média :</h3>
                <RadioButton class="btn" v-model="type" value="music" inputId="music" name="type" />
                <label for="music" class="p-radiobutton-label">Musique</label>
                <RadioButton class="btn" v-model="type" value="video" inputId="video" name="type" />
                <label for="video" class="p-radiobutton-label">Vidéo</label>
            </div>
            <div class="url">
                <InputText v-model="url" style="width: 600px;" placeholder="URL du média" maxlength="100"/>
                <div class="info">
                    <i class="pi pi-question-circle" @click="visible = true" style="margin-left: 0.8rem;"></i>
                    <Dialog v-model:visible="visible" modal header="URL du média" :style="{ width: '60%' }">
                        <h2>Voici des exemples d'URL de média</h2>
                        <h3>Youtube :</h3>
                        <ul style="list-style-type: none;">
                            <li><i class="listeIcone pi pi-times" style="color: red"></i>https://www.youtube.com/watch?v=dQw4w9WgXcQ</li>
                            <li><i class="listeIcone pi pi-check" style="color: green"></i>https://youtu.be/dQw4w9WgXcQ</li>
                            <li>Lien dans la section partage de la vidéo</li>
                        </ul>
                        <h3 style="margin-top: 2rem;">Spotify :</h3>
                        <ul style="list-style-type: none;">
                            <li><i class="listeIcone pi pi-check" style="color: green"></i>https://open.spotify.com/intl-fr/track/4PTG3Z6ehGkBFwjybzWkR8</li>
                            <li>Lien dans la section partage de la musique</li>
                        </ul>
                        <Button type="button" label="Quitter" @click="visible = false" style="margin-top: 2.5rem;"></Button>
                    </Dialog>
                </div>
            </div>
            <Button type="button" label="Créer le média" @click="createMedia()" style="margin-top: 2.5rem;"></Button>
            <p class="error"></p>
        </Panel>
    </main>
    <Footer />
</template>

<script setup>
import { ref } from 'vue';
import { ax } from '../router/router';
import { useRouter } from 'vue-router';

const visible = ref(false);
const name = ref('');
const url = ref('');
const type = ref('music');
const router = useRouter();

const createMedia = async () => {
    await ax.post('/media/create', {
        title: name.value,
        url: url.value,
        mediaType: type.value
    }).then((res) => {
        router.push('/media/' + type.value + '/' + res.data.data.id);
    }).catch((err) => {
        console.log("error:", err);
        const error = document.querySelector('.error');
        if (err.response.status == 400 || err.response.status == 422) {
            error.textContent = "Paramètre(s) invalides";
        } else {
            error.textContent = "Erreur lors de la création du média";
        }
    });
}
</script>

<style scoped>
main {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: min-content;
    min-height: 90vh;
    padding: 2.5rem 0;
    width: 100%;
    background-color: var(--color-background-soft);
}

.panel {
    width: 80%;
    height: min-content;
    margin: 2rem;
    padding: 2rem;
    box-shadow: 0 0 20px 0 var(--primary-color);
}

.type {
    display: flex;
    align-items: center;
    margin-top: 1.5rem;
}

.type .btn {
    margin-left: 2rem;
}

.type label {
    margin-left: 0.6rem;
}

.url {
    display: flex;
    align-items: center;
    margin-top: 1.5rem;
}

.info {
    display: flex;
    justify-content: center;
    align-items: center;
}

.info i {
    font-size: 1.2rem;
    color: var(--primary-color);
}

.listeIcone {
    font-size: 1rem;
    margin-right: 0.5rem;
}

.error {
    margin-top: 1rem;
    color: red;
}
</style>