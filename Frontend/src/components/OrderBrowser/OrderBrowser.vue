<template>
  <main>
    <div id="title">
      <h1>{{ msg }}</h1>
    </div>

    <SearchFilter @searchChange="onSearchChange" @submit.prevent="" />

    <DateFilter
      @startDateChange="onStartDateChange"
      @endDateChange="onEndDateChange"
    />

    <OrderTable :orders="orders" v-bind:key="orders" />
  </main>
</template>

<script>
import SearchFilter from "./Filters/SearchFilter.vue";
import DateFilter from "./Filters/DateFilter.vue";
import OrderTable from "./OrderTable/OrderTable.vue";
import axios from "axios";

export default {
  name: "OrderBrowser",
  props: {
    msg: String,
  },
  components: {
    SearchFilter,
    DateFilter,
    OrderTable,
  },
  data: function () {
    return {
      search: "",
      startDate: "",
      endDate: "",
      orders: [],
    };
  },
  created: function () {
    this.callServerFilter();
  },
  methods: {
    onSearchChange(value) {
      this.search = value;
      this.callServerFilter();
    },
    onStartDateChange(value) {
      this.startDate = value;
      this.callServerFilter();
    },
    onEndDateChange(value) {
      // must later than start date
      this.endDate = value;
      this.callServerFilter();
    },
    callServerFilter() {
      var data = {
        search: this.search,
        startDate: this.startDate,
        endDate: this.endDate,
      };
      axios({
        method: "POST",
        url: "http://127.0.0.1:8090/filter",
        data: data,
        headers: { "content-type": "text/plain" },
      })
        .then((result) => {
          this.orders = result.data;
          console.log(this.orders);
          console.log(result.data);
        })
        .catch((error) => {
          console.error(error);
        });
    },
  },
};
</script>

<style>
#title {
  padding-bottom: 20px;
  width: 500px;
  margin: auto;
}
</style>
