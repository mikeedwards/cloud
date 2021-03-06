import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import i18n from "./i18n";
import VCalendar from "v-calendar";

Vue.use(VCalendar, {});

Vue.config.productionTip = false;

new Vue({
    i18n,
    router,
    render: h => h(App),
}).$mount("#app");
