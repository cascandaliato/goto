<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png" />
    <hr />
    <button @click="api">call api</button>
    <br />
    {{ data }}
    <hr />
    <HelloWorld msg="Welcome to Your Vue.js App" />
  </div>
</template>

<script>
import HelloWorld from "./components/HelloWorld.vue";
import axios from "axios";

export default {
  name: "App",
  components: {
    HelloWorld
  },
  data: () => ({ data: {} }),
  methods: {
    async api() {
      try {
        const { data } = await axios.post("/api/shorten", {
          targetURL: `https://${Math.round(Math.random() * 100000)}.com/`
        });
        this.data = data;
      } catch (e) {
        console.log(e);
      }
    }
  },
  created() {
    // setInterval(this.api, 5000);
  }
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
