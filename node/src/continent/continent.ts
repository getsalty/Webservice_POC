import fs from "fs";

var pathName = "../common/data.json";

type Continent = {
    name: string;
    displayName: string;
    desc: string;
};

type Data = {
    continents: Continent[];
};

export const ListContinents = (name: string) => {
    const byteValue = fs.readFileSync(pathName, "utf8");

    const data = JSON.parse(byteValue) as Data;

    if (!name) {
        return data.continents;
    }

    return data.continents?.find((o) => o.name === name);
};
