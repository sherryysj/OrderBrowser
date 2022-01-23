<template>
  <main>
    <div id="title">
      <h1>{{ msg }}</h1>
    </div>

    <Search />
  </main>
</template>

<script>
import Search from "./Search.vue";

export default {
  name: "OrderBrowser",
  props: {
    msg: String,
  },
  components: {
    Search,
  },
};
</script>

<style>
#title {
  padding-bottom: 20px;
}
</style>
