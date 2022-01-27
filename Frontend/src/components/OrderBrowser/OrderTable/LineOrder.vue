<template>
  <tr>
    <td>
      <p style="margin-bottom: 0">{{ orderName }}</p>
      <p style="margin-bottom: 0">{{ productName }}</p>
    </td>
    <td>{{ customerCompany }}</td>
    <td>{{ customerName }}</td>
    <td>{{ orderDate }}</td>
    <td>${{ deliveredAmount }}</td>
    <td>${{ totalAmount }}</td>
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
  name: "LineOrder",
};
</script>
