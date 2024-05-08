<template>
    <main>
        <div class="form">
            <h1 class="title" style="font-weight: bold;">Nouveau compte:</h1>
            <h3>Adresse Email</h3>
            <InputGroup class="email">
                <InputGroupAddon>
                    <i class="pi pi-envelope"></i>
                </InputGroupAddon>
                <InputText class="input" v-model="email" placeholder="Email" />
            </InputGroup>
            <h3>Mot de passe</h3>
            <InputGroup class="password">
                <InputGroupAddon>
                    <i class="pi pi-lock"></i>
                </InputGroupAddon>
                <Password class="input" v-model="password" placeholder="Mot de passe" promptLabel="Choisie un mot de passe" weakLabel="Trop simple" mediumLabel="Normal" strongLabel="Complexe" toggleMask/>
            </InputGroup>
            <Button class="btn" label="Nouveau compte" @click="signUpFunc" :loading="visible"></Button>
            <div class="error"><h3 class="errorTxt"></h3></div>
            <div class="signin">
                <p>Déja un compte ?</p>
                <router-link :to="{ path: '/signin', query: { redirect: redirect } }">Se connecter</router-link>
            </div>
        </div>
    </main>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { ax } from '../router/router'

const route = useRoute();
const redirect = route.query.redirect;

const email = ref('');
const password = ref('');
const router = useRouter();
const visible = ref(false);

const signUpFunc = () => {
    visible.value = true;
    ax.post('/auth/signup', {
        email: email.value,
        password: password.value
    }).then(async (res: any) => {
        // delete token from local storage
        localStorage.removeItem('token');
        localStorage.setItem('token', res.data.data.token);
        ax.defaults.headers.common['Authorization'] = 'Bearer ' + res.data.data.token;
        console.log("Signed up")
        if (redirect && redirect != '/signin' && redirect != '/signup')
            router.push(redirect);
        else
            router.push('/home');
    }).catch((err: any) => {
        console.log(err);
        visible.value = false;
        const errorTxt = document.querySelector('.errorTxt');
        if (errorTxt) {
            errorTxt.textContent = 'Invalide email ou mot de passe';
        }
        password.value = '';
    });
}
</script>

<style scoped>
main {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100vh;
    width: 100%;
    background-color: var(--primary-color);
}

.title {
    margin-bottom: 1rem;
}

.input {
    height: 60px;
}

.form {
    padding: 5rem;
    border-radius: 10px;
    height: min-content;
    width: min-content;
    display: flex;
    flex-direction: column;
    align-items: start;
    background-color: var(--color-background-soft);
}

.email {
    width: 400px;
    margin-bottom: 3rem;
}

.signin {
    margin-top: 1rem;
    display: flex;
    align-items: center;
    justify-content: center;
}

.signin a {
    margin-left: 0.5rem;
    color: var(--color-primary);
    font-weight: 700;
}

.password {
    width: 400px;
    margin-bottom: 3rem;
}

.error {
    display: flex;
    align-items: center;
    justify-content: center;
    margin-top: 0.8rem;
    width: 400px;
}

.errorTxt {
    color: rgb(155, 0, 0);
    font-weight: bolder;
    text-align: center;
}

.btn {
    width: 400px;
    height: 60px;
    font-size: 20px;
    font-weight: 700;
    margin-top: 1rem;
    text-align: center;
    display: flex;
    align-items: center;
    justify-content: center;
}
</style>
