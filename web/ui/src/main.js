import "@babel/polyfill";
//import "./assets/safari.scss"
import Vue from "vue";
import router from './router'
import 'bootstrap'
import App from "./App";


import VueResource from 'vue-resource';
Vue.use(VueResource);

import VCalendar from 'v-calendar'
Vue.use(VCalendar);

Vue.config.productionTip = false;

// Attempt at some error handling
Vue.config.errorHandler = function (err, vm, info) {
//  Vue.notify({ type: "error", text: err });
};

window.addEventListener("unhandledrejection", function (err, promise) {
//  Vue.notify({ type: "error", text: err });
});

window.onerror = function (msg, url, line, col, error) {
//  Vue.notify({ type: "error", text: msg });
};

// Start timer. Provides a reactive timestamp, updated each second
//store.dispatch("time/start");

    // Launch vue app!
    window.Vue = new Vue({
      el: "#app",
      render: h => h(App),
      router,
    });

