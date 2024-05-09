import './assets/main.css'
import 'primeicons/primeicons.css'

import { createApp } from 'vue'
import App from './App.vue'
import PrimeVue from 'primevue/config'
import 'primevue/resources/themes/aura-dark-green/theme.css'
import InputText from 'primevue/inputtext'
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';
import InpuGroup from 'primevue/inputgroup'
import InputGroupAddon from 'primevue/inputgroupaddon'
import Password from 'primevue/password'
import Button from 'primevue/button'
import ProgressSpinner from 'primevue/progressspinner';
import Splitter from 'primevue/splitter';
import SplitterPanel from 'primevue/splitterpanel';
import Dialog from 'primevue/dialog';
import ToastService from 'primevue/toastservice';
import Toast from 'primevue/toast';
import Fieldset from 'primevue/fieldset';
import ConfirmDialog from 'primevue/confirmdialog';
import Panel from 'primevue/panel';
import RadioButton from 'primevue/radiobutton';
import ConfirmationService from 'primevue/confirmationservice';
import TabView from 'primevue/tabview';
import TabPanel from 'primevue/tabpanel';
import Tooltip from 'primevue/tooltip';

import routers from './router/router'

const appli = createApp(App)

appli.use(PrimeVue)
appli.use(routers)
appli.use(ToastService)
appli.use(ConfirmationService)
appli.directive('tooltip', Tooltip);
appli.component('IconField', IconField)
appli.component('RadioButton', RadioButton)
appli.component('TabView', TabView)
appli.component('TabPanel', TabPanel)
appli.component('Panel', Panel)
appli.component('ConfirmDialog', ConfirmDialog)
appli.component('Fieldset', Fieldset)
appli.component('InputIcon', InputIcon)
appli.component('Dialog', Dialog)
appli.component('Toast', Toast)
appli.component('Splitter', Splitter)
appli.component('SplitterPanel', SplitterPanel)
appli.component('ProgressSpinner', ProgressSpinner)
appli.component('InputText', InputText)
appli.component('Password', Password)
appli.component('InputGroup', InpuGroup)
appli.component('Button', Button)
appli.component('InputGroupAddon', InputGroupAddon)
appli.mount('#app')

