<!-- CustomAlert.vue -->
<template>
  <div>
    <div class="calbg"></div>
    <div class="calert col col-sm">
      <div class="btna" style="width:100%;">
        <svg
          width="18"
          height="18"
          class="flr"
          color="red"
          @click="closeAlert"
          style="position:absolute;top:3px;right:4px;"
        >
          <use :xlink:href="svgLocation + '#close'" />
        </svg>
      </div>
      <div>
         <slot></slot>
      </div>
    </div>
  </div>
</template>

<script>
import svgLoc from "../../assets/all.svg";
export default {
    data() {
    return {
      svgLocation: svgLoc,
    };
  },
  methods: {
    closeAlert() {
      this.$emit("close");
    },
  },
};
</script>

<style scoped>
.calert {
  -webkit-tap-highlight-color: transparent;
  -webkit-touch-callout: none;
  -webkit-user-select: none;
  -khtml-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
  -moz-user-select: none;
  -webkit-user-select: none;
  -ms-user-select: none;
  user-select: none;
  -webkit-user-drag: none;
  -webkit-touch-callout: none;
}

.calert {
  margin: auto;
  width: 300px;
  padding: 10px;
  z-index: 8;
  transform: translate(-50%, 10%);
  position: fixed;
  overflow: hidden;
  z-index: 1032;
  top: 20%;
  left: 50%;
  transform: translate(-50%, -20%);
 
}
.dark .calert{
  background-color: #090909;
}
.falert {
  width: 500px;
  max-width: 90vw;
}

.falert iframe {
  width: 100%;
  height: 500px;
  border: none;
  border-radius: 7px;
}

.calbg {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 1031;
  background-color: rgba(0, 0, 0, 0.5);
}

#closePopup {
  transition: 0.5s;
  color: #911;
}

#closePopup:hover {
  color: #711;
}
</style>
