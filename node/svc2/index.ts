import express, { Express, Request } from "express";
import * as ax from "axios";

const axios = ax.default

const PORT: number = parseInt(process.env.PORT || "8081");
const app: Express = express();

app.get("/finalupper", (req, res) => {
  for (let header in req.headers) {
    if (header.toLowerCase() === "baggage") {
      console.log(`finalupper header ${header} equals ${req.headers[header]}`)
    }
  }
  if (!req.query.subject || typeof req.query.subject !== "string") {
    res.status(400).send()
    return
  }
  const subject = req.query.subject as string;
  const uppered = subject.toUpperCase()
  res.status(200).send(uppered)
});

app.listen(PORT, () => {
  console.log(`Listening for requests on http://localhost:${PORT}`);
});