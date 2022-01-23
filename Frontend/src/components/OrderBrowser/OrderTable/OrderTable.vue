<template>
  <table class="table table-striped table-sm" id="order-table">
    <thead>
      <tr>
        <th scope="col">Order name</th>
        <th scope="col">Customer company</th>
        <th scope="col">Customer name</th>
        <th scope="col">
          Order date
          <button v-if="sortUp" v-on:click="setSortup()" id="sortButton">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="16"
              height="16"
              fill="currentColor"
              class="bi bi-sort-down-alt"
              viewBox="0 0 16 16"
            >
              <path
                d="M3.5 3.5a.5.5 0 0 0-1 0v8.793l-1.146-1.147a.5.5 0 0 0-.708.708l2 1.999.007.007a.497.497 0 0 0 .7-.006l2-2a.5.5 0 0 0-.707-.708L3.5 12.293V3.5zm4 .5a.5.5 0 0 1 0-1h1a.5.5 0 0 1 0 1h-1zm0 3a.5.5 0 0 1 0-1h3a.5.5 0 0 1 0 1h-3zm0 3a.5.5 0 0 1 0-1h5a.5.5 0 0 1 0 1h-5zM7 12.5a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 0-1h-7a.5.5 0 0 0-.5.5z"
              />
            </svg>
          </button>

          <button v-else v-on:click="setSortup()" id="sortButton">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="16"
              height="16"
              fill="currentColor"
              class="bi bi-sort-up"
              viewBox="0 0 16 16"
            >
              <path
                d="M3.5 12.5a.5.5 0 0 1-1 0V3.707L1.354 4.854a.5.5 0 1 1-.708-.708l2-1.999.007-.007a.498.498 0 0 1 .7.006l2 2a.5.5 0 1 1-.707.708L3.5 3.707V12.5zm3.5-9a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5zM7.5 6a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1h-5zm0 3a.5.5 0 0 0 0 1h3a.5.5 0 0 0 0-1h-3zm0 3a.5.5 0 0 0 0 1h1a.5.5 0 0 0 0-1h-1z"
              />
            </svg>
          </button>
        </th>
        <th scope="col">Delivered Amount</th>
        <th scope="col">Total Amount</th>
      </tr>
    </thead>
    <tbody>
      <LineOrder
        v-for="order in orders"
        v-bind:key="order"
        :orderName="order.orderName"
        customerName="Sherry"
        customerCompany="company"
        orderDate="1999-1-1"
        :deliveredAmount="11.3"
        :totalAmount="24"
      />
    </tbody>
  </table>

  <PageControl :orderAmount="orderAmount" @pageChange="onPageChange" />
</template>

<script>
import LineOrder from "./LineOrder";
import PageControl from "./PageControl";

export default {
  name: "OrderTable",
  data: function () {
    return {
      sortUp: true,
      orderAmount: 100,
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
    };
  },
  methods: {
    setSortup() {
      this.sortUp = !this.sortUp;
    },
    onPageChange(value) {
      console.log(value);
    },
  },
  components: {
    LineOrder,
    PageControl,
  },
};
</script>

<style>
#order-table {
  width: 95%;
  margin: auto;
  margin-top: 30px;
  margin-bottom: 10px;
}
#sortButton {
  padding: 0;
  border: none;
  background-color: #fff;
}
</style>
