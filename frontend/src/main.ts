import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import './main.scss';

const app = createApp(App);
app.directive('doubleclick', {
  created(el, binding, vnode, prevVnode) {
    let counter = 0;
    let lastClick = new Date();
    el.addEventListener('click', (e: any) => {
      if (new Date().getTime() - lastClick.getTime() > 800) {
        counter = 0;
        lastClick = new Date();
        return;
      }

      lastClick = new Date();
      counter += 1;
      if (counter > 1) {
        binding.value(e.pageX / window.innerWidth);
        counter = 0;
      }
    });
  },
});

app.use(router).mount('#app');
