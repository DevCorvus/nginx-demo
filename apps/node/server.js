const express = require("express");

const PORT = 8000;

const app = express();

app.get("/", (req, res) => {
  res.send('<h1 style="color: gold">Hello world from Node</h1>');
});

app.listen(PORT, () => {
  console.log("Server running on port", PORT);
});
