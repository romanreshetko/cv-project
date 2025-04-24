import { Request, Response } from "express"
import yaml from "js-yaml";
import { renderToStaticMarkup } from "react-dom/server";
import ResumePage from "../components/resumePage";
import { ResumeData } from "../types/resumeTypes";
import puppeteer from "puppeteer";
import archiver from "archiver";
import fs from "fs";
import path from "path";

export const generateResume = async (req: Request, res: Response) => {
    try {
        if (!req.file) {
            res.status(400).json({ message: "File error" });
            return;
        }

        const fileContent = req.file.buffer.toString("utf-8");
        const resumeData = yaml.load(fileContent) as ResumeData;

        if (!resumeData || typeof resumeData !== "object") {
            res.status(400).json({ message: "Invalid YAML format" });
            return;
        }

        const styles = fs.readFileSync(path.join(__dirname, "../globals.css"), "utf-8");

        const htmlContent = renderToStaticMarkup(ResumePage({resumeData}));

        const fullHtml = `
        <html>
        <head>
            <style>${styles}</style>
        </head>
        <body>
            ${htmlContent}
        </body>
        </html>`;
        
        const browser = await puppeteer.launch({
            headless: true,
            args: ['--no-sandbox', '--disable-setuid-sandbox']
        });
        const page = await browser.newPage();

        await page.setContent(fullHtml, {waitUntil: 'load'});

        const pdfBuffer = await page.pdf({format: 'A4'});
        await browser.close();

        res.setHeader("Content-Type", "application/zip");
        res.setHeader("Content-Disposition", 'attachment; filename="resume.zip"');
        //res.end(pdfBuffer);

        const archive = archiver("zip", {zlib: {level: 9}});

        archive.append(fullHtml, {name: "resume.html"});
        archive.append(Buffer.from(pdfBuffer), {name: "resume.pdf"});

        archive.pipe(res);
        archive.finalize();
      

        //res.status(200).json({ message: "Resume page generated successfully", html: fullHtml });

    } catch (error) {
        console.log(error);
        res.status(500).json({ message: "Server error" });
    }
}