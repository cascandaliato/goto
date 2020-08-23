<template>
  <div
    class="mt-16 bg-white p-4 text-lg flex justify-center items-center border-2 border-red-500"
  >
    <span class="text-red-600">
      <FontAwesomeIcon :icon="['fas', 'times']" size="2x" />
    </span>
    <span class="ml-4">{{ message }}</span>
  </div>
</template>

<script>
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

export default {
  props: ["message"],
  components: { FontAwesomeIcon }
};
</script>

<style></style>
