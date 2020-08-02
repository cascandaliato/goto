import Vue from "vue";
import App from "./App.vue";
import VTooltip from "v-tooltip";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faClipboard } from "@fortawesome/free-regular-svg-icons";
import "./assets/styles/index.css";

library.add(faClipboard);

Vue.use(VTooltip);
Vue.config.productionTip = false;

new Vue({
  render: h => h(App)
}).$mount("#app");
