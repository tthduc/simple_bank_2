import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import PrimeVue from 'primevue/config';
import Aura from '@primevue/themes/aura';

// Create the app
const app = createApp(App)

// Mount the router to the app
app.use(router)
app.use(PrimeVue, {
    theme: {
        preset: Aura
    }
});

// Mount the app to the #app element
app.mount('#app')
