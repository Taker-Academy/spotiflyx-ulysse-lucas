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
import ProgressSpinner from 'primevue/progressspinner';
import Splitter from 'primevue/splitter';
import SplitterPanel from 'primevue/splitterpanel';
import Dialog from 'primevue/dialog';
import routers from './router/router'

const appli = createApp(App)

appli.use(PrimeVue)
appli.use(routers)
appli.component('Dialog', Dialog)
appli.component('Splitter', Splitter)
appli.component('SplitterPanel', SplitterPanel)
appli.component('ProgressSpinner', ProgressSpinner)
appli.component('InputText', InputText)
appli.component('Password', Password)
appli.component('InputGroup', InpuGroup)
appli.component('Button', Button)
appli.component('InputGroupAddon', InputGroupAddon)
appli.mount('#app')

