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
import routers from './router/router'

const appli = createApp(App)

appli.use(PrimeVue)
appli.use(routers)
appli.component('InputText', InputText)
appli.component('Password', Password)
appli.component('InputGroup', InpuGroup)
appli.component('Button', Button)
appli.component('InputGroupAddon', InputGroupAddon)
appli.mount('#app')

