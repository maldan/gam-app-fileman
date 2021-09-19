import { createApp } from 'vue';
// @ts-ignore
import App from './App.vue';
import Router from './router';
import Store from './store';
import './main.scss';
import { init } from './gam_sdk_ui/vue/event';
//import devtools from '@vue/devtools';

const app = createApp(App);

init(app);

app.config.globalProperties.$store = Store;
app.use(Router).use(Store).mount('#app');

if (process.env.NODE_ENV === 'development') {
  //devtools.connect(/* host, port */);
}
