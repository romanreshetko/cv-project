import express from "express";
import multer from "multer"
import cors from 'cors'
import { generateResume } from "./handlers/generateResume";
import { getExample } from "./handlers/getExample";

const app = express();
const upload = multer({storage: multer.memoryStorage()});
const port = 3000;

app.use(express.json());
app.use(cors({origin: "*"}))

app.post('/generate', upload.single("file"), generateResume);

app.get('/example', getExample);

app.listen(port, () => {
    console.log("Server is running on port 3000");
})