<template>
  <form
    class="flex flex-row items-center justify-center flex-no-wrap rounded-full w-11/12 max-w-lg"
    style="box-shadow: 0 1px 6px 0 rgba(32,33,36,0.28);"
    @submit.prevent="submit"
  >
    <div
      class="p-2 flex items-center justify-center flex-wrap rounded-l-full bg-white border-2 border-transparent border-r-0 border-solid flex-grow flex-shrink transition-colors duration-150"
    >
      <div class="w-2"></div>
      <input
        v-model="targetURL"
        type="text"
        class="placeholder-gray-500 focus:outline-none tracking-wide border-0 flex-grow flex-shrink"
        placeholder="https://vuejs.org/v2/guide/forms.html"
      />
    </div>
    <button
      type="submit"
      ref="submitButton"
      class="bg-blue-500 border-2 border-blue-500 border-solid text-white p-2 pr-3 rounded-r-full transition duration-100 ease-linear focus:outline-none uppercase focus:bg-pink-500 focus:border-pink-500 hover:bg-pink-500 hover:border-pink-500 tracking-wide"
    >
      shorten
    </button>
  </form>
</template>

<script>
import axios from "axios";

export default {
  data: () => ({
    targetURL: ""
  }),
  methods: {
    async api() {
      try {
        const url = `${
          process.env.NODE_ENV === "production" ? "/goto" : ""
        }/api/shorten`;
        const { data } = await axios.post(url, {
          targetURL: this.targetURL
        });
        this.$emit("success", data.shortURL);
      } catch (e) {
        this.$emit("failure", e);
      }
    },
    submit() {
      this.$refs.submitButton.focus();
      this.$emit("submit");
      this.api();
    }
  }
};
</script>

<style></style>
