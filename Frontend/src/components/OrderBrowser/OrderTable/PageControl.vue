<template>
  <div class="container">
    <div>
      <p style="margin-bottom: 0">Total Orders: {{ orderAmount }}</p>
    </div>
    <div class="displayInRow" id="page-controls">
      <div class="displayInRow" id="orderPerPage">
        <p style="margin-bottom: 0">Orders per page:</p>
        <select name="orderPerPage" style="margin-left: 10px">
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
        <input id="page-input" />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    orderAmount: {
        type:Number,
        default:50
    }
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
      if (this.page > 1) {
        this.page--;
        this.$emit("pageChange", this.page);
      }
    },
    pagePlus() {
      if (this.page < this.orderAmount / this.ordersPerPage) {
        this.page++;
        this.$emit("pageChange", this.page);
      }
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
