<template>
  <OrderBrowser msg="Welcome to Order Browse Page" />
</template>

<script>
import OrderBrowser from "./components/OrderBrowser.vue";

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
  margin-top: 60px;
}
</style>
