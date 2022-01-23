<template>
  <div id="date-container">
    <div id="date-title">
      <p id="date-title-text">Create Date</p>
    </div>
    <div class="row" id="date-input">
      <input type="date" class="col" id="date-input-start" />
      <p class="col" id="date-to">-</p>
      <input type="date" class="col" id="date-input-end" />
    </div>
  </div>
</template>

<script>
export default {
  name: "DateFilter",
};
</script>

<style>
#date-container {
  margin-top: 20px;
  margin-left: 20px;
}
#date-title {
  width: 100px;
  height: 25px;
}
#date-title-text {
  margin-bottom: 0;
  text-align: left;
  font-weight: bold;
}
#date-input {
  border-style: groove;
  width: 350px;
}
#date-to {
  flex-grow: 0.01;
  margin-bottom: 0;
}
#date-input-start {
  flex-grow: 0.5;
  border: none;
  padding: 5px;
}
#date-input-end {
  flex-grow: 0.5;
  border: none;
  padding: 5px;
}
</style>
