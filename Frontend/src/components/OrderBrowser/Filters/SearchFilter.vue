<template>
  <div class="row" id="search">
    <div class="col" id="search-icon">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="20"
        height="20"
        fill="currentColor"
        class="bi bi-search"
        viewBox="0 0 16 16"
      >
        <path
          d="M11.742 10.344a6.5 6.5 0 1 0-1.397 1.398h-.001c.03.04.062.078.098.115l3.85 3.85a1 1 0 0 0 1.415-1.414l-3.85-3.85a1.007 1.007 0 0 0-.115-.1zM12 6.5a5.5 5.5 0 1 1-11 0 5.5 5.5 0 0 1 11 0z"
        />
      </svg>
    </div>

    <div class="col" id="search-title">
      <p style="margin-bottom: 0">Search</p>
    </div>

    <div class="col-sm" id="search-input">
      <form>
        <input
          class="form-control"
          id="search-input"
          placeholder="search the order here"
          @change="handleChange"
        />
      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: "SearchFilter",
  methods: {
    handleChange(e) {
      var search = e.target.value;
      this.$emit("searchChange", search);
    },
  },
};
</script>

<style>
#search {
  align-items: center;
  margin-left: 40px;
  margin-right: 10px;
}
#search-icon {
  flex-grow: 0.01;
}
#search-title {
  flex-grow: 0.01;
  font-weight: bold;
  font-size: 20px;
  padding-left: 0;
}
#search-input {
  flex-grow: 0.85;
  width: 400px;
}
</style>
