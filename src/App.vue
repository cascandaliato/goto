<template>
  <div
    class="font-sans antialiased text-gray-800 flex flex-col items-center justify-start mt-16"
  >
    <img
      class="block w-3/4 max-w-md"
      src="https://via.placeholder.com/500x250"
      alt="app logo"
    />
    <form
      class="flex flex-col items-center justify-center w-full mt-16"
      @submit.prevent="submit"
    >
      <div
        class="inline-flex items-center justify-center rounded-full w-11/12 max-w-lg flex-wrap"
        style="box-shadow: 0 1px 6px 0 rgba(32,33,36,0.28);"
      >
        <div
          class="p-2 flex-wrap flex items-center justify-center rounded-l-full bg-white border-2 border-transparent border-r-0 border-solid flex-grow transition-colors duration-150"
        >
          <div class="w-2"></div>
          <input
            type="text"
            class="placeholder-gray-500 focus:outline-none tracking-wide border-0 flex-grow flex-shrink"
            placeholder="https://vuejs.org/v2/guide/forms.html"
            v-model="targetURL"
          />
        </div>
        <button
          ref="submitButton"
          class="bg-blue-500 border-2 border-blue-500 border-solid text-white p-2 pr-3 rounded-r-full transition duration-150 ease-linear focus:outline-none uppercase focus:bg-pink-500 focus:border-pink-500 tracking-wide"
        >
          shorten
        </button>
      </div>
    </form>
    <div class="loader mt-32" v-show="awaiting">Loading ...</div>
    <div
      class="mt-24 bg-white p-4 text-lg flex justify-center items-center border-2 border-gray-400 border-dashed"
      v-show="!awaiting && shortURL"
    >
      <div ref="shortURL" @click="selectText()">{{ shortURL }}</div>
      <div class="flex flex-col justify-center items-center">
        <span
          class="text-blue-500 cursor-pointer ml-4"
          @click="copyToClipboard"
          v-tooltip="{
            content: this.tooltipMessage,
            autoHide: true,
            hideOnTargetClick: false,
            offset: 2,
            delay: 50
          }"
          @mouseleave="resetTooltipMessage"
        >
          <i class="far fa-clipboard fa-2x"></i>
        </span>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import copy from "copy-to-clipboard";

export default {
  name: "App",
  components: {},
  data: () => ({
    shortURL: "",
    targetURL: "",
    focus: false,
    awaiting: false,
    tooltipMessage: "Copy to clipboard"
  }),
  methods: {
    async api() {
      try {
        this.awaiting = true;
        const url = `${
          process.env.NODE_ENV === "production" ? "/goto" : ""
        }/api/shorten`;
        const { data } = await axios.post(url, {
          targetURL: this.targetURL
        });
        this.shortURL = data.shortURL;
        this.awaiting = false;
      } catch (e) {
        this.awaiting = false;
      }
    },
    submit() {
      this.$refs.submitButton.focus();
      this.api();
    },
    selectShortURL() {
      this.$refs.shortURL.textselect();
    },
    selectText() {
      var sel, range;
      // var el = document.getElementById(id); //get element id
      var el = this.$refs.shortURL;
      if (window.getSelection && document.createRange) {
        //Browser compatibility
        sel = window.getSelection();
        if (sel.toString() == "") {
          //no text selection
          window.setTimeout(function() {
            range = document.createRange(); //range object
            range.selectNodeContents(el); //sets Range
            sel.removeAllRanges(); //remove all ranges from selection
            sel.addRange(range); //add Range to a Selection.
          }, 1);
        }
      } else if (document.selection) {
        //older ie
        sel = document.selection.createRange();
        if (sel.text == "") {
          //no text selection
          range = document.body.createTextRange(); //Creates TextRange object
          range.moveToElementText(el); //sets Range
          range.select(); //make selection.
        }
      }
    },
    copyToClipboard() {
      copy(this.shortURL);
      this.tooltipMessage = "Copied!";
    },
    resetTooltipMessage() {
      setTimeout(() => {
        this.tooltipMessage = "Copy to clipboard";
      }, 200);
    }
  }
};
</script>

