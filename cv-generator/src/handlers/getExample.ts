import { Request, Response } from "express"
import path from "path"

export const getExample = (req: Request, res: Response) => {
    const filePath = path.join(__dirname, "../../public", "correct.yaml");
    res.download(filePath, "example.yaml");
}