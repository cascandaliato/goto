<template>
  <div
    class="h-screen font-sans antialiased text-gray-800 flex flex-col items-center justify-start"
  >
    <Logo class="w-full h-auto mt-24 sm:mt-32" />
    <main
      class="flex flex-col items-center justify-start mt-8 w-11/12 max-w-lg"
    >
      <InputForm
        @submit="handleSubmit"
        @success="handleSuccess"
        @failure="handleFailure"
      />
      <Spinner class="mt-20" v-show="showSpinner" />
      <transition
        name="comp-custom-classes-transition"
        enter-active-class="animate__animated animate__fadeIn animate__faster"
        leave-active-class="animate__animated animate__fadeOut animate__faster"
        @after-leave="afterLeave"
      >
        <ShortURL class="mt-16" :shortURL="shortURL" v-show="showShortURL" />
      </transition>
    </main>
    <Footer
      class="text-xs text-gray-500 flex-grow flex flex-col justify-end items-center pb-1 mt-12"
    />
  </div>
</template>

<script>
import Spinner from "./components/Spinner.vue";
import Footer from "./components/Footer.vue";
import ShortURL from "./components/ShortURL.vue";
import InputForm from "./components/InputForm.vue";
import Logo from "./components/Logo.vue";

export default {
  name: "App",
  components: { Footer, InputForm, Spinner, ShortURL, Logo },
  data: () => ({
    shortURL: null,
    showSpinner: false,
    showShortURL: false
  }),
  methods: {
    handleSubmit() {
      this.showShortURL = false;
      if (!this.shortURL) this.showSpinner = true;
    },
    handleSuccess(shortURL) {
      this.showSpinner = false;
      this.showShortURL = true;
      this.shortURL = shortURL;
    },
    handleFailure() {
      this.showSpinner = false;
      this.showShortURL = true;
    },
    afterLeave() {
      this.showSpinner = true;
    }
  }
};
</script>

<style>
body {
  background-color: #f7fafc;
}
</style>
