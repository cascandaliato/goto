<template>
  <div>
    <div
      class="bg-gray-200 w-screen h-56 font-sans antialiased flex items-center justify-center text-gray-800 text-xl"
    >
      <div class="shadow-lg flex items-center justify-center max-w-3xl w-4/5 rounded-l-md">
        <input
          type="text"
          class="placeholder-gray-500 p-2 rounded-l-md w-full focus:outline-none border-2 border-r-0 border-solid border-transparent focus:border-blue-500 tracking-wide"
          placeholder="https://vuejs.org/v2/guide/forms.html"
          v-model="targetURL"
        />
        <button
          class="bg-blue-500 text-white p-2 rounded-r-md focus:outline-none lowercase border-2 border-solid border-transparent focus:bg-green-500 tracking-wide"
          @click="api"
        >shorten</button>
      </div>
    </div>
    {{ targetURL }} -> {{ shortURL }}
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "App",
  components: {},
  data: () => ({ shortURL: "", targetURL: "" }),
  methods: {
    async api() {
      try {
        const url = `${
          process.env.NODE_ENV === "production" ? "/goto" : ""
        }/api/shorten`;
        const { data } = await axios.post(url, {
          targetURL: this.targetURL
        });
        this.shortURL = data.shortURL;
      } catch (e) {
        console.log(e);
      }
    }
  }
};
</script>

<style>
</style>
