import express, { Express, Request } from "express";
import * as ax from "axios";

const axios = ax.default

const PORT: number = parseInt(process.env.PORT || "8080");
const finalupperAddr: string = process.env.FINALUPPER_ADDR || "localhost:8081"
const app: Express = express();

app.get("/uppercase", async (req: Request<{subject: string}>, res) => {
  if (!req.query.subject || typeof req.query.subject !== "string") {
    res.status(400)
    return
  }
  const subject = req.query.subject;
  const finalres = await axios.get(`http://${finalupperAddr}/finalupper?subject=${subject}`)
  const uppered = finalres.data
  res.send(uppered);
});

app.listen(PORT, () => {
  console.log(`Listening for requests on http://localhost:${PORT}`);
});