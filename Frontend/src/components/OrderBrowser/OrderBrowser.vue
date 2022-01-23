<template>
  <main>
    <div id="title">
      <h1>{{ msg }}</h1>
    </div>

    <Search />
    <DateFilter />
    <OrderTable />
  </main>
</template>

<script>
import Search from "./Search.vue";
import DateFilter from "./DateFilter.vue";
import OrderTable from "./OrderTable/OrderTable.vue";

export default {
  name: "OrderBrowser",
  props: {
    msg: String,
  },
  components: {
    Search,
    DateFilter,
    OrderTable,
  },
};
</script>

<style>
#title {
  padding-bottom: 20px;
  width: 500px;
  margin: auto;
}
</style>
