<template>
  <div
    class="bg-white p-4 text-lg flex justify-center items-center border-2 border-gray-400 border-dashed relative"
  >
    <Scissors
      @click="selectText()"
      class="inline-block h-5 w-5 absolute left-2 -top-07 cursor-pointer"
    />

    <div id="shortURL" ref="shortURL" @click="selectText()">{{ shortURL }}</div>
    <div class="flex flex-col justify-center items-center">
      <span
        class="text-blue-500 hover:text-pink-500 cursor-pointer ml-4 transition duration-75 ease-linear"
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
        <FontAwesomeIcon :icon="['far', 'clipboard']" size="2x" />
      </span>
    </div>
  </div>
</template>

<script>
import copy from "copy-to-clipboard";
import Scissors from "./Scissors.vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

export default {
  props: ["shortURL"],
  components: { Scissors, FontAwesomeIcon },
  data: () => ({
    tooltipMessage: "Click to copy"
  }),
  methods: {
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
.tooltip {
  display: block !important;
  z-index: 10000;
}

.tooltip .tooltip-inner {
  background: #3b82f6;
  color: white;
  font-weight: 500;
  border-radius: 6px;
  padding: 5px 10px 4px;
}

.tooltip .tooltip-arrow {
  width: 0;
  height: 0;
  border-style: solid;
  position: absolute;
  margin: 5px;
  border-color: #3b82f6;
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
