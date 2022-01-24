<template>
  <main>
    <div id="title">
      <h1>{{ msg }}</h1>
    </div>

    <Search />

    <DateFilter />

    <div>
      <p id="total-amount">Total Amount: ${{ totalAmount }}</p>
    </div>

    <OrderTable />
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
      totalAmount: 100,
    };
  },
};
</script>

<style>
#title {
  padding-bottom: 20px;
  width: 500px;
  margin: auto;
}
#total-amount {
  text-align: left;
  margin-left: 50px;
  margin-top: 10px;
  margin-bottom: 0;
  font-weight: bold;
}
</style>
