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
    <PageControl
      :orderAmount="orderAmount"
      @pageChange="onPageChange"
      @ordersShownChange="onOrdersShownChange"
    />
  </main>
</template>

<script>
import Search from "./Search.vue";
import DateFilter from "./DateFilter.vue";
import OrderTable from "./OrderTable/OrderTable.vue";
import PageControl from "./PageControl.vue";

export default {
  name: "OrderBrowser",
  props: {
    msg: String,
  },
  components: {
    Search,
    DateFilter,
    OrderTable,
    PageControl,
  },
  data: function () {
    return {
      totalAmount: 100,
      orderAmount: 100,
      page: 1,
      ordersPerPage: 5,
    };
  },
  methods: {
    onPageChange(value) {
      this.page = value;
    },
    onOrdersShownChange(value) {
      this.ordersPerPage = value;
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
#total-amount {
  text-align: left;
  margin-left: 50px;
  margin-top: 10px;
  margin-bottom: 0;
  font-weight: bold;
}
</style>
