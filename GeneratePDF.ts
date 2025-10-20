import { env } from "node:process";
import * as puppeteer from "puppeteer";
import * as fs from "@std/fs";

const statics = env.GENERATED;

if (!statics) {
  throw "No generated directory configured (env GENERATED).";
}

const server = Deno.serve(async (req) => {
  if (req.method !== "GET") {
    return new Response(
      "",
      {
        status: 404
      }
    );
  }

  const url = new URL(req.url);

  const path = url.pathname;

  const filepath = statics + path;
  const fileExists = await fs.exists(filepath);
  
  if (!fileExists) {
    return new Response(
      "",
      {
        status: 404
      }
    );   
  }

  const content = await Deno.readTextFile(filepath);

  let contentType = "";
  if (filepath.endsWith(".html")) {
    contentType = "text/html; charset=utf-8";
  } else if (filepath.endsWith(".css")) {
    contentType = "text/css; charset=utf-8";
  }

  return contentType ?
    new Response(
      content,
      {
        status: 200,
        headers: {
          "Content-Type": contentType
        }
      }
    ) :
    new Response(
      "",
      {
        status: 404
      }
    );
});

const browser = await puppeteer.launch();

const page = await browser.newPage();

await page.goto("http://localhost:8000/backend.html", {
  waitUntil: "networkidle2"
});

await page.pdf({
  path: 'test.pdf'
});

await browser.close();

await server.shutdown();