<style>
.loader {
  color: #ed64a6;
  font-size: 20px;
  margin: 100px auto;
  width: 1em;
  height: 1em;
  border-radius: 50%;
  position: relative;
  text-indent: -9999em;
  -webkit-animation: load4 1.3s infinite linear;
  animation: load4 1.3s infinite linear;
  -webkit-transform: translateZ(0);
  -ms-transform: translateZ(0);
  transform: translateZ(0) scale(0.4);
}
@-webkit-keyframes load4 {
  0%,
  100% {
    box-shadow: 0 -3em 0 0.2em, 2em -2em 0 0em, 3em 0 0 -1em, 2em 2em 0 -1em,
      0 3em 0 -1em, -2em 2em 0 -1em, -3em 0 0 -1em, -2em -2em 0 0;
  }
  12.5% {
    box-shadow: 0 -3em 0 0, 2em -2em 0 0.2em, 3em 0 0 0, 2em 2em 0 -1em,
      0 3em 0 -1em, -2em 2em 0 -1em, -3em 0 0 -1em, -2em -2em 0 -1em;
  }
  25% {
    box-shadow: 0 -3em 0 -0.5em, 2em -2em 0 0, 3em 0 0 0.2em, 2em 2em 0 0,
      0 3em 0 -1em, -2em 2em 0 -1em, -3em 0 0 -1em, -2em -2em 0 -1em;
  }
  37.5% {
    box-shadow: 0 -3em 0 -1em, 2em -2em 0 -1em, 3em 0em 0 0, 2em 2em 0 0.2em,
      0 3em 0 0em, -2em 2em 0 -1em, -3em 0em 0 -1em, -2em -2em 0 -1em;
  }
  50% {
    box-shadow: 0 -3em 0 -1em, 2em -2em 0 -1em, 3em 0 0 -1em, 2em 2em 0 0em,
      0 3em 0 0.2em, -2em 2em 0 0, -3em 0em 0 -1em, -2em -2em 0 -1em;
  }
  62.5% {
    box-shadow: 0 -3em 0 -1em, 2em -2em 0 -1em, 3em 0 0 -1em, 2em 2em 0 -1em,
      0 3em 0 0, -2em 2em 0 0.2em, -3em 0 0 0, -2em -2em 0 -1em;
  }
  75% {
    box-shadow: 0em -3em 0 -1em, 2em -2em 0 -1em, 3em 0em 0 -1em, 2em 2em 0 -1em,
      0 3em 0 -1em, -2em 2em 0 0, -3em 0em 0 0.2em, -2em -2em 0 0;
  }
  87.5% {
    box-shadow: 0em -3em 0 0, 2em -2em 0 -1em, 3em 0 0 -1em, 2em 2em 0 -1em,
      0 3em 0 -1em, -2em 2em 0 0, -3em 0em 0 0, -2em -2em 0 0.2em;
  }
}
@keyframes load4 {
  0%,
  100% {
    box-shadow: 0 -3em 0 0.2em, 2em -2em 0 0em, 3em 0 0 -1em, 2em 2em 0 -1em,
      0 3em 0 -1em, -2em 2em 0 -1em, -3em 0 0 -1em, -2em -2em 0 0;
  }
  12.5% {
    box-shadow: 0 -3em 0 0, 2em -2em 0 0.2em, 3em 0 0 0, 2em 2em 0 -1em,
      0 3em 0 -1em, -2em 2em 0 -1em, -3em 0 0 -1em, -2em -2em 0 -1em;
  }
  25% {
    box-shadow: 0 -3em 0 -0.5em, 2em -2em 0 0, 3em 0 0 0.2em, 2em 2em 0 0,
      0 3em 0 -1em, -2em 2em 0 -1em, -3em 0 0 -1em, -2em -2em 0 -1em;
  }
  37.5% {
    box-shadow: 0 -3em 0 -1em, 2em -2em 0 -1em, 3em 0em 0 0, 2em 2em 0 0.2em,
      0 3em 0 0em, -2em 2em 0 -1em, -3em 0em 0 -1em, -2em -2em 0 -1em;
  }
  50% {
    box-shadow: 0 -3em 0 -1em, 2em -2em 0 -1em, 3em 0 0 -1em, 2em 2em 0 0em,
      0 3em 0 0.2em, -2em 2em 0 0, -3em 0em 0 -1em, -2em -2em 0 -1em;
  }
  62.5% {
    box-shadow: 0 -3em 0 -1em, 2em -2em 0 -1em, 3em 0 0 -1em, 2em 2em 0 -1em,
      0 3em 0 0, -2em 2em 0 0.2em, -3em 0 0 0, -2em -2em 0 -1em;
  }
  75% {
    box-shadow: 0em -3em 0 -1em, 2em -2em 0 -1em, 3em 0em 0 -1em, 2em 2em 0 -1em,
      0 3em 0 -1em, -2em 2em 0 0, -3em 0em 0 0.2em, -2em -2em 0 0;
  }
  87.5% {
    box-shadow: 0em -3em 0 0, 2em -2em 0 -1em, 3em 0 0 -1em, 2em 2em 0 -1em,
      0 3em 0 -1em, -2em 2em 0 0, -3em 0em 0 0, -2em -2em 0 0.2em;
  }
}

