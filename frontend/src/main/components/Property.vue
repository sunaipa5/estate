<template>
  <div>
    <div v-for="property in properties" :key="property.id">
      <div class="col col-xxl imgCol">
        <Splide :options="options">
          <SplideSlide v-for="(image, index) in images" :key="index">
            <img :src="'/img/' + image" class="center" />
          </SplideSlide>
        </Splide>
      </div>
      <br />
      <div class="col col-xxl">
        <h2>{{ property.Title }}</h2>
        <p>{{ property.Centare }} m2 | {{ property.Room_number }}</p>
        <p>{{ property.Address }}</p>
      </div>
      <br />
      <div class="col col-xxl">
        <pre>
        {{property.Description}}
      </pre>
      </div>
    </div>
  </div>
</template>

<script>
import "@splidejs/vue-splide/css";
import axios from "axios";
import { Splide, SplideSlide } from "@splidejs/vue-splide";

export default {
  props: ["pageName"],
  data() {
    return {
      properties: [],
      images: [],
      currentIndex: 0,
      translateX: 0,
      intervalId: null,
    };
  },
  created() {
    this.getPropertyDetails();
  },
  setup() {
    const options = {
      rewind: true,
      type: "loop",
      autoplay: true,
      padding: "8rem",
      gap: "40px",
    };

    return { options };
  },
  methods: {
    getPropertyDetails() {
      const postData = {
        PageName: this.pageName,
      };
      axios
        .post("/api/getProperty", postData)
        .then((response) => {
          if (response != undefined) {
            console.log(response.data);
            this.properties = response.data;
            document.title = this.properties[0].Title;
            this.images = this.properties[0].File_names.trim().split(" ");
          }
        })
        .catch((error) => {
          console.error("Get Error:", error);
        });
    },
    components: {
      Splide,
      SplideSlide,
    },
  },
};

</script>

<style>
.splide__slide {
  width: 100%;
  height: 50vh;
  border-radius: 7px;
  background: #ececec;
}

.dark .splide__slide {
  background-color: #111;
}

.splide__slide img {
  max-height: 100%;
}

pre {
    white-space: pre-wrap;
    margin-left: 0;
}
</style>