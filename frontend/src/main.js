import 'core-js/stable';
import 'regenerator-runtime/runtime';
import Vue from 'vue';
import App from './App.vue';

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';
import './assets/tailwind.css'
import router from './router'

import Donut from 'vue-css-donut-chart';
import 'vue-css-donut-chart/dist/vcdonut.css';


import {

    TDropdown,
   
  } from 'vue-tailwind/dist/components';

import VueTailwind from 'vue-tailwind'

const components={
    TDropdown
  }
Vue.use(VueTailwind,components)

Vue.use(Donut);

Wails.Init(() => {
	new Vue({
        router,
        render: h => h(App)
    }).$mount('#app');
});
