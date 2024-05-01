<template>
    <main>
        <div class="form">
            <h1 class="title" style="font-weight: bold;">Sign Up:</h1>
            <h3>Email address</h3>
            <InputGroup class="email">
                <InputGroupAddon>
                    <i class="pi pi-envelope"></i>
                </InputGroupAddon>
                <InputText class="input" v-model="email" placeholder="Enter email" />
            </InputGroup>
            <h3>Set Password</h3>
            <InputGroup class="password">
                <InputGroupAddon>
                    <i class="pi pi-lock"></i>
                </InputGroupAddon>
                <Password class="input" v-model="password" placeholder="Enter password" toggleMask/>
            </InputGroup>
            <Button class="btn" @click="signUpFunc">Sign Up</Button>
            <div class="signin">
                <p>Already have an account?</p>
                <router-link to="/signin">Sign In</router-link>
            </div>
        </div>
    </main>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router'; // Import useRouter for programmatic navigation

const email = ref('');
const password = ref('');
const router = useRouter(); // Create a router instance

const signUpFunc = () => {
    console.log(email.value, password.value);
    // Here you can add logic to handle signup, e.g., send data to backend

    // After successful signup, navigate to another route
    router.push('/');
}
</script>

<style scoped>
main {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100vh;
    width: 100vw;
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
