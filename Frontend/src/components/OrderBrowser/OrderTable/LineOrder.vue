<template>
  <tr>
    <td>
      <p style="margin-bottom: 0">{{ orderName }}</p>
      <p style="margin-bottom: 0">{{ productName }}</p>
    </td>
    <td>{{ customerCompany }}</td>
    <td>{{ customerName }}</td>
    <td>{{ orderDate }}</td>
    <td>{{ deliveredAmountString() }}</td>
    <td>{{ totalAmountString() }}</td>
  </tr>
</template>

<script>
export default {
  props: {
    orderName: String,
    productName: String,
    customerCompany: String,
    customerName: String,
    orderDate: String,
    deliveredAmount: Number,
    totalAmount: Number,
  },
  methods: {
    deliveredAmountString() {
      if (this.deliveredAmount == 0) {
        return "/";
      } else {
        var string = "$" + this.deliveredAmount.toFixed(4);
        return string;
      }
    },
    totalAmountString() {
      if (this.totalAmount == 0) {
        return "Not available without unit price";
      } else {
        var string = "$" + this.totalAmount.toFixed(4);
        return string;
      }
    },
  },
  name: "LineOrder",
};
</script>
