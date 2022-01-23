<template>
  <main>
    <div id="title">
      <h1>{{ msg }}</h1>
    </div>

    <Search />
    <DateFilter />
  </main>
</template>

<script>
import Search from "./Search.vue";
import DateFilter from "./DateFilter.vue";

export default {
  name: "OrderBrowser",
  props: {
    msg: String,
  },
  components: {
    Search,
    DateFilter,
  },
};
</script>

<style>
#title {
  padding-bottom: 20px;
}
</style>
