<template>
  <form
    class="flex flex-row items-center justify-center flex-no-wrap rounded-full w-full transition duration-75 ease-linear"
    @submit.prevent="submit"
  >
    <div
      class="p-2 flex items-center justify-center flex-wrap rounded-l-full bg-white border-2 border-transparent border-r-0 border-solid flex-grow flex-shrink"
    >
      <div class="w-2"></div>
      <input
        aria-label="Target URL"
        v-model="targetURL"
        ref="targetURLInput"
        type="text"
        class="placeholder-gray-500 focus:outline-none tracking-wide border-0 flex-grow flex-shrink"
        placeholder="Enter a URL to shorten..."
      />
    </div>
    <button
      type="submit"
      ref="submitButton"
      class="bg-blue-500 border-2 border-blue-500 border-solid text-white p-2 pr-3 rounded-r-full transition duration-75 ease-linear focus:outline-none uppercase focus:bg-pink-500 focus:border-pink-500 hover:bg-pink-500 hover:border-pink-500 tracking-wide font-semibold"
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
        const { data } = await axios.post("/api/shorten", {
          targetURL: this.targetURL
        });
        if (data.error) {
          this.$emit("failure", new Error(data.message));
        } else {
          this.$emit("success", data.shortURL);
        }
      } catch (e) {
        this.$emit("failure", e);
      }
    },
    submit() {
      this.$refs.submitButton.focus();
      this.$emit("submit");
      this.api();
    }
  },
  mounted() {
    this.$refs.targetURLInput.focus();
  }
};
</script>

<style scoped>
form {
  box-shadow: 0 1px 2px 0 rgba(60, 64, 67, 0.302),
    0 1px 3px 1px rgba(60, 64, 67, 0.149);
}

form:hover,
form:focus-within {
  box-shadow: 0 1px 3px 0 rgba(60, 64, 67, 0.302),
    0 4px 8px 3px rgba(60, 64, 67, 0.149);
}
</style>
