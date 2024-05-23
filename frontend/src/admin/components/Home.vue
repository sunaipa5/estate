<template>
  <div>
    <div class="center">
      <div class="col tx-cen">
        <h3>İstanbul</h3>
        <h1>{{timeTurkey}}</h1>
      </div>
      <div class="col tx-cen">
        <h3>Tokyo</h3>
        <h1>{{timeTokyo}}</h1>
      </div>
      <div class="col tx-cen">
        <h3>London</h3>
        <h1>{{timeLondon}}</h1>
      </div>
      <div class="col tx-cen">
        <h3>Florida</h3>
        <h1>{{timeFlorida}}</h1>
      </div>
    </div>
    <div>
      <h2>En son eklenenler</h2>
      <div class="container">
        <div
          class="card ca-md hover"
          v-for="property in properties"
          :key="property.id"
          @click="this.$router.push('/admin/ilan/'+property.Page_name);"
        >
          <div class="inner">
            <img :src="'/img/' + property.File_names.split(' ')[0]" />
            <div class="card-col">
              <p>{{ property.Title }}</p>
              <p>{{ property.Centare }} m2 | {{ property.Room_number }}</p>
              <p>{{ property.Address }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>


<script>
import { ref, onMounted } from "vue";
import axios from "axios";

export default {
  data() {
    return {
      properties: [],
    };
  },
  setup() {
    const timeFlorida = ref("");
    const timeTurkey = ref("");
    const timeTokyo = ref("");
    const timeLondon = ref("");

    const locations = {
      Florida: -4,
      Turkey: 3,
      Tokyo: 9,
      London: 1,
    };

    const updateTime = () => {
      const now = new Date();

      const formatTime = (hours, minutes) => {
        return `${hours.toString().padStart(2, "0")}:${minutes
          .toString()
          .padStart(2, "0")}`;
      };

      for (const [location, offset] of Object.entries(locations)) {
        const localTime = new Date(now.getTime() + offset * 3600 * 1000);
        const hours = localTime.getUTCHours();
        const minutes = localTime.getUTCMinutes();
        const formattedTime = formatTime(hours, minutes);

        switch (location) {
          case "Florida":
            timeFlorida.value = formattedTime;
            break;
          case "Turkey":
            timeTurkey.value = formattedTime;
            break;
          case "Tokyo":
            timeTokyo.value = formattedTime;
            break;
          case "London":
            timeLondon.value = formattedTime;
            break;
        }
      }
    };

    onMounted(() => {
      updateTime();
      setInterval(updateTime, 1000);
    });

    return { timeFlorida, timeTurkey, timeTokyo, timeLondon };
  },
  created() {
    this.getProperties();
  },
  methods: {
    getProperties() {
      axios
        .get("/api/getProperties/1/all/all")
        .then((response) => {
          if (response != undefined) {
            console.log(response.data);
            this.properties = response.data.properties;
          }
        })
        .catch((error) => {
          console.error("Get Error:", error);
        });
    },
  },
};
</script>