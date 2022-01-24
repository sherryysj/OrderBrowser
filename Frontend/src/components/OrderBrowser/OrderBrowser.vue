<template>
  <main>
    <div id="title">
      <h1>{{ msg }}</h1>
    </div>

    <Search @searchChange="onSearchChange" />

    <DateFilter />

    <OrderTable :orders="ordersFiltered" />
  </main>
</template>

<script>
import Search from "./Search.vue";
import DateFilter from "./DateFilter.vue";
import OrderTable from "./OrderTable/OrderTable.vue";

export default {
  name: "OrderBrowser",
  props: {
    msg: String,
  },
  components: {
    Search,
    DateFilter,
    OrderTable,
  },
  data: function () {
    return {
      search: "",
      orders: [
        { orderName: "Runoob" },
        { orderName: "Google" },
        { orderName: "Taobao" },
        { orderName: "Taobao2" },
        { orderName: "Taobao3" },
        { orderName: "Taobao4" },
        { orderName: "Taobao5" },
        { orderName: "Taobao6" },
        { orderName: "Taobao7" },
        { orderName: "Taobao8" },
        { orderName: "Taobao0" },
      ],
      ordersFiltered: [
        { orderName: "Runoob" },
        { orderName: "Google" },
        { orderName: "Taobao" },
        { orderName: "Taobao2" },
        { orderName: "Taobao3" },
        { orderName: "Taobao4" },
        { orderName: "Taobao5" },
      ],
    };
  },
  methods: {
    onSearchChange(value) {
      this.search = value;
      // call orders filter
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
