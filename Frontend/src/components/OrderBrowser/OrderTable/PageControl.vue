<template>
  <div class="container">
    <div>
      <p style="margin-bottom: 0">Total Orders: {{ orderAmount }}</p>
    </div>
    <div class="displayInRow" id="page-controls">
      <div class="displayInRow" id="orderPerPage">
        <p style="margin-bottom: 0">Orders per page:</p>
        <select
          name="orderPerPage"
          style="margin-left: 10px"
          @change="handleOrdersPerPage"
        >
          <option value="5">5</option>
          <option value="10">10</option>
          <option value="15">15</option>
          <option value="20">20</option>
          <option value="25">25</option>
        </select>
      </div>
      <div class="displayInRow" id="page-move">
        <button class="page-buttons" v-on:click="pageMinus()">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="16"
            height="16"
            fill="currentColor"
            class="bi bi-chevron-left"
            viewBox="0 0 16 16"
          >
            <path
              fill-rule="evenodd"
              d="M11.354 1.646a.5.5 0 0 1 0 .708L5.707 8l5.647 5.646a.5.5 0 0 1-.708.708l-6-6a.5.5 0 0 1 0-.708l6-6a.5.5 0 0 1 .708 0z"
            />
          </svg>
        </button>
        <p style="margin-bottom: 0">{{ page }}</p>
        <button class="page-buttons" v-on:click="pagePlus()">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="16"
            height="16"
            fill="currentColor"
            class="bi bi-chevron-right"
            viewBox="0 0 16 16"
          >
            <path
              fill-rule="evenodd"
              d="M4.646 1.646a.5.5 0 0 1 .708 0l6 6a.5.5 0 0 1 0 .708l-6 6a.5.5 0 0 1-.708-.708L10.293 8 4.646 2.354a.5.5 0 0 1 0-.708z"
            />
          </svg>
        </button>
      </div>
      <div class="displayInRow">
        <p style="margin-bottom: 0">Go to</p>
        <input id="page-input" @change="handlePageChange" />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    orderAmount: {
      type: Number,
      default: 50,
    },
  },
  data: function () {
    return {
      page: 1,
      ordersPerPage: 5,
    };
  },
  name: "PageControl",
  methods: {
    pageMinus() {
      // only minus page value when page bigger than 1
      if (this.page > 1) {
        this.page--;
        this.$emit("pageChange", this.page);
      }
    },
    pagePlus() {
      // only plus page value when page smaller than max page
      if (this.page < this.countMaxPage()) {
        this.page++;
        this.$emit("pageChange", this.page);
      }
    },
    handleOrdersPerPage(e) {
      var ordersNumber = e.target.value;
      this.ordersPerPage = ordersNumber;
      // if orders shown per page more than order amount, reset and jump to page 1
      if (this.ordersPerPage >= this.orderAmount) {
        this.page = 1;
        this.$emit("pageChange", this.page);
      }
      this.$emit("ordersShownChange", this.ordersPerPage);
    },
    handlePageChange(e) {
      var page = e.target.value;
      // only change page value when page input bigger than 0 and smaller or equal to max page
      if (page > 0 && page <= this.countMaxPage()) {
        this.page = page;
        this.$emit("pageChange", this.page);
      }
    },
    countMaxPage() {
      var maxPage = Math.ceil(this.orderAmount / this.ordersPerPage);
      return maxPage;
    },
  },
};
</script>

<style>
.page-buttons {
  width: 30px;
  margin-left: 10px;
  margin-right: 10px;
  border: none;
  background-color: #fff;
}
#page-controls {
  justify-content: center;
}
.displayInRow {
  align-items: center;
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  margin-left: 10px;
  margin-right: 10px;
}
#page-input {
  width: 50px;
  margin-left: 10px;
}
#orderPerPage {
  margin-left: -70px;
  margin-right: 20px;
}
</style>
