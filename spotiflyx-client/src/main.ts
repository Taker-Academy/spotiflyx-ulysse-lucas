import './assets/main.css'
import 'primeicons/primeicons.css'

import { createApp } from 'vue'
import App from './App.vue'
import PrimeVue from 'primevue/config'
import 'primevue/resources/themes/aura-dark-green/theme.css'
import InputText from 'primevue/inputtext'
import InpuGroup from 'primevue/inputgroup'
import InputGroupAddon from 'primevue/inputgroupaddon'
import Password from 'primevue/password'
import Button from 'primevue/button'

const app = createApp(App)

app.use(PrimeVue)
app.component('InputText', InputText)
app.component('Password', Password)
app.component('InputGroup', InpuGroup)
app.component('Button', Button)
app.component('InputGroupAddon', InputGroupAddon)
app.mount('#app')

