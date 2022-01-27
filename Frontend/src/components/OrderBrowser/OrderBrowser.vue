<template>
  <main>
    <div id="title">
      <h1>{{ msg }}</h1>
    </div>

    <SearchFilter @searchChange="onSearchChange" />

    <DateFilter
      @startDateChange="onStartDateChange"
      @endDateChange="onEndDateChange"
    />

    <OrderTable :orders="ordersFiltered" />
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
      orders: [
        { orderName: "Runoob", orderDate: "1991-1-1" },
        { orderName: "Google", orderDate: "1991-1-2" },
        { orderName: "Taobao", orderDate: "1991-1-3" },
        { orderName: "Taobao2", orderDate: "1991-1-4" },
        { orderName: "Taobao3", orderDate: "1991-1-5" },
        { orderName: "Taobao4", orderDate: "1991-1-6" },
        { orderName: "Taobao5", orderDate: "1991-1-7" },
        { orderName: "Taobao6" },
        { orderName: "Taobao7" },
        { orderName: "Taobao8" },
        { orderName: "Taobao0" },
      ],
      ordersFiltered: [
        { orderName: "Runoob", orderDate: "2020-01-02T15:34:12Z" },
        { orderName: "Google", orderDate: "2020-01-15T17:34:12Z" },
        { orderName: "Taobao", orderDate: "2020-01-05T05:34:12Z" },
        { orderName: "Taobao2", orderDate: "2020-02-02T15:34:12Z" },
        { orderName: "Taobao3", orderDate: "2020-01-03T10:34:12Z" },
        { orderName: "Taobao4", orderDate: "2020-01-02T15:34:14Z" },
        { orderName: "Taobao5", orderDate: "2020-01-15T17:34:16Z" },
      ],
    };
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
      var data = [this.search, this.startDate, this.endDate];
      axios({
        method: "POST",
        url: "http://127.0.0.1:8090/filter",
        data: data,
        headers: { "content-type": "text/plain" },
      })
        .then((result) => {
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
