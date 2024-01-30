<template>
  <div id="app">
    <button @click="downloadCsv">Download CSV</button>
  </div>
</template>

<script>
export default {
  name: "App",
  methods: {
    async downloadCsv() {
      try {
        const res = await fetch("http://localhost:8080/download_csv");
        const bom = new Uint8Array([0xef, 0xbb, 0xbf]);
        const blob = new Blob([bom, res], { type: "text/csv" });
        const link = document.createElement("a");
        link.href = window.URL.createObjectURL(blob);
        link.download = "downloaded_file.csv";
        link.click();
      } catch (error) {
        console.error("Failed to download CSV:", error);
      }
    },
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
