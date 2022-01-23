<template>
  <tr>
    <td>{{ orderName }}</td>
    <td>{{ customerCompany }}</td>
    <td>{{ customerName }}</td>
    <td>{{ orderDate }}</td>
    <td>{{ deliveredAmount }}</td>
    <td>{{ totalAmount }}</td>
  </tr>
</template>

<script>
export default {
  props: {
    orderName: String,
    customerCompany: String,
    customerName: String,
    orderDate: String,
    deliveredAmount: Number,
    totalAmount: Number,
  },
  name: "LineOrder",
};
</script>
