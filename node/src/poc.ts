import express from "express";
import { ListContinents } from "./continent/continent";
import { GetImage } from "./continent/imageHandler";

const app = express();

app.listen(3000, () => console.log("Ready"));

app.get("/continent", (req, res) => {
    res.send(ListContinents(""));
});

app.get("/continent/:name", (req, res) => {
    const continent = ListContinents(req.params.name);
    if (!continent) {
        res.status(404).send("Could not find continent with desired name.");
    } else {
        res.send(continent);
    }
});

app.get("/continent/image/:name", (req, res) => {
    const image = GetImage(req.params.name);
    if (!image) {
        res.status(404).send("Could not find image with desired name.");
    } else {
        res.type("image/svg+xml");
        res.send(image);
    }
});
