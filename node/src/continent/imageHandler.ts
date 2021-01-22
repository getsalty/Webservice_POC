import fs from "fs";

const dirPathName = "../common/images/";

export const GetImage = (name: string) => {
    if (name.length === 0) {
        return;
    }

    const imagePath = dirPathName + name + ".svg";

    return fs.readFileSync(imagePath, "utf8");
};
