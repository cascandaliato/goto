import Vue from "vue";
import App from "./App.vue";
import VTooltip from "v-tooltip";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faClipboard } from "@fortawesome/free-regular-svg-icons";
import { faTimes } from "@fortawesome/free-solid-svg-icons";
import "./assets/styles/index.css";

library.add(faClipboard);
library.add(faTimes);

Vue.use(VTooltip);
Vue.config.productionTip = false;

new Vue({
  render: h => h(App)
}).$mount("#app");