.tooltip {
  display: block !important;
  z-index: 10000;
}

.tooltip .tooltip-inner {
  background: #4299e1;
  color: white;
  border-radius: 6px;
  padding: 5px 10px 4px;
}

.tooltip .tooltip-arrow {
  width: 0;
  height: 0;
  border-style: solid;
  position: absolute;
  margin: 5px;
  border-color: #4299e1;
  z-index: 1;
}

.tooltip[x-placement^="top"] {
  margin-bottom: 5px;
}

.tooltip[x-placement^="top"] .tooltip-arrow {
  border-width: 5px 5px 0 5px;
  border-left-color: transparent !important;
  border-right-color: transparent !important;
  border-bottom-color: transparent !important;
  bottom: -5px;
  left: calc(50% - 5px);
  margin-top: 0;
  margin-bottom: 0;
}

.tooltip[x-placement^="bottom"] {
  margin-top: 5px;
}

.tooltip[x-placement^="bottom"] .tooltip-arrow {
  border-width: 0 5px 5px 5px;
  border-left-color: transparent !important;
  border-right-color: transparent !important;
  border-top-color: transparent !important;
  top: -5px;
  left: calc(50% - 5px);
  margin-top: 0;
  margin-bottom: 0;
}

.tooltip[x-placement^="right"] {
  margin-left: 5px;
}

.tooltip[x-placement^="right"] .tooltip-arrow {
  border-width: 5px 5px 5px 0;
  border-left-color: transparent !important;
  border-top-color: transparent !important;
  border-bottom-color: transparent !important;
  left: -5px;
  top: calc(50% - 5px);
  margin-left: 0;
  margin-right: 0;
}

.tooltip[x-placement^="left"] {
  margin-right: 5px;
}

.tooltip[x-placement^="left"] .tooltip-arrow {
  border-width: 5px 0 5px 5px;
  border-top-color: transparent !important;
  border-right-color: transparent !important;
  border-bottom-color: transparent !important;
  right: -5px;
  top: calc(50% - 5px);
  margin-left: 0;
  margin-right: 0;
}

.tooltip.popover .popover-inner {
  background: #f9f9f9;
  color: black;
  padding: 24px;
  border-radius: 5px;
  box-shadow: 0 5px 30px rgba(black, 0.1);
}

.tooltip.popover .popover-arrow {
  border-color: #f9f9f9;
}

.tooltip[aria-hidden="true"] {
  visibility: hidden;
  opacity: 0;
  transition: opacity 0.15s, visibility 0.15s;
}

.tooltip[aria-hidden="false"] {
  visibility: visible;
  opacity: 1;
  transition: opacity 0.15s;
}
</style>
