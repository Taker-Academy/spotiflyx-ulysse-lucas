<script>
import Popup from '../components/popup.vue';
import Footer from '../components/Footer.vue';
import NavBar from '../components/NavBar.vue';

export default {
    components: {
        NavBar,
        Popup,
        Footer
    },
    name: 'Account'
}
</script>

<template>
    <Popup />
    <navBar />
    <main>
        <div>
            <div class="top">
                <ConfirmDialog></ConfirmDialog>
                <Fieldset class="field" legend="Mon compte">
                    <h2>Email: {{ email }}</h2>
                    <Button @click="logout()">Se déconnecter</Button>
                    <Button @click="confirmDelete()" severity="danger">Supprimer mon compte</Button>
                </Fieldset>
            </div>
            <div class="passwordForm">
                <Toast />
                <Fieldset legend="Changement de mot de passe">
                    <div class="pwd">
                        <Password class="pwdInput" v-model="old" :feedback="false" placeholder="Ancien mot de passe" style="width: 100%;" toggleMask/>
                        <Password class="pwdInput" v-model="newpwd" :feedback="false" placeholder="Nouveau mot de passe" style="width: 100%;" toggleMask/>
                    </div>
                    <div class="save">
                        <Button @click="savePwd()">Sauvegarder</Button>
                        <p class="errorMsg"></p>
                    </div>
                </Fieldset>
            </div>
        </div>
    </main>
    <Footer />
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useConfirm } from "primevue/useconfirm";
import { ax } from '../router/router';
import { useToast } from 'primevue/usetoast';
import Toast from 'primevue/toast';
import Button from 'primevue/button';

const toast = useToast();
const confirm = useConfirm();
const router = useRouter();

const email = ref('test@gmail.com');
const old = ref('');
const newpwd = ref('');

ax.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token');
ax.get('/user').then((res) => {
    email.value = res.data.data.email;
}).catch((err) => {
    console.log(err);
});

const confirmDelete = () => {
    confirm.require({
        message: ' Est tu certain de vouloir supprimer ton compte ?',
        header: 'Confirmation',
        icon: 'pi pi-exclamation-triangle',
        rejectClass: 'p-button-secondary p-button-outlined',
        acceptClass: 'p-button-danger',
        rejectLabel: 'Annuler',
        acceptLabel: 'Supprimer',
        accept: () => {
            deleteAccount();
        },
        reject: () => {}
    });
};

const logout = () => {
    localStorage.removeItem('token');
    router.push('/signin');
}

const deleteAccount = async () => {
    localStorage.removeItem('token');
    await ax.delete('/user').then((res) => {
        router.push('/signup');
    }).catch((err) => {
        console.log(err);
    });
}

const savePwd = async () => {
    await ax.put('/user', {
        oldPassword: old.value,
        newPassword: newpwd.value
    }).then((res) => {
        old.value = '';
        newpwd.value = '';
        toast.add({severity:'success', summary: ' Mot de passe modifié', detail: 'Votre mot de passe a bien été modifié', life: 6000});
    }).catch((err) => {
        console.log(err);
        const errorMsg = document.querySelector('.errorMsg');
        if (errorMsg) {
            errorMsg.textContent = 'Mot de passe incorrect';
        }
        old.value = '';
        newpwd.value = '';
    });
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

.field {
    padding: 1rem 2rem 2rem 2rem;
    width: 500px;
}

.field button {
    text-align: center;
}

.top {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    width: 100%;
    height: min-content;
    margin-top: 4rem;
    padding: 2rem;
}

.top Button {
    margin-top: 1.5rem;
}

.passwordForm {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    width: 100%;
    margin-top: 6rem;
    padding: 1rem;
}

.pwd {
    width: 500px;
    margin-top: 1.5rem;
    display: flex;
    flex-direction: column;
    align-items: center;
}

.pwd .pwdInput {
    margin-bottom: 1.5rem;
    width: 100%;
    border: none;
}

.save {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}

.save Button {
    margin-bottom: 1rem;
}

.save p {
    color: rgb(155, 0, 0);
    font-weight: bolder;
    text-align: center;
    margin-bottom: 1rem;
}
</style>