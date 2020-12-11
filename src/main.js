import { library } from "@fortawesome/fontawesome-svg-core";
import { faGithub, faTwitter } from "@fortawesome/free-brands-svg-icons";
import { faClipboard } from "@fortawesome/free-regular-svg-icons";
import { faTimes } from "@fortawesome/free-solid-svg-icons";
import VTooltip from "v-tooltip";
import Vue from "vue";
import App from "./App.vue";
import "./assets/styles/index.css";

library.add(faClipboard);
library.add(faTimes);
library.add(faGithub);
library.add(faTwitter);

Vue.use(VTooltip);
Vue.config.productionTip = false;

new Vue({
  render: h => h(App)
}).$mount("#app");
