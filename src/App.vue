<template>
  <div
    class="h-screen font-sans antialiased text-gray-800 flex flex-col  items-center justify-start"
  >
    <Logo class="w-11/12 max-w-lg mt-24 sm:mt-32" />
    <main
      class="flex flex-col items-center justify-start mt-8 sm:mt-10 w-11/12 max-w-lg"
    >
      <InputForm
        @submit="handleSubmit"
        @success="handleSuccess"
        @failure="handleFailure"
      />
      <Spinner class="mt-20" v-show="showSpinner" />
      <transition
        name="short-url-transition"
        enter-active-class="animate__animated animate__fadeIn animate__faster"
        leave-active-class="animate__animated animate__fadeOut animate__faster"
        @after-leave="afterLeave"
      >
        <ShortURL class="mt-16" :shortURL="shortURL" v-show="showShortURL" />
      </transition>
      <transition
        name="error-transition"
        enter-active-class="animate__animated animate__fadeIn animate__faster"
        leave-active-class="animate__animated animate__fadeOut animate__faster"
        @after-leave="afterLeave"
      >
        <Error class="mt-16" :message="errorMessage" v-show="showError" />
      </transition>
    </main>
    <Footer
      class="text-xs text-gray-700 flex-grow flex flex-col justify-end items-center pb-1 mt-12"
    />
  </div>
</template>

<script>
import Spinner from "./components/Spinner.vue";
import Error from "./components/Error.vue";
import Footer from "./components/Footer.vue";
import ShortURL from "./components/ShortURL.vue";
import InputForm from "./components/InputForm.vue";
import Logo from "./components/Logo.vue";

const bgStyles = ["bg-gradient-to-r", "from-white", "via-gray-100", "to-white"];

export default {
  name: "App",
  components: { Error, Footer, InputForm, Logo, ShortURL, Spinner },
  beforeCreate() {
    document.querySelector("body").classList.add(...bgStyles);
  },
  beforeDestroy() {
    document.querySelector("body").classList.remove(...bgStyles);
  },
  data: () => ({
    showShortURL: false,
    shortURL: null,
    showError: false,
    errorMessage: null,
    showSpinner: false,
    firstRequest: true
  }),
  methods: {
    handleSubmit() {
      this.showShortURL = false;
      this.showError = false;
      this.showSpinner = false;

      if (this.firstRequest) this.showSpinner = true;
      this.firstRequest = false;
    },
    handleSuccess(shortURL) {
      this.showSpinner = false;

      this.showShortURL = true;
      this.shortURL = shortURL;
    },
    handleFailure(e) {
      this.showSpinner = false;

      this.showError = true;
      this.errorMessage =
        e.message.charAt(0).toUpperCase() + e.message.slice(1);
    },
    afterLeave() {
      this.shortURL = null;
      this.errorMessage = null;
      this.showSpinner = true;
    }
  }
};
</script>

<style scoped></style>
