<template>
  <footer class="text-center">
    <span>
      Logo font
      <a
        class="hover:underline text-gray-900"
        href="https://www.fontspring.com/fonts/los-andes-type/fibra/fibra-light"
        title="fibra-light"
        >fibra-light</a
      >
      by
      <a
        class="hover:underline text-gray-900"
        href="https://losandestype.myportfolio.com/"
        title="Los Andes Type"
        >Los Andes Type</a
      >
    </span>
    <span>
      Icons by
      <a
        class="hover:underline text-gray-900"
        href="https://www.flaticon.com/authors/freepik"
        title="Freepik"
        >Freepik</a
      >
      and
      <a
        class="hover:underline text-gray-900"
        href="https://www.flaticon.com/authors/pixel-perfect"
        title="Pixel perfect"
        >Pixel perfect</a
      >
      from
      <a
        class="hover:underline text-gray-900"
        href="https://www.flaticon.com/"
        title="Flaticon"
        >flaticon.com</a
      >
    </span>
    <span>
      and Dave Gandy from
      <a
        class="hover:underline text-gray-900"
        href="http://fontawesome.io"
        title="Font Awesome"
        >fontawesome.io</a
      >
    </span>
    <span>
      © 2020
      <a
        href="https://cascandaliato.com"
        title="Author's home page"
        class="hover:underline text-gray-900"
        >Carmelo Scandaliato</a
      >&nbsp;
      <a
        href="https://github.com/cascandaliato/goto"
        title="Project's GitHub repository"
        class="hover:underline text-gray-900"
        ><FontAwesomeIcon :icon="['fab', 'github']" size="lg"/></a
      >&nbsp;
      <a
        href="https://twitter.com/cascandaliato"
        title="Author's Twitter profile"
        class="hover:underline text-gray-900"
        ><FontAwesomeIcon :icon="['fab', 'twitter']" size="lg"
      /></a>
    </span>
  </footer>
</template>

<script>
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

export default {
  components: { FontAwesomeIcon }
};
</script>

<style></style>
