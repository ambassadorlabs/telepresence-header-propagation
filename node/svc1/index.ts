import express, { Express, Request } from "express";
import * as ax from "axios";

const axios = ax.default

const PORT: number = parseInt(process.env.PORT || "8080");
const svc2Addr: string = process.env.SVC2_ADDR || "localhost:8081"
const app: Express = express();

app.get("/uppercase", async (req: Request<{subject: string}>, res) => {
  for (let header in req.headers) {
    if (header.toLowerCase() === "baggage") {
      console.log(`uppercase header ${header} equals ${req.headers[header]}`)
    }
  }
  if (!req.query.subject || typeof req.query.subject !== "string") {
    res.status(400)
    return
  }
  const subject = req.query.subject;
  const finalres = await axios.get(`http://${svc2Addr}/finalupper?subject=${subject}`)
  const uppered = finalres.data
  res.send(uppered);
});

app.listen(PORT, () => {
  console.log(`Listening for requests on http://localhost:${PORT}`);
});