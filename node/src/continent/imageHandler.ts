import fs from "fs";

const dirPathName = "../common/images/";

export const GetImage = (name: string) => {
    const imagePath = dirPathName + name + ".svg";

    try {
        return fs.readFileSync(imagePath, "utf8");
    } catch (error) {
        return "";
    }
};
