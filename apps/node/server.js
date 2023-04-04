const express = require("express");

const PORT = 8000;

const app = express();

app.get("/", (req, res) => {
  res.send("Hello world from Node");
});

app.listen(PORT, () => {
  console.log("Server running on port", PORT);
});
