<template>
  <OrderBrowser msg="Order Browsing Page" />
</template>

<script>
import OrderBrowser from "./components/OrderBrowser/OrderBrowser.vue";

export default {
  name: "App",
  components: {
    OrderBrowser,
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 20px;
  min-width: 600px;
}
</style>
