<template>
    <Toast />
</template>

<script setup>
import { useToast } from 'primevue/usetoast';
import Toast from 'primevue/toast';
import { useRouter, useRoute } from 'vue-router';

const toast = useToast();
const route = useRoute();

const connexionPopup = () => {
    toast.add({severity:'success', summary: '⠀Bienvenue', detail: 'Vous voilà connecté !', life: 6000});
    localStorage.removeItem('connected');
};

if (route.path != '/signin' && route.path != '/signup' &&
    localStorage.getItem('connected') == 'true') {
    connexionPopup();
}

const expose = { connexionPopup };
</script>