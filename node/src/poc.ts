import express from "express";
import { ListContinents } from "./continent/continent";
import { GetImage } from "./continent/imageHandler";

const app = express();

app.listen(3000, () => console.log("Ready"));

app.get("/continent", (req, res) => {
    res.send(ListContinents(""));
});

app.get("/continent/:name", (req, res) => {
    res.send(ListContinents(req.params.name));
});

app.get("/continent/image/:name", (req, res) => {
    res.send(GetImage(req.params.name));
});
