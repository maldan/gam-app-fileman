import { createApp } from 'vue';
// @ts-ignore
import App from './App.vue';
import Router from './router';
import Store from './store';
import './main.scss';
import Event from './gam_sdk_ui/vue/event';
//import devtools from '@vue/devtools';
import UI from './gam_sdk_ui/vue/ui';

const app = createApp(App);

app.config.globalProperties.$store = Store;
app.use(Router).use(Store).use(Event).use(UI).mount('#app');

if (process.env.NODE_ENV === 'development') {
  //devtools.connect(/* host, port */);
}
